package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
)

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		return []byte{}
	}

	return buff.Bytes()
}

// String returns string pointer
func String(data string) *string {
	return &data
}

// StringUnref returns empty string if argument is nil, otherwise it returns string value.
func StringUnref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// Int creates reference from given data value.
func Int(data int) *int {
	return &data
}

// IntUnref returns int pointer referenced value, in case pointer is nil, it returns default int value 0.
func IntUnref(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

// Int64 creates reference from given data value.
func Int64(data int64) *int64 {
	return &data
}

// Bool creates reference from given data value.
func Bool(data bool) *bool {
	return &data
}

// BoolUnref returns false if argument is nil, otherwise it returns bool value.
func BoolUnref(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// Float64 creates reference from given data value.
func Float64(data float64) *float64 {
	return &data
}

// PrintJSON marshals value to json and returns result. If error occurs it is being skipped.
func PrintJSON(data interface{}) string {
	js, _ := json.Marshal(data)
	return string(js)
}
