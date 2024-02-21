package main

import "github.com/urfave/cli/v2"

func nFlag(opts *Options) *cli.Int64Flag {
	return &cli.Int64Flag{
		Aliases:     []string{"n"},
		Value:       15,
		Usage:       "output only `n` most frequent words, 0 means all words",
		Destination: &opts.topKWords,
	}
}

func prettyFlag(opts *Options) *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:        "pretty",
		Aliases:     []string{"p"},
		Value:       false,
		Usage:       "apply pretty print",
		Destination: &opts.prettyPrint,
	}
}

func inFileFlag(opts *Options) *cli.StringFlag {
	return &cli.StringFlag{
		Name:        "file",
		Aliases:     []string{"f"},
		Value:       "",
		Usage:       "read from the given `FILE`",
		Destination: &opts.inFile,
	}

}

func outFileFlag(opts *Options) *cli.StringFlag {
	return &cli.StringFlag{
		Name:        "out",
		Aliases:     []string{"o"},
		Value:       "",
		Usage:       "output to the given `FILE`",
		Destination: &opts.outFile,
	}

}

func percentageFlag(opts *Options) *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:        "letter-percentage",
		Aliases:     []string{"l"},
		Value:       false,
		Usage:       "enhance output with letter percentage statistics",
		Destination: &opts.letterPercentage,
	}

}
