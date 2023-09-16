package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var opts = &Options{}

func main() {

	app := &cli.App{
		Name:      "analyzer",
		HelpName:  "analyzer",
		Usage:     "Analyze word frequency and letter count of a file piped from stdin",
		UsageText: "analyzer [global options]",

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
