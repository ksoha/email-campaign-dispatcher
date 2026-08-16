package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"log"

	"github.com/ksoha/email-dispatcher/internal/database"
	"github.com/ksoha/email-dispatcher/internal/mail"
	"github.com/ksoha/email-dispatcher/internal/models"
	"github.com/ksoha/email-dispatcher/internal/response"
)

func CreateCampaignHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var campaign models.CreateCampaignRequest

		err := json.NewDecoder(r.Body).Decode(&campaign)
		if err != nil {
			response.WriteError(
				w,
				http.StatusBadRequest,
				"Invalid request",
			)
			return
		}

		err = database.CreateCampaign(db, campaign)
		if err != nil {
			log.Println("CreateCampaign error:", err)
			response.WriteError(
				w,
				http.StatusInternalServerError,
				"Failes tp create campaign",
			)
			return
		}

		//Return success response
		err = response.WriteJSON(
			w,
			http.StatusAccepted,
			map[string]string{
				"message": "Campaign created successfully",
			},
		)

		if err != nil {
			response.WriteError(
				w,
				http.StatusInternalServerError,
				"Failed to encode response",
			)
			return
		}

	}
}

func GetCampaignsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//fetch all campaigns from the database
		campaigns, err := database.GetCampaigns(db)
		if err != nil {
			response.WriteError(
				w,
				http.StatusInternalServerError,
				"Failed to retrieve campaigns",
			)
			return
		}

		err = response.WriteJSON(
			w,
			http.StatusAccepted,
			campaigns,
		)
		if err != nil {
			response.WriteError(
				w,
				http.StatusInternalServerError,
				"Failed to fetch campaigns",
			)
			return
		}
	}
}

// function that handles sending for a small batch of recipients
func SendCampaignHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println(" SEND CAMPAIGN HANDLER HIT")

		// Get campaign ID from URL
		campaignID := r.PathValue("id")

		log.Println("Campaign ID from URL:", campaignID)
		// Fetch the selected campaign
		campaign, err := database.GetCampaignByID(db, campaignID)
		if err != nil {

			log.Println("GetCampaignByID error:", err)

			response.WriteError(
				w,
				http.StatusNotFound,
				"Campaign not found",
			)
			return
		}

		// Fetch recipients from database
		recipients, err := database.GetRecipient(db)
		if err != nil {
			response.WriteError(
				w,
				http.StatusInternalServerError,
				"Failed to fetch recipients",
			)
			return
		}

		// Dispatch campaign to recipients
		err = mail.SendCampaign(
			campaign,
			recipients,
		)
		if err != nil {
			response.WriteError(
				w,
				http.StatusInternalServerError,
				"Failed to send campaign",
			)
			return
		}

		// Success
		err = response.WriteJSON(
			w,
			http.StatusOK,
			map[string]string{
				"message": "Campaign sent successfully",
			},
		)
		if err != nil {
			response.WriteError(
				w,
				http.StatusInternalServerError,
				"Failed to encode response",
			)
			return
		}
	}
}
