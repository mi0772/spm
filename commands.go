package main

import (
	"strconv"
)

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
