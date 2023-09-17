package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type WordFrequency struct {
	data []Word
}

func (obj *WordFrequency) print() {
	fmt.Println("word,count")
	for _, word := range obj.data {
		fmt.Printf("%s,%d\n", word.text, word.count)
	}
}

func (obj *WordFrequency) prettyPrint() {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Word", "Count"})
	table.SetAlignment(tablewriter.ALIGN_CENTER)

	for _, word := range obj.data {
		table.Append([]string{word.text, strconv.Itoa(word.count)})
	}

	table.Render()
}

func (obj *WordFrequency) render() {
	if opts.prettyPrint {
		obj.prettyPrint()
	} else {
		obj.print()
	}
}
