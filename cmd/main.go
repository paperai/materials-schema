package main

import (
	"log"
	"text/template"

	"golang.org/x/xerrors"
)

func main() {
	t, err := template.ParseFiles("../templ/index.html")
	if err != nil {
		log.Fatal(xerrors.Errorf("fail to parse html. err: %w", err))
	}

	log.Println(t)
}
