package mail

import (
	"bytes"
	"html/template"

	"github.com/ksoha/email-dispatcher/internal/models"
)

// Data that will be available inside email.tmpl
type TemplateData struct {
	Recipient models.Recipient
	Campaign  models.Campaign
}

// function to execute the template
func executeTemplate(
	recipient models.Recipient,
	campaign models.Campaign,
) (string, error) {

	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer

	data := TemplateData{
		Recipient: recipient,
		Campaign:  campaign,
	}

	err = t.Execute(&tpl, data)
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}
