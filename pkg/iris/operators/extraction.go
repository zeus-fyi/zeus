package iris_operators

import (
	"reflect"
	"strings"
)

type IrisRoutingResponseETL struct {
	Source        string `json:"source"`
	ExtractionKey string `json:"extractionKey"`
	DataType      string `json:"dataType"`
	Value         any    `json:"result"`
}

func (r *IrisRoutingResponseETL) ExtractKeyValue(m map[string]interface{}) {
	if r.ExtractionKey == "" {
		r.Value = m
		return
	}

	// Splitting the key on comma to navigate the nested structure
	keys := strings.Split(r.ExtractionKey, ",")

	var current interface{} = m
	for _, key := range keys {
		// Type assert current to a map
		if asMap, ok := current.(map[string]interface{}); ok {
			// If the key exists, update current. Otherwise, set r.Value to nil and return
			if value, exists := asMap[key]; exists {
				current = value
			} else {
				r.Value = nil
				return
			}
		} else {
			r.Value = nil
			return
		}
	}

	r.Value = current
	if r.Value == nil {
		return
	}
	r.DataType = reflect.TypeOf(r.Value).String()
}
