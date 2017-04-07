package main

import (
	"fmt"

	"log"

	"github.com/kfirstri/go-email-service/models"
)

// Consts for all files and formatsc
const usersFile = "users.csv"
const groupsFile = "group_members.csv"
const preferencesFile = "user_email_preferences.csv"
const emailsFile = "emails.csv"
const outputFolder = "sent_emails/"
const emailSubjectFormat = "%v_%v_subject.txt"
const emailBodyFormat = "%v_%v_body.html"

// Users is the map with all the server users
var Users = make(map[int]*models.User)

// Groups are the server's user group
var Groups = make(map[int][]int)

func loadUserAndGroups() error {
	var err error

	// Create a new User for each record in usersFile
	err = ReadFile(usersFile, handleUsersFile)

	if err != nil {
		return err
	}

	// Fill the Groups map
	err = ReadFile(groupsFile, handleGroupsFile)

	if err != nil {
		return err
	}

	// Read user's Preferences
	err = ReadFile(preferencesFile, handlePreferencesFile)

	if err != nil {
		return err
	}

	return nil
}

func sendEmail(email *models.Email, recipients []*models.User, subject, body string) {
	for _,rec := range recipients {
		log.Printf("%v,%v", email.ID, rec.ID)
		log.Printf(emailSubjectFormat, email.ID, rec.ID)
		log.Printf(emailBodyFormat, email.ID, rec.ID)
	}
}

func main() {
	var err error

	// Load all data
	err = loadUserAndGroups()

	if err != nil {
		panic(fmt.Sprintf("Error loading data: %v", err))
	}

	// Start going over emails and handle them.
	err = ReadFile(emailsFile, handleEmailsFile)

	if err != nil {
		panic(fmt.Sprintf("Error loading emails file: %v", err))
	}
}
