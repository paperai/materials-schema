package conv

import (
	"encoding/json"

	"github.com/paperai/materials-schema/internal/entity"
)

func JSONToStruct(jsonStr []byte) ([]*entity.Schema, error) {
	st := new([]*entity.Schema)
	if err := json.Unmarshal([]byte(jsonStr), st); err != nil {
		return nil, err
	}

	return *st, nil
}
