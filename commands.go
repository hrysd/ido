package main

import (
	"bytes"
	"fmt"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	listCommand,
}

var listCommand = cli.Command{
	Name:        "list",
	Usage:       "Show all stored hook",
	Description: "Show all stored hook in `~/.ido`",
	Action:      list,
}

func list(c *cli.Context) {
	hooks := AllHooks()

	var lines bytes.Buffer

	for i, h := range hooks {
		lines.Write([]byte(fmt.Sprintf("%v: %v", h.Name, h.Token)))

		if i < (len(hooks) - 1) {
			lines.Write([]byte("\n"))
		}
	}

	fmt.Println(lines.String())
}
