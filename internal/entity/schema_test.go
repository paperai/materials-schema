package entity_test

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"text/template"

	"github.com/paperai/materials-schema/internal/conv"
)

func Test_HTML(t *testing.T) {
	f, err := os.Open("../../testdata/test_schema.json")
	if err != nil {
		t.Errorf("fail to open json file: %v", err)
		return
	}
	defer f.Close()
	j, err := ioutil.ReadAll(f)
	if err != nil {
		t.Errorf("fail to read json: %v", err)
		return
	}
	es, err := conv.JSONToStruct(j)
	if err != nil {
		t.Errorf("fail to convert json: %v", err)
		return
	}

	templ, err := template.ParseFiles("../../templ/index.html")
	if err != nil {
		t.Errorf("fail to parse html: %v", err)
		return
	}

	schema := es[0]

	reader, err := schema.HTML(templ)
	if err != nil {
		t.Errorf("fail to encode html: %v", err)
		return
	}

	var htmlBytes []byte
	for {
		buf := make([]byte, 4096)
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error(err)
			return
		}
		if n > 0 {
			htmlBytes = append(htmlBytes, buf[:n]...)
		}
	}
	htmlStr := string(htmlBytes)

	if result := schema.Name; !strings.Contains(htmlStr, result) {
		t.Errorf("%s word didn't contain", result)
		return
	}
	if result := schema.Properties[0].Name; !strings.Contains(htmlStr, result) {
		t.Errorf("%s word didn't contain", result)
		return
	}
	if result := schema.Properties[0].Color; !strings.Contains(htmlStr, result) {
		t.Errorf("%s word didn't contain", result)
		return
	}
	if result := schema.Properties[0].Type; !strings.Contains(htmlStr, result) {
		t.Errorf("%s word didn't contain", result)
		return
	}
	if result := schema.Synonyms[0]; !strings.Contains(htmlStr, result) {
		t.Errorf("%s word didn't contain", result)
		return
	}
	if result := schema.SchemaName; !strings.Contains(htmlStr, result) {
		t.Errorf("%s word didn't contain", result)
		return
	}
	if result := "存在しない適当なワード"; strings.Contains(htmlStr, result) {
		t.Errorf("%s word contain", result)
		return
	}
}
