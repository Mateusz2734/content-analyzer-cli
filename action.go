package main

import (
	"os"
	"strings"

	"github.com/andrew-d/go-termutil"
	"github.com/urfave/cli/v2"
)

func mainAction(cCtx *cli.Context) error {
	sb := &strings.Builder{}

	letterCounter := &LetterCount{}
	wordFreq := &WordFrequency{}

	if termutil.Isatty(os.Stdin.Fd()) && opts.inFile == "" {
		return cli.Exit("Usage: ca [global options]\nTry 'ca -h' for more information.", 0)
	}

	data := input()

	letterCounter.data, wordFreq.data = analyze(data)

	letterCounter.render(sb)
	wordFreq.render(sb)

	output(sb.String())

	return nil
}
