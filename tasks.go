package main

import (
	"fmt"
	"log"
	"os"
	"spm/database"
	"spm/help"
	"strconv"
	"syscall"

	"github.com/sethvargo/go-password/password"
	"golang.org/x/term"
)

var Tasks map[string]interface{} = make(map[string]interface{})

func init() {
	Tasks["all"] = func(args []string) {
		displayResult(database.Search("*"))
	}
	Tasks["query"] = func(args []string) {
		queryCmd.Parse(os.Args[2:])
		fmt.Println("serching for label : ", *queryFieldArg)
		displayResult(database.Search(*queryFieldArg))
	}
	Tasks["reset"] = func(args []string) {
		fmt.Print("Enter New Master Password: ")
		var newPasswordSlice = make([]byte, 32)
		newPasswordSlice, err := term.ReadPassword(syscall.Stdin)
		if err != nil {
			os.Exit(0)
		}
		fmt.Println()
		database.ChangeMasterPassword(string(newPasswordSlice))
		fmt.Println("master password sucessfully changed")
	}
	Tasks["store"] = func(args []string) {
		v := os.Args[2:]
		sc.Label = storeCmd.String("l", "", "Label for new password")
		sc.Account = storeCmd.String("a", "", "Account for this password")
		sc.Password = storeCmd.String("p", "", "The password")

		storeCmd.Parse(v)
		if storeCmd.NFlag() < 3 {
			help.DisplayUsage()
			os.Exit(1)
		}

		fmt.Println("memorize new password for : ", *sc.Label, *sc.Account, *sc.Password)
		database.Memorize(*sc.Label, *sc.Account, *sc.Password)
	}

	Tasks["delete"] = func(args []string) {
		v := os.Args[2:]
		deleteCmd.Parse(v)
		database.Delete(deleteArgId)
	}

	Tasks["version"] = func(args []string) {
		fmt.Println("rp - simple password manager, v1.0.1 (C) Carlo Di Giuseppe, 16-01-2023")
	}

	Tasks["generate"] = func(args []string) {

		v := os.Args[2:]
		gc.Length = generateCmd.String("l", "10", "Length of password")
		gc.Special = generateCmd.String("s", "1", "Number of special chars")
		generateCmd.Parse(v)
		if generateCmd.NFlag() < 2 {
			help.DisplayUsage()
			os.Exit(1)
		}
		v1, _ := strconv.Atoi(*gc.Length)
		v2, _ := strconv.Atoi(*gc.Special)

		fmt.Printf("generate %v password length with %v special chars\n", *gc.Length, *gc.Special)
		res, err := password.Generate(v1, 5, v2, false, false)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("This is your password:", res)
	}
}
