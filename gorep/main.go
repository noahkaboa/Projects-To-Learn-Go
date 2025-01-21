package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Result struct {
	found  bool
	result []string
}

func main() {

	if contains(os.Args, "--") {
		os.Args = os.Args[1:]
	}

	filename := os.Args[1]
	term := os.Args[2]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	searchResult := fileCheck(file, term)
	found := searchResult.found
	results := searchResult.result

	if found {
		for i := 0; i < len(results); i++ {
			fmt.Println(strings.TrimSpace(results[i]))
		}
	}

}

func fileCheck(file *os.File, term string) Result {
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

	for w := 1; w <= 3; w++ {
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

func contains(slice []string, term string) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == term {
			return true
		}
	}
	return false
}

func worker(id int, jobs <-chan string, results chan<- string, searchTerm string) {
	for j := range jobs {
		if strings.Contains(j, searchTerm) {
			results <- j
		} else {
			results <- ""
		}
	}
}
