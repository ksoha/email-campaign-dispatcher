package mail

import (
	"sync"

	"github.com/ksoha/email-dispatcher/internal/models"
)

func SendCampaign(campaign models.Campaign, recipients []models.Recipient) error {

	const workerCount = 5

	//channel through which workers receive the recipients
	ch := make(chan models.Recipient)

	var wg sync.WaitGroup //WaitGroup to wait for all the workers to finish

	//start the workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)

		//start a worker goroutine
		go EmailWorker(
			i,
			ch,
			campaign,
			&wg,
		)
	}
	// Send recipients into the channel
	for _, recipient := range recipients {
		ch <- recipient
	}

	// No more recipients will be sent
	close(ch)

	// Wait until all workers finish
	wg.Wait()

	return nil
}
