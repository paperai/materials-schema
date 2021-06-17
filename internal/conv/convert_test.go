package conv_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/paperai/materials-schema/internal/conv"
	"github.com/paperai/materials-schema/internal/entity"
)

func expectedJSON(j []byte) ([]*entity.Schema, error) {
	expected := new([]*entity.Schema)
	if err := json.Unmarshal(j, expected); err != nil {
		return nil, err
	}
	return *expected, nil
}

func Test_JSONToStruct(t *testing.T) {
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

	expected, err := expectedJSON(j)
	if err != nil {
		t.Error(err)
		return
	}

	f2, err := os.Open("../../testdata/test_schema2.json")
	if err != nil {
		t.Errorf("fail to open json file: %v", err)
		return
	}
	defer f2.Close()

	j2, err := ioutil.ReadAll(f2)
	if err != nil {
		t.Errorf("fail to read json: %v", err)
		return
	}

	expected2, err := expectedJSON(j2)
	if err != nil {
		t.Error(err)
		return
	}

	for _, te := range []struct {
		title    string
		input    []byte
		expected []*entity.Schema
		isError  bool
	}{
		{
			title:    "it is success for file1 to convert json",
			input:    j,
			expected: expected,
			isError:  false,
		},
		{
			title:    "it is success for file2 to convert json",
			input:    j2,
			expected: expected2,
			isError:  false,
		},
		{
			title:    "it is fail to convert json",
			input:    j,
			expected: expected2,
			isError:  true,
		},
	} {
		result, err := conv.JSONToStruct(te.input)
		if err != nil {
			t.Error(err)
			return
		}

		if diff := cmp.Diff(result, te.expected); (diff != "") != te.isError {
			t.Errorf("response was mismatch. (-result, +expected)\n%s", diff)
		}
	}

}
