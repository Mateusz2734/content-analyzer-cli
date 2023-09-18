package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var opts = &Options{}

func main() {

	app := &cli.App{
		Name:      "ca",
		HelpName:  "ca",
		Usage:     "Analyze word frequency and letter count of a file",
		UsageText: "ca [global options]",

		Flags: []cli.Flag{
			nFlag,
			prettyFlag,
			inFileFlag,
			outFileFlag,
			percentageFlag,
		},

		Action: mainAction,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
