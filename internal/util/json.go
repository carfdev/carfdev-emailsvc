package util

import (
	"bytes"
	"encoding/json"
)

// StrictUnmarshal deserializes JSON while disallowing unknown fields.
func StrictUnmarshal(data []byte, v any) error {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	return dec.Decode(v)
}
