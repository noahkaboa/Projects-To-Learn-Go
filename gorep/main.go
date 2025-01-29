package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

type Result struct {
	found  bool
	result []string
	path   string
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
		found := searchResult[0].found
		results := searchResult[0].result

		if found {
			for i := 0; i < len(results); i++ {
				fmt.Println(strings.TrimSpace(results[i]))
			}
		}
	} else if filenameFlag == "" {
		f, err := os.Open(dirFlag)
		if err != nil {
			fmt.Println(err)
			return
		}
		searchResult := dirCheck(f, termFlag, workerCountFlag)

		for i := 0; i < len(searchResult); i++ {
			if searchResult[i].found {
				fmt.Println(searchResult[i].path)

				for j := 0; j < len(searchResult[i].result); j++ {
					fmt.Println(strings.TrimSpace(searchResult[i].result[j]))
				}
			}
		}

	} else {
		fmt.Println("Both file and directory provided. Exiting")
		return
	}

}

func dirCheck(dir *os.File, term string, worker int) []Result {
	files, err := dir.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	results := []Result{}

	for _, v := range files {
		if v.IsDir() {
			results = append(results, dirCheckFromInfo(v, term, worker)...)
		} else {
			results = append(results, fileCheckFromInfo(v, term, worker)...)
		}
	}

	return results
}

func dirCheckFromInfo(dirInfo fs.FileInfo, term string, worker int) []Result {
	f, err := os.Open(dirInfo.Name())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dirCheck(f, term, worker)
}

func fileCheckFromInfo(fileInfo fs.FileInfo, term string, worker int) []Result {
	f, err := os.Open(fileInfo.Name())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return fileCheck(f, term, worker)
}

func fileCheck(file *os.File, term string, workerCount int) []Result {
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

	path, _ := os.Getwd()
	results := make(chan string, fileLen)
	endResult := Result{false, []string{}, path + file.Name()}

	for w := 1; w <= workerCount; w++ {
		go worker(jobs, results, term)
	}

	for a := 1; a <= fileLen; a++ {
		r := <-results
		if r != "" {
			endResult.found = true
			endResult.result = append(endResult.result, r)
		}
	}
	close(jobs)

	return []Result{endResult}

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
