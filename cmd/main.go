package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/paperai/materials-schema/internal/conv"
	"golang.org/x/xerrors"
)

const (
	schemaFileName = "schema.md"
)

func main() {
	t, err := template.ParseFiles("./templ/index.html")
	if err != nil {
		log.Fatal(xerrors.Errorf("fail to parse html. err: %w", err))
	}

	f, err := os.Open("./schemas.json")
	if err != nil {
		log.Fatalf("fail to open json file: %v", err)
	}
	defer f.Close()

	j, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("fail to read json: %v", err)
	}

	schemas, err := conv.JSONToStruct(j)
	if err != nil {
		log.Fatalf("fail to convert json: %v", err)
	}

	materialHTMLStrs := make([]string, len(schemas))
	for i, schema := range schemas {
		reader, err := schema.HTML(t)
		if err != nil {
			log.Fatalf("fail to encode html: %v", err)
		}

		var htmlBytes []byte
		for {
			buf := make([]byte, 4096)
			n, err := reader.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
				return
			}
			if n > 0 {
				htmlBytes = append(htmlBytes, buf[:n]...)
			}
		}
		materialHTMLStrs[i] = string(htmlBytes)
	}

	if err := writeToREADME(materialHTMLStrs); err != nil {
		log.Fatal(err)
	}
}

func writeToREADME(htmlStrs []string) error {
	file, err := os.Create(schemaFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString("#material-schema\n"); err != nil {
		return err
	}

	if _, err := file.WriteString(strings.Join(htmlStrs, "<br>")); err != nil {
		return err
	}

	return nil
}
