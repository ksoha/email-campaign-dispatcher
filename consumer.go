package main

import "fmt"

func emailWorker(id int, ch chan Recipient) {

	//listen to the channel
	for recipient := range ch {
		fmt.Println(id, recipient)
	}

}
