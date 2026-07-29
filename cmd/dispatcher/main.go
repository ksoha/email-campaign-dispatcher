package main

import (
	"sync"

	"github.com/ksoha/email-dispatcher/internal/mail"
)

func main() {

	//cerating channel
	//not giving size inside make to make it unbuffered
	recipient := make(chan mail.Recipient)

	//running producer and consumer inside the main thread/function will cause a deadlock
	//running them in separate go rountines
	go func() {
		mail.LoadRecipients("emails.csv", recipient)
	}()

	//using wait group to wait for the all the goroutines to finish their work
	var wg sync.WaitGroup

	//creating different workers to consume the data
	workerCount := 5

	for i := 0; i < workerCount; i++ {
		wg.Add(1) //adding a worker to the wait group
		go mail.EmailWorker(i, recipient, &wg)
	}

	wg.Wait() //waiting for all the workers to finish their work

}
