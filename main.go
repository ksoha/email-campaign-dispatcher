package main

import (
	"bytes"
	"html/template"
	"sync"
)

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

	//using wait group to wait for the all the goroutines to finish their work
	var wg sync.WaitGroup

	//creating different workers to consume the data
	workerCount := 5

	for i := 0; i < workerCount; i++ {
		wg.Add(1) //adding a worker to the wait group
		go emailWorker(i, recipient, &wg)
	}

	wg.Wait() //waiting for all the workers to finish their work

}

// function to execute the template
func executeTemplate(r Recipient) (string, error) {
	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer //buffer to hold the executed template

	//passing the recipient , because its a struct which has access to the fields(name in template)
	e := t.Execute(&tpl, r) //executing the template with the recipient data
	if e != nil {
		return "", e
	}

	return tpl.String(), nil //returning the executed template as a string

}
