package entity

import (
	"bytes"
	"io"
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

func (s *Schema) HTML(templ *template.Template) (io.Reader, error) {
	var buf bytes.Buffer
	err := templ.Execute(&buf, struct {
		Schema  string
		Schemas []*Schema
	}{
		Schema:  s.Name,
		Schemas: s,
	})
	return &buf, err
}
