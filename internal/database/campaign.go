package database

import (
	"database/sql"

	"github.com/ksoha/email-dispatcher/internal/models"
)

func CreateCampaign(db *sql.DB, campaign models.CreateCampaignRequest) error {

	query := `
	INSERT INTO campaigns (
	user_id, name, subject, body
	) 
	VALUES ($1 , $2, $3, $4)
	`
	_, err := db.Exec(
		query,
		campaign.UserID, //will not be used after jwt
		campaign.Name,
		campaign.Subject,
		campaign.Body,
	)
	if err != nil {
		return err
	}

	return nil

}

func GetCampaigns(db *sql.DB) ([]models.Campaign, error) {

	query := `
		SELECT
			id,
			user_id,
			name,
			subject,
			body,
			created_at,
			updated_at
		FROM campaigns
		ORDER BY created_at DESC
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []models.Campaign

	for rows.Next() {

		var campaign models.Campaign

		err := rows.Scan(
			&campaign.ID,
			&campaign.UserID,
			&campaign.Name,
			&campaign.Subject,
			&campaign.Body,
			&campaign.CreatedAt,
			&campaign.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		campaigns = append(campaigns, campaign)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return campaigns, nil
}
