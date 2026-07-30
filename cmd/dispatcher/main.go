package main

import (
	"log"
	"net/http"

	"github.com/ksoha/email-dispatcher/internal/database"
	"github.com/ksoha/email-dispatcher/internal/handlers"
)

func main() {

	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//create router
	router := http.NewServeMux()

	//register the handler for the /recipients endpoint
	router.HandleFunc(
		"GET /recipients",
		handlers.GetRecipientsHandler(db),
	)

	// Start server
	log.Println("Server running on :8080")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
