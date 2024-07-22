package util

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// OutputFormat represents the available output formats.
type OutputFormat string

const (
	TableOutput OutputFormat = "table"
	RawOutput   OutputFormat = "raw"
	JSONOutput  OutputFormat = "json"
)

var tableAlignment = tablewriter.ALIGN_CENTER
var tableCenterSeparator = "|"
var tableColumnSeparator = "|"
var tableRowSeparator = "-"
var tableHeaderAlignment = tablewriter.ALIGN_CENTER
var tableHeaderLine = true
var tableBorder = true
var tableRowLine = true
var tableLeftBorder = true
var tableTopBorder = true
var tableRightBorder = true
var tableBottomBorder = false
var tablePadding = "   "
var tableNoWhiteSpace = false

func TableWriter(data [][]string, cols ...string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(cols)
	table.AppendBulk(data)
	table.SetAlignment(tableAlignment)
	table.SetCenterSeparator(tableCenterSeparator)
	table.SetColumnSeparator(tableColumnSeparator)
	table.SetRowSeparator(tableRowSeparator)
	table.SetHeaderAlignment(tableHeaderAlignment)
	table.SetHeaderLine(tableHeaderLine)
	table.SetBorder(tableBorder)
	table.SetRowLine(tableRowLine)
	table.SetBorders(tablewriter.Border{Left: tableLeftBorder, Top: tableTopBorder, Right: tableRightBorder, Bottom: tableBottomBorder})
	table.SetTablePadding(tablePadding)
	table.SetNoWhiteSpace(tableNoWhiteSpace)

	table.Render()
}

func RenderOutput(data [][]string, cols []string, format OutputFormat) {
	switch format {
	case TableOutput:
		TableWriter(data, cols...)
	case RawOutput:
		for _, output := range data {
			fmt.Println(strings.Join(output, ", "))
		}
	case JSONOutput:
		jsonData := make(map[string]string)
		for _, output := range data {
			jsonData[output[0]] = output[1]
		}
		jsonOutput, err := json.MarshalIndent(jsonData, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling to JSON:", err)
			return
		}
		fmt.Println(string(jsonOutput))
	}
}
