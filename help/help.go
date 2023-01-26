package help

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func DisplayUsage() {

	fmt.Printf("Invalid usage, please read the following instructions ...\n\n")
	headerFmt := color.New(color.FgRed, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgCyan).SprintfFunc()

	tbl := table.New("COMMAND", "PARAMETER", "EXAMPLE VALUE", "DESCRIPTION")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	tbl.AddRow("query", "-l <<label>>", "-l google.com", "search <<label>> into record ")
	tbl.AddRow("all", "", "", "return all record")
	tbl.AddRow("store", "-l <<label>> -a <<account>> -p <<password>>", "-l google.com -a account@gmail.com -p pippo", "create new entry")
	tbl.AddRow("reset", "", "", "change master password with new one asked in terminal")
	tbl.AddRow("generate", "-l <<length>> -s <<special chars>>", "-l 20 -s 2", "generate 20 chars password with 2 special char")
	tbl.Print()
}
