package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

var (
	regx         string
	contextLines = flag.Int("c", 3, "number of context lines")
	ignoreCase   = flag.Bool("i", false, "ignore case")
	pattern      = flag.String("type", "go", "file extension")
)

//Results is a result of parsing
type Results struct {
	path  string
	lines string
}

func main() {
	flag.Parse()

	root := os.Args[len(os.Args)-1]

	if *ignoreCase {
		regx = "(?i)" + os.Args[len(os.Args)-2]
	} else {
		regx = "(?-i)" + os.Args[len(os.Args)-2]
	}

	results, err := worker(root)

	if err != nil {
		log.Fatal(err)
	}

	for _, r := range results {

		fmt.Printf("%s:\n%s", r.path, r.lines)
	}
}

func worker(root string) ([]Results, error) {
	done := make(chan struct{})
	defer close(done)

	paths, errc := walkFiles(done, root)

	resch := make(chan Results)

	var wg sync.WaitGroup

	numGoroutine := 5

	wg.Add(numGoroutine)

	for i := 0; i < numGoroutine; i++ {
		go func() {
			parseFile(paths, resch, done)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(resch)
	}()

	var results []Results

	for r := range resch {
		results = append(results, r)
	}

	if err := <-errc; err != nil {
		return nil, err
	}

	return results, nil
}

func walkFiles(done chan struct{}, root string) (chan string, chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)

	go func() {
		defer close(paths)

		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.Mode().IsRegular() {
				if filepath.Ext(info.Name()) == "."+*pattern {
					select {
					case paths <- path:
					case <-done:
						return fmt.Errorf("walk canceled")
					}
				}
			}

			return nil
		})
	}()

	return paths, errc
}

func parseFile(paths chan string, res chan Results, done chan struct{}) {
	for path := range paths {
		file, err := os.Open(path)

		if err != nil {
			log.Println("file opening err:", err)
			return
		}

		defer file.Close()

		fileScanner := bufio.NewScanner(file)

		re := regexp.MustCompile(regx)

		var lineNumber int
		var context string
		var contextLineCounter int
		var mathedLines string

		fileScanner.Scan()

		for {
			lineNumber++
			contextLineCounter++

			context += fmt.Sprintf("%d: %s\n", lineNumber, fileScanner.Text())

			if contextLineCounter == *contextLines {
				if re.FindString(context) != "" {
					mathedLines += context
				}

				context = "\n"
				contextLineCounter = 0
			}

			if !fileScanner.Scan() {
				if re.FindString(context) != "" {
					mathedLines += context
				}

				break
			}
		}

		select {
		case res <- Results{path, mathedLines}:
		case <-done:
			return
		}
	}
}
