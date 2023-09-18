package main

import (
	"fmt"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type LetterCount struct {
	data []int
	sum  int
}

func (obj *LetterCount) calcSum() {
	obj.sum = 0
	for _, count := range obj.data {
		obj.sum += count
	}
}

func (obj *LetterCount) print(builder *strings.Builder) {
	if opts.letterPercentage {
		obj.calcSum()

		builder.WriteString("letter,count,percentage\n")
	} else {
		builder.WriteString("letter,count\n")
	}

	for i, count := range obj.data {
		if opts.letterPercentage {
			builder.WriteString(fmt.Sprintf("%c,%d,%.2f%%\n", i+97, count, (float32(count)/float32(obj.sum))*100))
		} else {
			builder.WriteString(fmt.Sprintf("%c,%d\n", i+97, count))
		}
	}
	builder.WriteString("\n")
}

func (obj *LetterCount) prettyPrint(builder *strings.Builder) {
	var tableData [][]string

	if opts.letterPercentage {
		for i := 0; i < 3; i++ {
			tableData = append(tableData, []string{})
		}

		for i, count := range obj.data {
			tableData[0] = append(tableData[0], fmt.Sprintf("%c", i+97))
			tableData[1] = append(tableData[1], fmt.Sprint(count))
			tableData[2] = append(tableData[2], fmt.Sprintf("%.2f%%", (float32(count)/float32(obj.sum))*100))
		}

	} else {
		for i := 0; i < 2; i++ {
			tableData = append(tableData, []string{})
		}

		for i, count := range obj.data {
			tableData[0] = append(tableData[0], fmt.Sprintf("%c", i+97))
			tableData[1] = append(tableData[1], fmt.Sprint(count))
		}

	}

	table := tablewriter.NewWriter(builder)
	table.SetRowLine(true)
	table.SetAlignment(tablewriter.ALIGN_CENTER)

	for _, row := range tableData {
		table.Append(row)
	}

	table.Render()
	builder.WriteString("\n")
}

func (obj *LetterCount) render(builder *strings.Builder) {
	if opts.prettyPrint {
		obj.prettyPrint(builder)
	} else {
		obj.print(builder)
	}
}
