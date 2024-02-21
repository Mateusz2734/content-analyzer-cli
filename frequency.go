package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type WordFrequency struct {
	data []Word
}

func (obj *WordFrequency) print(builder *strings.Builder) {
	builder.WriteString("word,count\n")
	for _, word := range obj.data {
		builder.WriteString(fmt.Sprintf("%s,%d\n", word.text, word.count))
	}
}

func (obj *WordFrequency) prettyPrint(builder *strings.Builder) {
	table := tablewriter.NewWriter(builder)

	table.SetHeader([]string{"Word", "Count"})
	table.SetAlignment(tablewriter.ALIGN_CENTER)

	for _, word := range obj.data {
		table.Append([]string{word.text, strconv.Itoa(word.count)})
	}

	table.Render()
}

func (obj *WordFrequency) render(builder *strings.Builder, opts *Options) {
	if opts.prettyPrint {
		obj.prettyPrint(builder)
	} else {
		obj.print(builder)
	}
}
