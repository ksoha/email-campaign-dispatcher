package mail

import (
	"bytes"
	"html/template"
)

func executeTemplate(r Recipient) (string, error) {
	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer

	if err := t.Execute(&tpl, r); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
