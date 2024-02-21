package main

import (
	"os"
	"strings"

	"github.com/andrew-d/go-termutil"
	"github.com/urfave/cli/v2"
)

func mainAction(opts *Options) func(c *cli.Context) error {
	return func(cCtx *cli.Context) error {
		sb := &strings.Builder{}

		letterCounter := &LetterCount{}
		wordFreq := &WordFrequency{}

		if termutil.Isatty(os.Stdin.Fd()) && opts.inFile == "" {
			return cli.Exit("Usage: ca [global options]\nTry 'ca -h' for more information.", 0)
		}

		data := input(opts)

		letterCounter.data, wordFreq.data = analyze(data, opts)

		letterCounter.render(sb, opts)
		wordFreq.render(sb, opts)

		output(sb.String(), opts)

		return nil
	}
}
