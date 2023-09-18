package main

import (
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v2"
)

func output(content string) {
	if opts.outFile == "" {
		stdin_output(content)
	} else {
		file_output(content)
	}
}

func stdin_output(content string) {
	fmt.Print(content)
}

func file_output(content string) {
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
		return stdin_input()
	} else {
		return file_input()
	}
}

func stdin_input() []byte {
	stdin, err := io.ReadAll(os.Stdin)

	if err != nil {
		cli.Exit("Cannot read from stdin", 1)
	}

	return stdin
}

func file_input() []byte {
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
