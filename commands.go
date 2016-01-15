package main

import (
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
)

const FILENAME = ".ido"

var Commands = []cli.Command{
	setCommand,
}

var setCommand = cli.Command{
	Name:        "set",
	Usage:       "Store hook name and token",
	Description: "Not yet",
	Action:      doSet,
}

func doSet(c *cli.Context) {
	name := c.Args().Get(0)
	token := c.Args().Get(1)

	storeTokenWithName(name, token)
}

func storeTokenWithName(name string, token string) {
	homeDirectory := os.Getenv("HOME")
	ioutil.WriteFile(".ido")
}

func configurationPath() string {
	return filepath.Join(os.Getenv("HOME"), FILENAME)
}
