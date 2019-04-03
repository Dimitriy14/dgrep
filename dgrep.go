package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var regx string

type match struct {
	lineNumber int
	strs       string
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Not enough arguments")
		fmt.Println("Want main.go [options] [string] [root path]")
		return
	}

	contextLines := flag.Int("ctx", 3, "an int")

	ignoreCase := flag.Bool("case", false, "a bool")

	pattern := flag.String("p", "go", "-p=go")

	flag.Parse()

	root := os.Args[len(os.Args)-1]

	if *ignoreCase {
		regx = "(?i)" + os.Args[len(os.Args)-2]
	} else {
		regx = "(?-i)" + os.Args[len(os.Args)-2]
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Mode().IsRegular() {
			if filepath.Ext(info.Name()) == "."+*pattern {
				wg.Add(1)

				go parseWithIndex(path, *contextLines)
			}
		}
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	wg.Wait()
}

func parseFile(path string, contextLines int) {
	dat, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(regx)

	var matches []match

	lines := strings.Split(string(dat), "\n")

	for i, line := range lines {
		if re.FindString(line) != "" {
			var context []string

			if len(lines) < contextLines {
				context = append(context, line)

			} else if i > len(lines)-contextLines {
				for j := -contextLines + 1; j <= 0; j++ {
					context = append(context, lines[i+j])
				}
			} else {
				for j := 0; j < contextLines; j++ {
					context = append(context, lines[i+j])
				}
			}

			matches = append(matches, match{i + 1, strings.Join(context, "\n")})
		}
	}

	if matches != nil {
		fmt.Println(path + " : ")
		for _, m := range matches {
			fmt.Printf("Match on line %v:\n%v\n", m.lineNumber, m.strs)
		}
	}

	wg.Done()
}

type Match struct {
	lines string
	index []int
}

func parseWithIndex(path string, contextLines int) {
	myregexp := ".*" + regx

	for i := 0; i < contextLines; i++ {
		myregexp += ".*\n?"
	}

	dat, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(myregexp)

	res := re.FindAllString(string(dat), -1)

	indexes := re.FindAllStringIndex(string(dat), -1)

	var matches []Match

	for i, r := range res {
		matches = append(matches, Match{r, indexes[i]})
	}

	fmt.Println(path + " : ")

	for _, m := range matches {
		fmt.Println("lines: ", m.lines, "index:", m.index)
	}

	wg.Done()
}
