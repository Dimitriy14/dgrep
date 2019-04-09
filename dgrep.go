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

	wg sync.WaitGroup
)

type jobRequest struct {
	path string
	info os.FileInfo
}

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

	reqch := make(chan jobRequest)
	done := make(chan struct{})

	numberOfJob := 3

	for i := 0; i < numberOfJob; i++ {
		go parseFile(i, reqch, done)
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Mode().IsRegular() {
			if filepath.Ext(info.Name()) == "."+*pattern {
				reqch <- jobRequest{path, info}
			}
		}
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	for i := 0; i < numberOfJob; i++ {
		done <- struct{}{}
	}

}

func parseFile(id int, request chan jobRequest, done chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Printf("Routine №:%d is finished\n", id)
			return
		case req := <-request:
			file, err := os.Open(req.path)

			if err != nil {
				log.Println("file opening err:", err)
				return
			}

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

			if mathedLines != "" {
				fmt.Printf("\n%s: \t%s", req.path, mathedLines)
			}
		}
	}
}
