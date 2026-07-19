package main

import "time"

// struct of a receipient
type Recipient struct {
	Name  string
	Email string
}

func main() {

	//cerating channel
	//not giving size inside make to make it unbuffered
	recipient := make(chan Recipient)

	//running producer and consumer inside the main thread/function will cause a deadlock
	//running them in separate go rountines
	go func() {
		loadRecipients("emails.csv", recipient)
	}()

	go emailWorker(1, recipient)

	time.Sleep(3 * time.Second)
}
