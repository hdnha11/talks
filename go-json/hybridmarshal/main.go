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

func (r LogRecord) MarshalJSON() ([]byte, error) {
	type copyLogRecord LogRecord // another trick to remove MarshalJSON method
	record, err := json.Marshal(copyLogRecord(r))
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(r.Data)
	if err != nil {
		return nil, err
	}

	record[len(record)-1] = ','
	return append(record, data[1:]...), nil
}

// STOP TYPE OMIT

func (r *LogRecord) UnmarshalJSON(data []byte) error {
	type copyLogRecord LogRecord
	var v copyLogRecord
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	var d map[string]any
	_ = json.Unmarshal(data, &d)
	delete(d, "service")
	delete(d, "time")

	*r = LogRecord(v)
	r.Data = d
	return nil
}

func main() {
	// START OMIT
	logs := []LogRecord{
		{Service: "direct-billing", Time: 1711871705, Data: map[string]any{
			"log":     "hi",
			"success": true}},
		{Service: "direct-order", Time: 1711871728, Data: map[string]any{
			"message": "hello"}},
		{Service: "direct-shipper", Time: 1711871733, Data: map[string]any{
			"ip":     "192.168.1.1",
			"module": "http"}},
	}

	data, err := json.MarshalIndent(logs, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data)
	// STOP OMIT
}
