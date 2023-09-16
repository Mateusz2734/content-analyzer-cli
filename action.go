package main

import (
	"io"
	"os"

	"github.com/andrew-d/go-termutil"
	"github.com/urfave/cli/v2"
)

func mainAction(cCtx *cli.Context) error {
	letterCounter := &LetterCount{pretty: true}
	wordFreq := &WordFrequency{pretty: true}

	if termutil.Isatty(os.Stdin.Fd()) {
		return cli.Exit("Usage: analyzer [global options]\nTry 'analyzer -h' for more information.", 0)
	}

	stdin, err := io.ReadAll(os.Stdin)

	if err != nil {
		panic("Error: cannot read from stdin")
	}

	letterCounter.data, wordFreq.data = Analyze(stdin)

	letterCounter.calcSum()
	letterCounter.render()

	wordFreq.render()
	return nil
}
