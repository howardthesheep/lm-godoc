package web

import (
	"html/template"
)

var (
	PackagesTemplate *template.Template
	FilesTemplate    *template.Template
	IndexTemplate    *template.Template
)

func InitializeTemplates() error {
	var err error

	IndexTemplate, err = template.ParseFiles("./www/index.html")
	if err != nil {
		return err
	}

	PackagesTemplate, err = template.ParseFiles("./www/packages.html")
	if err != nil {
		return err
	}

	FilesTemplate, err = template.ParseFiles("./www/files.html")
	if err != nil {
		return err
	}

	return nil
}
