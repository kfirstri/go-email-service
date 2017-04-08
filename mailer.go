package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"strings"
	"text/template"

	"fmt"
	"os"

	"github.com/kfirstri/go-email-service/models"
)

// parseDataAndTemplate gets the JSON and both file names,
// return a JSON map, and the data templates as a map
// with the keys "subject" and "body"
func parseDataAndTemplate(data, subjectFile, bodyFile string) (map[string]interface{}, map[string]*template.Template) {
	var templates = make(map[string]*template.Template)

	// Parse data
	var dataJSON = make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &dataJSON)

	if err != nil {
		log.Printf("Error parsing JSON: %v", err)
	}

	// Read subject template file
	content, err := ioutil.ReadFile(subjectFile)

	if err != nil {
		log.Printf("Error parsing file %v: %v", subjectFile, err)
	}

	templates["subject"] = makeTemplate(string(content[:len(content)]))

	// Read subject template file
	content, err = ioutil.ReadFile(bodyFile)

	if err != nil {
		log.Printf("Error parsing file %v: %v", bodyFile, err)
	}

	templates["body"] = makeTemplate(string(content[:len(content)]))

	return dataJSON, templates
}

func makeTemplate(text string) *template.Template {
	// handle files without dot
	text = strings.Replace(text, "{{ ", "{{ .", -1)

	t := template.Must(template.New("email_body").Parse(text))

	return t
}

func sendEmail(email *models.Email, recipients []*models.User, data map[string]interface{}, templates map[string]*template.Template) {
	// Execute the templates for each recipient
	for _, rec := range recipients {
		// Set current recipient data
		data["first_name"] = rec.Firstname
		data["last_name"] = rec.Lastname
		data["email_address"] = rec.Email

		// Get file's names
		subjectFileName := fmt.Sprintf(emailSubjectFormat, email.ID, rec.ID)
		bodyFileName := fmt.Sprintf(emailBodyFormat, email.ID, rec.ID)

		// Write both files
		fakeEmailFile(subjectFileName, data, templates["subject"])
		fakeEmailFile(bodyFileName, data, templates["body"])

		log.Printf("%v,%v", email.ID, rec.ID)
		addToLog(fmt.Sprintf("%v,%v", email.ID, rec.ID))
	}
}

func fakeEmailFile(fileName string, data map[string]interface{}, template *template.Template) {
	fs, err := os.Create(outputFolder + fileName)

	if err != nil {
		fmt.Printf("err creating file '%v': %v", fileName, err)
		return
	}

	err = template.Execute(fs, data)

	if err != nil {
		fmt.Printf("err writing file '%v': %v", fileName, err)
	}

	fs.Close()
}
