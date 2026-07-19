package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

//producer is what will read the data
//making the producer a part of main package

// will recieve the data file in this function
func loadRecipients(filePath string, ch chan Recipient) error {
	//logic to read the file
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer f.Close()

	//reading the csv data
	r := csv.NewReader(f)
	record, err := r.ReadAll()
	if err != nil {
		return err
	}

	//looping through the records
	for _, record := range record[1:] {
		fmt.Println(record)

		//send -> consumer
		//implementing a queue tyoe structure using channels
		//using an unbuffered channel(capacity of 0)

		//sending the instance of Recipient struct into tthe channel
		ch <- Recipient{
			Name:  record[0], //name
			Email: record[1],
		} //blocking

	}

	return nil
}
