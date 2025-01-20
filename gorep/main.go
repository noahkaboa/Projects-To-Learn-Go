package main

import (
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

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", fileCheck(data, term))

}

func fileCheck(file []byte, term string) Result {
	fileString := string(file)
	if strings.Contains(fileString, term) {
		return Result{found: true, result: term}
	} else {
		return Result{found: false, result: ""}
	}
}

func contains(slice []string, term string) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == term {
			return true
		}
	}
	return false
}
