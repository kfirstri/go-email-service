package models

import (
	"strconv"
)

// Email represents email record from emails.csv
type Email struct {
	ID              int
	Type            string
	RecipientType   string
	RecipientID     int
	SenderAddress   string
	SubjectTemplate string
	BodyTemplate    string
	Data            string
}

// NewEmailFromRecord takes a record read from a csv files and
// Returns a pointer to the new created Email object
func NewEmailFromRecord(record []string) *Email {
	var currEmail Email

	// Parse email id and recipient id to int
	emailid, _ := strconv.Atoi(record[0])
	recipientid, _ := strconv.Atoi(record[3])

	// Fill rest of the struct
	currEmail.ID = emailid
	currEmail.Type = record[1]
	currEmail.RecipientType = record[2]
	currEmail.RecipientID = recipientid
	currEmail.SenderAddress = record[4]
	currEmail.SubjectTemplate = record[5]
	currEmail.BodyTemplate = record[6]
	currEmail.Data = record[7]

	return &currEmail
}

// GetRecipients gets the user map and returns the real recipients of the email
func (currEmail *Email) GetRecipients(users *map[int]*User, groups *map[int][]int) []*User {
	var userIDs = make([]int, 0)
	var Recipients = make([]*User, 0)

	// Get the user IDs to check
	if currEmail.RecipientType == "group_mail" {
		userIDs = append(userIDs, (*groups)[currEmail.RecipientID]...)
	} else if currEmail.RecipientType == "direct" {
		userIDs = append(userIDs, currEmail.RecipientID)
	}

	// Filter according to preferences
	for _, userID := range userIDs {
		currUser,exists := (*users)[userID]

		if exists {
			isEnabled, exists := currUser.Preferences[currEmail.Type]

			if isEnabled || !exists {
				Recipients = append(Recipients, currUser)
			}
		}
	}

	return Recipients
}