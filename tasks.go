package main

import (
	"fmt"
	"log"
	"os"
	"spm/database"
	"spm/help"
	"spm/models"
	"syscall"

	"github.com/sethvargo/go-password/password"
	"golang.org/x/term"
)

var Tasks map[string]interface{} = make(map[string]interface{})

func init() {
	Tasks["all"] = all
	Tasks["query"] = query
	Tasks["reset"] = reset
	Tasks["store"] = memorize
	Tasks["delete"] = delete
	Tasks["version"] = version
	Tasks["generate"] = generateTask
}

func all(args []string) {
	displayResult(database.Search("*"))
}

func query(args []string) {
	var p string
	var ok bool
	if p, ok = GetParameterAsString(args, 2); !ok {
		help.DisplayUsage()
	}
	fmt.Println("serching for label : ", p)
	displayResult(database.Search(p))
}

func reset(args []string) {
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

func memorize(args []string) {

	var label, account, password string
	var ok bool

	if label, ok = GetParameterAsString(args, 2); !ok {
		help.DisplayUsage()
	}
	if account, ok = GetParameterAsString(args, 3); !ok {
		help.DisplayUsage()
	}
	if password, ok = GetParameterAsString(args, 4); !ok {
		help.DisplayUsage()
	}
	var sc = &models.StoreCommand{Label: &label, Account: &account, Password: &password}

	fmt.Printf("memorize new password for label:%s, account:%s, password:%s\n", *sc.Label, *sc.Account, *sc.Password)
	database.Memorize(*sc.Label, *sc.Account, *sc.Password)
}

func delete(args []string) {
	var id int
	var ok bool
	if id, ok = GetParameterAsInt(args, 2); !ok {
		help.DisplayUsage()
	}

	database.Delete(&id)
}

func version() {
	fmt.Println("rp - simple password manager, v1.0.1 (C) Carlo Di Giuseppe, 16-01-2023")
}

func generateTask(args []string) {
	var length, numspecial int
	var ok bool
	if length, ok = GetParameterAsInt(args, 2); !ok {
		help.DisplayUsage()
	}
	if numspecial, ok = GetParameterAsInt(args, 3); !ok {
		help.DisplayUsage()
	}
	var gc = &models.GenerateCommand{Length: length, Special: numspecial}

	fmt.Printf("generate %v password length with %v special chars\n", gc.Length, gc.Special)
	res, err := password.Generate(gc.Length, gc.Length/4, gc.Special, false, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("This is your password:", res)
}
