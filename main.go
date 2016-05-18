package main // import "github.com/leanlabsio/blacksmith"

import (
	"github.com/codegangsta/cli"
	"github.com/leanlabsio/blacksmith/cmd"
	"os"
)

var Version string = "dev"

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
	app.Version = Version
	app.Run(os.Args)
}
