package database

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"spm/models"
	"spm/security"
	"strings"
	"time"
)

const DB_NAME = "spm.db"

var userHomeDir string
var MasterPassword string

func init() {
	var err error
	userHomeDir, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
}

func filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, v := range ss {
		if test(v) {
			ret = append(ret, v)
		}
	}
	return
}

func fetchAllRecords() ([]models.Entry, int) {

	var result models.Entries
	databaseFileName, err := os.Open(userHomeDir + "/" + DB_NAME)
	if err != nil {
		databaseFileName = CreateNewDatabase()
	}
	defer databaseFileName.Close()

	byteValue, err := io.ReadAll(databaseFileName)
	if err != nil && err != io.EOF {
		panic(err)
	}

	if len(byteValue) > 0 {
		//decrypt content with master password
		byteValue, err = security.Decrypt(byteValue, MasterPassword)
		if err != nil {
			fmt.Println("cannot decrypt database, wrong master password ?")
			os.Exit(0)
		}
		json.Unmarshal(byteValue, &result)
	}
	size := len(result.Entries)

	return filter(result.Entries, func(entry models.Entry) bool {
		return !entry.Deleted
	}), size
}

func CreateNewDatabase() *os.File {
	newFile, err := os.Create(userHomeDir + "/" + DB_NAME)
	if err != nil {
		log.Fatal(err)
	}

	Memorize("spm", "master password", MasterPassword)
	return newFile
}

func ChangeMasterPassword(newPassword string) {
	content, _ := fetchAllRecords()

	var entries = models.Entries{Entries: content}
	file, _ := json.MarshalIndent(entries, "", " ")
	var err error
	file, err = security.Encrypt(file, newPassword)
	if err != nil {
		log.Fatal(err)
	}
	_ = os.WriteFile(userHomeDir+"/"+DB_NAME, file, 0644)
}

func Memorize(label, account, password string) {

	dbArray, totalRecord := fetchAllRecords()
	db := dbArray[:]
	var founded bool = false

	//cerco eventuale record esistente
	for i, v := range db {
		if v.Label == label {
			db[i].ModifiedAt = time.Now()
			db[i].Password = password
			db[i].Account = account
			founded = true
		}
	}

	if !founded {
		entry := models.Entry{Id: totalRecord + 1, Label: label, Account: account, Password: password}
		entry.CreatedAt = time.Now()
		db = append(db, entry)
	}

	cryptAndSave(db)
}

func Search(label string) []models.Entry {
	var result = make([]models.Entry, 0)
	entries, _ := fetchAllRecords()

	for _, v := range entries {
		if label == "*" || (len(label) != 0 && strings.Contains(v.Label, label)) {
			result = append(result, v)
		}
	}

	return result
}

func ExistDB() bool {
	_, error := os.Stat(userHomeDir + "/" + DB_NAME)
	return !os.IsNotExist(error)
}

func Delete(id *int) {
	dbArray, _ := fetchAllRecords()
	db := dbArray[:]

	//cerco eventuale record esistente
	for i, v := range db {
		if v.Id == *id {
			db[i].Deleted = true
		}
	}
	cryptAndSave(db)
}

func cryptAndSave(db []models.Entry) {
	var entries = models.Entries{Entries: db}
	file, _ := json.MarshalIndent(entries, "", " ")
	var err error
	file, err = security.Encrypt(file, MasterPassword[:])
	if err != nil {
		log.Fatal(err)
	}
	_ = os.WriteFile(userHomeDir+"/"+DB_NAME, file, 0644)
}
