package examples

import (
	"bytes"
	"encoding/json"
)

// StructToMap returns a map where keys are a struct non empty fields and their values
func StructToMap(item interface{}) (map[string]interface{}, error) {
	var data map[string]interface{}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(item)
	err := json.Unmarshal(reqBodyBytes.Bytes(), &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
