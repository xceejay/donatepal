package services

import (
	"bytes"
	"html/template"

	log "github.com/sirupsen/logrus"
)

type TemplateEngine struct{}

func (templateEngine TemplateEngine) ProcessFile(fileName string, vars interface{}) string {
	tmpl, err := template.ParseFiles(fileName)

	if err != nil {
		log.Error(err)

	}

	return templateEngine.process(tmpl, vars)
}

func (templateEngine TemplateEngine) process(t *template.Template, vars interface{}) string {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, vars)
	if err != nil {
		log.Error(err)
	}
	return tmplBytes.String()
}
