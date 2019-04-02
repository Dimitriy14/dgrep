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

type matche struct {
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

	if root == "" {
		root = "."
	}

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

				go parseFile(path, *contextLines)
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

	var matches []matche

	lines := strings.Split(string(dat), "\n")

	for i, line := range lines {
		if re.FindString(line) != "" {
			var context []string

			if i > len(lines)-contextLines {
				for j := -contextLines; j < 0; j++ {
					context = append(context, lines[i+j])
				}
			} else {
				for j := 0; j < contextLines; j++ {
					context = append(context, lines[i+j])
				}
			}

			matches = append(matches, matche{i + 1, strings.Join(context, "\n")})
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
