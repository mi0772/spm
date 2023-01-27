package main

import (
	"flag"
	"spm/models"
	"strconv"
)

var queryCmd = flag.NewFlagSet("query", flag.ExitOnError)
var queryFieldArg = queryCmd.String("l", "label", "specify label to search in")
var storeCmd = flag.NewFlagSet("store", flag.ExitOnError)
var deleteCmd = flag.NewFlagSet("delete", flag.ExitOnError)
var deleteArgId = deleteCmd.Int("id", 0, "specify a record identifier")
var generateCmd = flag.NewFlagSet("generate", flag.ExitOnError)
var sc models.StoreCommand
var gc models.GenerateCommand

func GetParameterAsInt(args []string, index uint) (res int, ok bool) {
	if checkLength(args, index) {
		return 0, false
	}

	res, err := strconv.Atoi(args[index])
	if err != nil {
		return 0, false
	}
	return res, true
}

func GetParameterAsString(args []string, index uint) (res string, ok bool) {
	if checkLength(args, index) {
		return "", false
	}
	res = args[index]
	return res, true
}

func checkLength(args []string, index uint) bool {
	return uint(len(args)-1) < index
}
