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
	switch err {
	case io.EOF:
		return nil
	case io.ErrUnexpectedEOF:
		return nil
	default:
		return err
	}
}
