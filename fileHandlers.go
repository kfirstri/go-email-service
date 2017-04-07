package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"strconv"

	"github.com/kfirstri/go-email-service/models"
)

// ReadFile reads the csv file specified in filename
func ReadFile(filename string, handleRecord func([]string)) error {
	// Open the file
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		return err
	}

	// Create a new CSV reader
	csvReader := csv.NewReader(bufio.NewReader(file))

	// Start reading
	csvRecord, err := csvReader.Read()

	// Loop until eof
	for err != io.EOF {

		// Make sure we don't some different error
		if err != nil {
			return err
		}

		handleRecord(csvRecord)
		csvRecord, err = csvReader.Read()
	}

	return nil
}

// handleUsersFile gets a record, create a new User struct and saves a
// pointer to it in the Users map
func handleUsersFile(record []string) {
	user, err := models.NewUserFromRecord(record)

	if err != nil {
		fmt.Printf("Error in creating new User: %v", err)
	}

	// Add the new user to the Users map
	Users[user.ID] = user
}

// handleGroupsFile gets a group record, create a new record in the Groups maps
// and adds the userID to that group
func handleGroupsFile(record []string) {
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
}

// handlePreferencesFile gets a preferences record and adds it to the User's preferences map
func handlePreferencesFile(record []string) {
	userID, _ := strconv.Atoi(record[0])
	emailType := record[1]
	isEnabled, _ := strconv.ParseBool(record[2])

	Users[userID].Preferences[emailType] = isEnabled
}

// handleEmailsFile handles emails
func handleEmailsFile(record []string) {
	var currEmail = models.NewEmailFromRecord(record)

	// Get the recipients
	recipients := currEmail.GetRecipients(&Users, &Groups)

	// If there are not recipients don't send the email
	if len(recipients) == 0 {
		return
	}

	// Parse templates

	// Send Emails
	sendEmail(currEmail, recipients, "", "")
}
