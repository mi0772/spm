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

	tbl := table.New("COMMAND", "PARAMENTER", "EXAMPLE VALUE", "DESCRIPTION")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	tbl.AddRow("query", "-l <<label>>", "-l google.com", "search <<label>> into record ")
	tbl.AddRow("all", "", "", "return all record")
	tbl.AddRow("store", "-l <<label>> -a <<account>> -p <<password>>", "-l google.com -a account@gmail.com -p pippo", "create new entry")
	tbl.AddRow("reset", "", "", "change master password with new one asked in terminal")

	tbl.Print()
}
