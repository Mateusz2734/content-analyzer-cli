package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	opts := &Options{}

	app := &cli.App{
		Name:      "ca",
		HelpName:  "ca",
		Usage:     "Analyze word frequency and letter count of a file",
		UsageText: "ca [global options]",
		Version:   "1.0.0",

		Flags: []cli.Flag{
			nFlag(opts),
			prettyFlag(opts),
			inFileFlag(opts),
			outFileFlag(opts),
			percentageFlag(opts),
		},

		Action: mainAction(opts),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
