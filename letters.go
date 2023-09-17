package main

import (
	"fmt"
	"os"

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

func (obj *LetterCount) print() {
	if opts.letterPercentage {
		fmt.Println("letter,count,percentage")
	} else {
		fmt.Println("letter,count")
	}

	for i, count := range obj.data {
		if opts.letterPercentage {
			fmt.Printf("%c,%d,%.2f%%\n", i+97, count, (float32(count)/float32(obj.sum))*100)
		} else {
			fmt.Printf("%c,%d\n", i+97, count)
		}
	}
	fmt.Println("")
}

func (obj *LetterCount) prettyPrint() {
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

	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetAlignment(tablewriter.ALIGN_CENTER)

	for _, row := range tableData {
		table.Append(row)
	}

	table.Render()
	fmt.Println("")
}

func (obj *LetterCount) render() {
	if opts.prettyPrint {
		obj.prettyPrint()
	} else {
		obj.print()
	}
}
