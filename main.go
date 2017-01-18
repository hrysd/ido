package main

import (
	"os"

	"github.com/codegangsta/cli"
	//"github.com/hrysd/ido/internal/pipe"
)

func main() {
	app := cli.NewApp()
	app.Name = "ido"
	app.Usage = "Not yet"
	app.Commands = Commands

	app.Action = func(c *cli.Context) {
		//hookName := c.Args().Get(0)

    room := idobata.GetRooms("esminc", "hrysd").Rooms[0]
		idobata.CreateMessage("hoge", room.id)
	}

	app.Run(os.Args)
}
