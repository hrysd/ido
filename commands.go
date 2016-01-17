package main

import (
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	setCommand,
}

var setCommand = cli.Command{
	Name:        "set-hook",
	Usage:       "Store hook name and token",
	Description: "Not yet",
	Action:      doSet,
}

func doSet(c *cli.Context) {
	panic("Not implemented yet")
}