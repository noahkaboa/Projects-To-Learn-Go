package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	targetURL := "https://noahkaboa.github.io"

	targetElement := "h1"

	resp, err := http.Get(targetURL)
	if err != nil {
		fmt.Println("Trouble sending get request")
		fmt.Println(err)
	}

	fmt.Println(resp.Header)

	bodyBytes := streamBody(resp.Body)

	defer resp.Body.Close()

	fmt.Println(string(bodyBytes))

	doc, err := html.Parse(strings.NewReader(string(bodyBytes)))
	if err != nil {
		fmt.Println(err)
	}

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == targetElement {
			fmt.Println(n.FirstChild.Data)
		}
	}
}

func streamBody(rc io.ReadCloser) []byte {
	bodyBytes := make([]byte, 0)
	for {
		streamedByte := make([]byte, 1)
		count, err := rc.Read(streamedByte)
		if err != nil {
			fmt.Println("Error with streaming body")
		}
		if count == 0 {
			return bodyBytes
		}
		bodyBytes = append(bodyBytes, streamedByte...)
	}
}
