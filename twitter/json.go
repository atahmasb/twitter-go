package twitter

import (
	"encoding/json"
	"io"
)

// UnmarshalJSON reads a stream and unmarshals the results in object v.
func UnmarshalJSON(v interface{}, stream io.Reader) error {

	decoder := json.NewDecoder(stream)
	decoder.UseNumber()
	err := decoder.Decode(&v)
	if err == io.EOF {
		return nil
	} else if err != nil {
		return err
	}

	return nil
}
