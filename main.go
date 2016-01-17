package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ido"
	app.Usage = "Not yet"
    app.Commands = Commands
	
    app.Action = func(c *cli.Context) {
        hookName := c.Args().Get(0)
		hook := DetectHook(hookName)

        post(hook.Token, scanStdout())
	}

	app.Run(os.Args)
}

func scanStdout() string {
	var lines bytes.Buffer
	scanner := bufio.NewScanner(os.Stdin)

	lines.Write([]byte("```\n"))

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		lines.Write(scanner.Bytes())
		lines.Write([]byte("\n"))
	}

	lines.Write([]byte("\n```"))

	return lines.String()
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
