package main

import (
	"fmt"
	"strconv"

	"github.com/kfirstri/go-email-service/models"
)

// Consts for all files and formats
const usersFile = "users.csv"
const groupsFile = "group_members.csv"
const preferencesFile = "user_email_preferences.csv"
const emailsFile = "emails.csv"
const outputFolder = "sent_emails/"
const emailSubjectFormat = "%v_%v_subject.txt"
const emailBodyFormat = "%v_%v_body.html"

func main() {
	// Initialize users and groups
	var Users = make(map[int]*models.User)
	var Groups = make(map[int][]int)
	var err error

	// Create a new User for each record in usersFile
	err = ReadFile(usersFile, func(record []string) {
		user, err := models.NewUserFromRecord(record)

		if err != nil {
			fmt.Printf("Error in creating new User: %v", err)
		}

		// Add the new user to the Users map
		Users[user.UserID] = user
	})

	if err != nil {
		fmt.Printf("Error in reading users file: %v", err)
	}

	for userID, user := range Users {
		fmt.Printf("User ID %v, User = %+v\n", userID, user)
	}

	// Fill the Groups map
	err = ReadFile(groupsFile, func(record []string) {
		groupID, _ := strconv.Atoi(record[0])
		userID, _ := strconv.Atoi(record[1])

		_, exists := Groups[groupID]

		// If the group's users array doesn't exists we need to allocate it
		if exists {
			Groups[groupID] = append(Groups[groupID], userID)
		} else {
			Groups[groupID] = make([]int, 1)
			Groups[groupID][0] = userID
		}
	})

	if err != nil {
		fmt.Printf("Error in reading groups file: %v", err)
	}

}
