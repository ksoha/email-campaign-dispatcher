package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()

	//listen to the channel
	for recipient := range ch {

		//simulate sending email
		smtpHost := "localhost"
		smtpPort := "1025"

		// formattedMsg := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\r\n%s\r\n", recipient.Email,
		// 	"Test email our email cheimpaign.")

		// //using a slice of bytes to sen the email message
		// msg := []byte(formattedMsg)

		msg, err := executeTemplate(recipient)
		if err != nil {
			fmt.Printf("Worker %d failed to execute template for %s: %v\n", id, recipient.Email, err)
			continue
		}

		fmt.Printf("Worker %d sending email to %s\n", id, recipient.Email)

		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "sohakulkarni09@gmail.com", []string{recipient.Email}, []byte(msg))
		if err != nil {
			log.Fatal(err)
		}

		//simulate some delay in sending email
		time.Sleep(50 * time.Millisecond)

		fmt.Printf("Worker %d sent email to %s\n", id, recipient.Email)
	}

}
