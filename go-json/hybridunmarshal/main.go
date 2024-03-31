package main

import (
	"encoding/json"
	"fmt"
)

// START TYPE OMIT
type LogRecord struct {
	Service string         `json:"service"`
	Time    int            `json:"time"`
	Data    map[string]any `json:"-"`
}

func (r *LogRecord) UnmarshalJSON(data []byte) error {
	var v struct {
		LogRecord
		UnmarshalJSON struct{} // shadow the UnmarshalJSON method
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	var d map[string]any
	_ = json.Unmarshal(data, &d)
	delete(d, "service")
	delete(d, "time")

	*r = v.LogRecord
	r.Data = d
	return nil
}

// STOP TYPE OMIT

func main() {
	// START OMIT
	data := []byte(`[
		{"time": 1711871705, "service": "direct-billing", "log": "hi", "success": true},
		{"time": 1711871728, "service": "direct-order", "message": "hello"},
		{"time": 1711871733, "service": "direct-shipper", "module": "http", "ip": "192.168.1.1"}
	]`)
	var logs []LogRecord
	if err := json.Unmarshal(data, &logs); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", logs)
	// STOP OMIT
}
