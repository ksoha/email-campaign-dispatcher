package main

import (
	"log"
	"net/http"

	"github.com/ksoha/email-dispatcher/internal/database"
	"github.com/ksoha/email-dispatcher/internal/handlers"
	"github.com/ksoha/email-dispatcher/internal/middleware"
)

func main() {

	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//create router
	router := http.NewServeMux()

	//route POST/signup
	router.HandleFunc(
		"POST /signup",
		handlers.SignUpHandler(db),
	)

	//route GET/Recipients
	router.HandleFunc(
		"GET /recipients",
		handlers.GetRecipientsHandler(db),
	)

	//route POST/Recipients/Import
	router.Handle(
		"POST /recipients/import",
		middleware.AuthMiddleware(
			handlers.CreateRecipientHandler(db),
		),
	)

	//route POST/login
	router.HandleFunc(
		"POST /login",
		handlers.LoginHandler(db),
	)

	router.HandleFunc(
		"POST /campaigns",
		handlers.CreateCampaignHandler(db),
	)

	router.Handle(
		"GET /campaigns",
		middleware.AuthMiddleware(
			handlers.GetCampaignsHandler(db),
		),
	)

	// route POST /campaigns/{id}/send
	router.HandleFunc(
		"POST /campaigns/{id}/send",
		handlers.SendCampaignHandler(db),
	)

	// Start server
	log.Println("Server running on :8080")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
