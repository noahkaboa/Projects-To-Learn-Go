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

	endResult := Result{false, []string{}}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, term) {
			endResult.found = true
			endResult.result = append(endResult.result, line)
		}
	}
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
