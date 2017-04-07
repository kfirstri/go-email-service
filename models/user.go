package models

import (
	"strconv"
)

// User model with data fields
type User struct {
	UserID      int
	Firstname   string
	Lastname    string
	Email       string
	Company     string
	Preferences map[string]bool
}

func makeUser(args ...string) (*User, error) {
	var currUser User

	// Parse user id to int
	userid, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, err
	}

	// Fill rest of the struct
	currUser.UserID = userid
	currUser.Firstname = args[1]
	currUser.Lastname = args[2]
	currUser.Email = args[3]
	currUser.Company = args[4]
	currUser.Preferences = make(map[string]bool)

	return &currUser, nil
}

// NewUserFromRecord takes a record read from a csv files and
// Returns a pointer to the new created User object or an error
func NewUserFromRecord(record []string) (*User, error) {
	return makeUser(record...)
}
