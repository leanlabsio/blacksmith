package main // import "github.com/leanlabsio/blacksmith"

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/leanlabsio/blacksmith/cmd"
)

var Version string = "dev"

func main() {
	app := cli.NewApp()
	app.Name = "blacksmithci"
	app.Authors = []cli.Author{
		{
			Name:  "cnam",
			Email: "support@leanlabs.io",
		},
		{
			Name:  "V",
			Email: "support@leanlabs.io",
		},
	}
	app.Commands = []cli.Command{
		cmd.DaemonCmd,
	}
	app.Version = Version
	app.Run(os.Args)
}
