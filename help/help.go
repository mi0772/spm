package help

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func DisplayUsage() {

	fmt.Printf("Invalid usage, please read the following instructions ...\n\n")
	headerFmt := color.New(color.FgRed, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgCyan).SprintfFunc()

	tbl := table.New("COMMAND", "PARAMETER", "EXAMPLE VALUE", "DESCRIPTION")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	tbl.AddRow("query", "<<label>>", "spm query google.com", "search google.com into record ")
	tbl.AddRow("all", "", "", "return all record")
	tbl.AddRow("store", "<<label>> <<account>> <<password>>", "google.com account@gmail.com pippo", "create new entry")
	tbl.AddRow("reset", "", "", "change master password with new one asked in terminal")
	tbl.AddRow("generate", "<<length>> <<special chars>>", "20 2", "generate 20 chars password with 2 special char")
	tbl.Print()
	os.Exit(0)
}
