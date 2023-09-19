package main

import "github.com/urfave/cli/v2"

var nFlag = &cli.Int64Flag{
	Aliases:     []string{"n"},
	Value:       15,
	Usage:       "output only `n` most frequent words, 0 means all words",
	Destination: &opts.topKWords,
}

var prettyFlag = &cli.BoolFlag{
	Name:        "pretty",
	Aliases:     []string{"p"},
	Value:       false,
	Usage:       "apply pretty print",
	Destination: &opts.prettyPrint,
}

var inFileFlag = &cli.StringFlag{
	Name:        "file",
	Aliases:     []string{"f"},
	Value:       "",
	Usage:       "read from the given `FILE`",
	Destination: &opts.inFile,
}

var outFileFlag = &cli.StringFlag{
	Name:        "out",
	Aliases:     []string{"o"},
	Value:       "",
	Usage:       "output to the given `FILE`",
	Destination: &opts.outFile,
}

var percentageFlag = &cli.BoolFlag{
	Name:        "letter-percentage",
	Aliases:     []string{"l"},
	Value:       false,
	Usage:       "enhance output with letter percentage statistics",
	Destination: &opts.letterPercentage,
}
