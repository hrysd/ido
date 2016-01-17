package main

import (
	"net/http"
	"net/url"
	"os"

	"github.com/codegangsta/cli"
	"github.com/hrysd/ido/internal/pipe"
)

func main() {
	app := cli.NewApp()
	app.Name = "ido"
	app.Usage = "Not yet"
	app.Commands = Commands

	app.Action = func(c *cli.Context) {
		hookName := c.Args().Get(0)
		hook := DetectHook(hookName)

		post(hook.Token, pipe.Read())
	}

	app.Run(os.Args)
}

func post(endpoint string, content string) {
	values := url.Values{}
	values.Add("source", content)

	_, err := http.PostForm(endpoint, values)

	if err != nil {
		panic(err) // XXX
	}
}
