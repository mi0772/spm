package models

import "time"

type StoreCommand struct {
	Label    *string
	Account  *string
	Password *string
}

type GenerateCommand struct {
	Length  *string
	Special *string
}

type Entry struct {
	Id         int       `json:"id"`
	Label      string    `json:"label"`
	Account    string    `json:"account"`
	Password   string    `json:"password"`
	Hint       string    `json:"hint,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at,omitempty"`
	Deleted    bool      `json:"deleted,omitempty"`
}

type Entries struct {
	Entries []Entry `json:"entries"`
}
