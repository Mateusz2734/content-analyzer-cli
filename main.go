package main

import (
	"io"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		HelpName: "content-analyzer",
		Usage:    "Analyze word frequency and letter count of a file piped from stdin",
		Action: func(cCtx *cli.Context) error {
			letterCounter := &LetterCount{pretty: true}
			wordFreq := &WordFrequency{pretty: true}

			stdin, err := io.ReadAll(os.Stdin)

			if err != nil {
				panic("Error: cannot get stdin")
			}

			letterCounter.data, wordFreq.data = Analyze(stdin)

			letterCounter.calcSum()
			letterCounter.render()

			wordFreq.render()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
