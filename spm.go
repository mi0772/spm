package main

import (
	"flag"
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

func main() {

	var existDB bool

	if len(os.Args) < 2 {
		help.DisplayUsage()
		os.Exit(0)
	}

	existDB = database.ExistDB()
	if existDB {
		fmt.Print("enter current master password: ")
	} else {
		fmt.Println("please, specify a master password for your new database: ")
	}

	var err error
	var mPasswordSlice = make([]byte, 32, 32)
	mPasswordSlice, err = term.ReadPassword(syscall.Stdin)
	if err != nil {
		os.Exit(0)
	}
	database.MasterPassword = string(mPasswordSlice)

	if !existDB {
		fmt.Println("This is your password, please take a note and don't forget it, or you cannot be able to access your data\n", database.MasterPassword)
	}
	fmt.Println()
	queryCmd := flag.NewFlagSet("query", flag.ExitOnError)
	queryFieldArg := queryCmd.String("l", "label", "specify label to search in")

	storeCmd := flag.NewFlagSet("store", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	deleteArgId := deleteCmd.Int("id", 0, "specify a record identifier")
	var sc models.StoreCommand
	sc.Label = storeCmd.String("l", "", "Label for new password")
	sc.Account = storeCmd.String("a", "", "Account for this password")
	sc.Password = storeCmd.String("p", "", "The password")

	switch os.Args[1] {
	case "all":
		displayResult(database.Search("*"))
	case "query":
		queryCmd.Parse(os.Args[2:])
		fmt.Println("serching for label : ", *queryFieldArg)
		displayResult(database.Search(*queryFieldArg))
	case "reset":
		fmt.Print("Enter New Master Password: ")
		var newPasswordSlice = make([]byte, 32, 32)
		newPasswordSlice, err = term.ReadPassword(syscall.Stdin)
		if err != nil {
			os.Exit(0)
		}
		fmt.Println()
		database.ChangeMasterPassword(string(newPasswordSlice))
		fmt.Println("master password sucessfully changed")
	case "store":
		v := os.Args[2:]
		storeCmd.Parse(v)
		if storeCmd.NFlag() < 3 {
			fmt.Println("devi specificare 3 parametri")
			os.Exit(1)
		}
		fmt.Println("memorize new password for : ", *sc.Label, *sc.Account, *sc.Password)
		database.Memorize(*sc.Label, *sc.Account, *sc.Password)
	case "delete":
		v := os.Args[2:]
		deleteCmd.Parse(v)
		database.Delete(deleteArgId)
	case "version":
		fmt.Println("rp - simple password manager, v1.0.1 (C) Carlo Di Giuseppe, 16-01-2023")
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
