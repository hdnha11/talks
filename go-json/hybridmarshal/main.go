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
	var v = struct {
		LogRecord
		MarshalJSON struct{} `json:"-"`
	}{LogRecord: r}
	record, err := json.Marshal(v)
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
