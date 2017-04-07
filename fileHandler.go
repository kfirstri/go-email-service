package main

import (
	"os"
	"bufio"
	"encoding/csv"
	"io"
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