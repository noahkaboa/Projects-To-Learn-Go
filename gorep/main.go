package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Result struct {
	found  bool
	result []string
}

func main() {

	var filenameFlag, termFlag, dirFlag string
	var ignoreCaseFlag bool
	var workerCountFlag int

	flag.StringVar(&filenameFlag, "f", "", "File that will be searched")
	flag.StringVar(&termFlag, "t", "", "Term that will be searched for")
	flag.StringVar(&dirFlag, "d", "", "Directory that will be searched")
	flag.BoolVar(&ignoreCaseFlag, "i", false, "Ignore case")
	flag.IntVar(&workerCountFlag, "w", 5, "How many workers in worker pool")

	flag.Parse()

	if filenameFlag == "" && dirFlag == "" {
		fmt.Println("Need to either provide -f or -d flag")
		return
	}

	if dirFlag == "" {
		file, err := os.Open(filenameFlag)
		if err != nil {
			fmt.Println(err)
		}

		searchResult := fileCheck(file, termFlag, workerCountFlag)
		found := searchResult.found
		results := searchResult.result

		if found {
			for i := 0; i < len(results); i++ {
				fmt.Println(strings.TrimSpace(results[i]))
			}
		}
	} else if filenameFlag == "" {

	} else {
		fmt.Println("Both file and directory provided. Exiting")
		return
	}

}

func fileCheck(file *os.File, term string, workerCount int) Result {
	counterScanner := bufio.NewScanner(file)
	jobs := make(chan string)
	fileLen := 0
	for counterScanner.Scan() {
		fileLen++
		line := counterScanner.Text()
		go func() {
			jobs <- line
		}()
	}

	results := make(chan string, fileLen)
	endResult := Result{false, []string{}}

	for w := 1; w <= workerCount; w++ {
		go worker(w, jobs, results, term)
	}

	for a := 1; a <= fileLen; a++ {
		r := <-results
		if r != "" {
			endResult.found = true
			endResult.result = append(endResult.result, r)
		}
	}
	close(jobs)

	return endResult

}

func worker(jobs <-chan string, results chan<- string, searchTerm string) {
	for j := range jobs {
		if strings.Contains(j, searchTerm) {
			results <- j
		} else {
			results <- ""
		}
	}
}
