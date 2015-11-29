package main

import (
	"github.com/codegangsta/cli"
	"github.com/vasiliy-t/blacksmithci/cmd"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "blacksmithci"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "cnam",
			Email: "support@leanlabs.io",
		},
		cli.Author{
			Name:  "V",
			Email: "support@leanlabs.io",
		},
	}
	app.Commands = []cli.Command{
		cmd.DaemonCmd,
	}
	app.Version = "0.0.1"
	app.Run(os.Args)
}
