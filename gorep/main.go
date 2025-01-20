package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Result struct {
	found  bool
	result string
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

	fmt.Printf("%+v\n", fileCheck(file, term))

}

func fileCheck(file *os.File, term string) Result {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, term) {
			return Result{true, line}
		}
	}
	return Result{false, ""}
}

func contains(slice []string, term string) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == term {
			return true
		}
	}
	return false
}
