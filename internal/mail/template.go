package mail

import (
	"bytes"
	"html/template"
)

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
