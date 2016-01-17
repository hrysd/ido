package main

import (
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

		hook.Post(pipe.Read())
	}

	app.Run(os.Args)
}
