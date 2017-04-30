package main // import "github.com/leanlabsio/blacksmith"

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/leanlabsio/blacksmith/cli"
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
		cli.DaemonCmd,
	}
	app.Version = Version
	app.Run(os.Args)
}
