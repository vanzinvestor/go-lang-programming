package main

import "html/template"

var (
	tmpls             *template.Template
	htmlTemplateFiles = []string{
		"welcome.html",
		"home.html",
	}
)

func init() {
	// load templates from files
	tmpls = template.Must(template.ParseFiles(htmlTemplateFiles...))
}
