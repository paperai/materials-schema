package entity

import (
	"bytes"
	"io"
	"strings"
	"text/template"
)

type Schema struct {
	ID         string   `json:"id,omitempty"`
	Name       string   `json:"name,omitempty"`
	Synonyms   []string `json:"synonyms,omitempty"`
	Properties []string `json:"property_ids,omitempty"`
	SchemaName string   `json:"schema,omitempty"`
}

// 今は使わないが後ほど使う可能性があるので取っておく
/*
type Property struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Color string `json:"color,omitempty"`
	Sort  uint16 `json:"sort,omitempty"`
}
*/

type outputSchema struct {
	ID         string
	Name       string
	Synonyms   string
	Properties string
	SchemaName string
}

func (s *Schema) HTML(templ *template.Template) (io.Reader, error) {
	var buf bytes.Buffer

	err := templ.Execute(&buf, struct {
		Schema *outputSchema
	}{
		Schema: &outputSchema{
			ID:         s.ID,
			Name:       s.Name,
			Synonyms:   strings.Join(s.Synonyms, ","),
			Properties: strings.Join(s.Properties, ","),
			SchemaName: s.SchemaName,
		},
	})
	return &buf, err
}
