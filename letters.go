package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type LetterCount struct {
	pretty bool
	data   []int
	sum    int
}

func (obj *LetterCount) calcSum() {
	obj.sum = 0
	for _, count := range obj.data {
		obj.sum += count
	}
}

func (obj *LetterCount) print() {
	for i, count := range obj.data {
		fmt.Printf("%c %d\n", i+97, count)
	}
}

func (obj *LetterCount) prettyPrint() {
	tableData := [2][]string{}

	for i, count := range obj.data {
		tableData[0] = append(tableData[0], fmt.Sprintf("%c", i+97))
		tableData[1] = append(tableData[1], fmt.Sprint(count))
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)

	for _, row := range tableData {
		table.Append(row)
	}

	table.Render()
}

func (obj *LetterCount) render() {
	if obj.pretty {
		obj.prettyPrint()
	} else {
		obj.print()
	}
}
