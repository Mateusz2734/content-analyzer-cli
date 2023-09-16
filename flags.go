package main

import "github.com/urfave/cli/v2"

var nFlag = &cli.Int64Flag{
	Aliases:     []string{"n"},
	Value:       0,
	Usage:       "Output only `n` most frequent words, 0 means all words",
	Destination: &opts.topKWords,
}

var prettyFlag = &cli.BoolFlag{
	Name:        "pretty",
	Aliases:     []string{"p"},
	Value:       false,
	Usage:       "Apply pretty print",
	Destination: &opts.prettyPrint,
}

var inFileFlag = &cli.StringFlag{
	Name:        "file",
	Aliases:     []string{"f"},
	Value:       "",
	Usage:       "Read from given `FILE`",
	Destination: &opts.inFile,
}

var outFileFlag = &cli.StringFlag{
	Name:        "out",
	Aliases:     []string{"o"},
	Value:       "",
	Usage:       "Output to a given `FILE`",
	Destination: &opts.outFile,
}

var percentageFlag = &cli.BoolFlag{
	Name:        "letter-percentage",
	Aliases:     []string{"l"},
	Value:       false,
	Usage:       "Enrich output with letter percentage stats",
	Destination: &opts.letterPercentage,
}
