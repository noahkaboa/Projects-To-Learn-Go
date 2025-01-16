package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	targetURL := "https://noahkaboa.github.io"

	resp, err := http.Get(targetURL)
	if err != nil {
		fmt.Println("Trouble sending get request")
		fmt.Println(err)
	}

	fmt.Println(resp.Header)

	bodyBytes := streamBody(resp.Body)

	defer resp.Body.Close()

	fmt.Println(string(bodyBytes))
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
