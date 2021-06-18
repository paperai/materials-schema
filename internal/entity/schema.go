package entity

import (
	"bytes"
	"io"
	"strings"
	"text/template"
)

type Schema struct {
	ID         string      `json:"id,omitempty"`
	Name       string      `json:"name,omitempty"`
	Synonyms   []string    `json:"synonyms,omitempty"`
	Properties []*Property `json:"properties,omitempty"`
	SchemaName string      `json:"schema,omitempty"`
}

type Property struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Color string `json:"color,omitempty"`
	Sort  uint16 `json:"sort,omitempty"`
}

func (s *Schema) HTML(templ *template.Template) (io.Reader, error) {
	type outputProperty struct {
		ID         string
		Name       string
		Type       string
		Color      string
		Sort       uint16
		Synonyms   string
		SchemaName string
	}

	outputs := make([]*outputProperty, len(s.Properties))
	for i, property := range s.Properties {
		outputs[i] = &outputProperty{
			ID:         s.ID,
			Name:       property.Name,
			SchemaName: s.SchemaName,
			Type:       property.Type,
			Color:      property.Color,
			Sort:       property.Sort,
			Synonyms:   strings.Join(s.Synonyms, ","),
		}
	}

	var buf bytes.Buffer
	err := templ.Execute(&buf, struct {
		Schema     string
		SchemaName string
		Synonyms   string
		Properties []*outputProperty
	}{
		Schema:     s.Name,
		Properties: outputs,
	})
	return &buf, err
}
