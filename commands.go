package main

import (
	"flag"
	"spm/models"
)

var queryCmd = flag.NewFlagSet("query", flag.ExitOnError)
var queryFieldArg = queryCmd.String("l", "label", "specify label to search in")
var storeCmd = flag.NewFlagSet("store", flag.ExitOnError)
var deleteCmd = flag.NewFlagSet("delete", flag.ExitOnError)
var deleteArgId = deleteCmd.Int("id", 0, "specify a record identifier")
var generateCmd = flag.NewFlagSet("generate", flag.ExitOnError)
var sc models.StoreCommand
var gc models.GenerateCommand
