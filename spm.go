package main

import (
	"fmt"
	"os"
	"sort"
	"spm/database"
	"spm/help"
	"spm/models"
	"syscall"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"golang.org/x/term"
)

func inizializeDatabase() {
	fmt.Print("please, specify a master password for your new database:")
	var err error
	var mPasswordSlice = make([]byte, 32, 32)
	mPasswordSlice, err = term.ReadPassword(syscall.Stdin)
	if err != nil {
		os.Exit(0)
	}
	database.MasterPassword = string(mPasswordSlice)
	database.CreateNewDatabase()
	fmt.Println("This is your password, please take a note and don't forget it, or you cannot be able to access your data\n", database.MasterPassword)
}

func readPassword() {
	fmt.Print("enter master password:")
	var err error
	var mPasswordSlice = make([]byte, 32)
	mPasswordSlice, err = term.ReadPassword(syscall.Stdin)
	if err != nil {
		os.Exit(0)
	}
	database.MasterPassword = string(mPasswordSlice)
}

func main() {

	var existDB bool
	if len(os.Args) < 2 {
		help.DisplayUsage()
		os.Exit(0)
	}

	existDB = database.ExistDB()
	if existDB {
		if os.Args[1] != "generate" {
			readPassword()
		}
	} else {
		inizializeDatabase()
	}

	fmt.Println()

	if v, ok := Tasks[os.Args[1]]; ok {
		v.(func([]string))(os.Args)
	} else {
		fmt.Println("Unknow command : ", os.Args[1])
		help.DisplayUsage()
	}
}

func displayResult(result []models.Entry) {

	sort.Slice(result, func(i, j int) bool {
		return result[i].Label < result[j].Label
	})

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgHiYellow).SprintfFunc()

	tbl := table.New("PASSWORD", "LABEL", "ACCOUNT", "ID", "CREATED AT", "MODIFIED AT")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, widget := range result {
		if widget.ModifiedAt.IsZero() {
			tbl.AddRow(widget.Password, widget.Label, widget.Account, widget.Id, widget.CreatedAt.Format("2006-01-02 15:04:05"), "-")
		} else {
			tbl.AddRow(widget.Password, widget.Label, widget.Account, widget.Id, widget.CreatedAt.Format("2006-01-02 15:04:05"), widget.ModifiedAt.Format("2006-01-02 15:04:05"))
		}

	}

	tbl.Print()
	if len(result) == 0 {
		fmt.Println("no records found")
	}
}
