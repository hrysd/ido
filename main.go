package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {
	var lines bytes.Buffer
	scanner := bufio.NewScanner(os.Stdin)

	lines.Write([]byte("```\n"))

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		lines.Write(scanner.Bytes())
		lines.Write([]byte("\n"))
	}

	lines.Write([]byte("\n```"))

	post(os.Args[1], lines.String())
}

func post(endpoint string, content string) {
	values := url.Values{}
	values.Add("source", content)

	response, err := http.PostForm(endpoint, values)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response)
}
