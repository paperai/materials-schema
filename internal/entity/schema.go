package entity

import (
	"bytes"
	"io"
	"text/template"
)

type Schema struct {
	ID         string      `json:"id,omitempty"`
	Name       string      `json:"name,omitempty"`
	Properties []*Property `json:"properties,omitempty"`
}

type Property struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Color string `json:"color,omitempty"`
	Sort  uint16 `json:"sort,omitempty"`
}

func (s *Schema) HTML(templ *template.Template) (io.Reader, error) {
	var buf bytes.Buffer
	err := templ.Execute(&buf, struct {
		Schema     string
		Properties []*Property
	}{
		Schema:     s.Name,
		Properties: s.Properties,
	})
	return &buf, err
}
