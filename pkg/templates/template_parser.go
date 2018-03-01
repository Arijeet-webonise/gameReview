package templates

import (
	"bytes"
	"html/template"
)

type ITemplateParser interface {
	ParseTemplate([]string, interface{}) (string, error)
}

func (tp *TemplateParser) ParseTemplate(templateFileName []string, data interface{}) (string, error) {
	var parsedTemplate string

	funcMap := template.FuncMap{
		"TestFunc": TestFunc,
	}

	t, err := template.New("titleTest").Funcs(funcMap).ParseFiles(templateFileName...)
	if err != nil {
		return parsedTemplate, err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return parsedTemplate, err
	}
	parsedTemplate = buf.String()
	return parsedTemplate, nil
}

type TemplateParser struct {
}

func TestFunc(num int) bool {
	if num == 0 {
		return true
	}
	return false
}
