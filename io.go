package main

import (
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v2"
)

func output(content string) {
	if opts.outFile == "" {
		stdinOutput(content)
	} else {
		fileOutput(content)
	}
}

func stdinOutput(content string) {
	fmt.Print(content)
}

func fileOutput(content string) {
	f, err := os.OpenFile(opts.outFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)

	if err != nil {
		cli.Exit("Cannot open given file", 1)
	}

	_, err = f.WriteString(content)

	if err != nil {
		cli.Exit("Cannot write to the file", 1)
	}

	if err := f.Close(); err != nil {
		cli.Exit("Cannot close file", 1)
	}
}

func input() []byte {
	if opts.inFile == "" {
		return stdinInput()
	} else {
		return fileInput()
	}
}

func stdinInput() []byte {
	content, err := io.ReadAll(os.Stdin)

	if err != nil {
		cli.Exit("Cannot read from stdin", 1)
	}

	return content
}

func fileInput() []byte {
	fileInfo, err := os.Stat(opts.inFile)

	if err != nil {
		cli.Exit("File does not exist", 1)
	}

	if fileInfo.IsDir() {
		cli.Exit("Cannot read from direcory", 1)
	}

	file, err := os.OpenFile(opts.inFile, os.O_RDONLY, 0666)

	if err != nil {
		cli.Exit("Cannot open given file", 1)
	}

	content, err := io.ReadAll(file)

	if err != nil {
		cli.Exit("Cannot read from given file", 1)
	}

	return content
}
