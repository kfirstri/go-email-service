package models

import (
	"strconv"
)

// User model with data fields
type User struct {
	ID      int
	Firstname   string
	Lastname    string
	Email       string
	Company     string
	Preferences map[string]bool
}

// NewUserFromRecord takes a record read from a csv files and
// Returns a pointer to the new created User object or an error
func NewUserFromRecord(record []string) (*User, error) {
	var currUser User

	// Parse user id to int
	userid, err := strconv.Atoi(record[0])

	if err != nil {
		return nil, err
	}

	// Fill rest of the struct
	currUser.ID = userid
	currUser.Firstname = record[1]
	currUser.Lastname = record[2]
	currUser.Email = record[3]
	currUser.Company = record[4]
	currUser.Preferences = make(map[string]bool)

	return &currUser, nil
}
