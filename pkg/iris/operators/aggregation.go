package iris_operators

import (
	"errors"
	"fmt"
	"reflect"
)

type Aggregation struct {
	Operator  string                   `json:"operator"`
	DataType  string                   `json:"dataType"`
	DataSlice []IrisRoutingResponseETL `json:"dataSlice"`
	//WindowFilter any    `json:"windowFilter,omitempty"`

	SumInt            int     `json:"sumInt,omitempty"`
	CurrentMaxInt     int     `json:"currentMaxInt,omitempty"`
	CurrentMaxFloat64 float64 `json:"CurrentMaxFloat64,omitempty"`
}
type IrisRoutingResponseETL struct {
	Source        string `json:"source"`
	ExtractionKey string `json:"extractionKey"`
	DataType      string `json:"dataType"`
	Value         any    `json:"result"`
}

func (r *IrisRoutingResponseETL) ExtractKeyValue(m map[string]any) {
	if r.ExtractionKey == "" {
		r.Value = m
		return
	}
	r.Value = m[r.ExtractionKey]
	if r.Value == nil {
		return
	}
	r.DataType = reflect.TypeOf(r.Value).String()
}

const (
	Max = "max"
	Sum = "sum"
)

func (a *Aggregation) AggregateOn(x any, y IrisRoutingResponseETL) error {
	switch a.Operator + a.DataType {
	case Max + DataTypeInt:
		val, ok := ConvertToInt(x)
		if !ok {
			return errors.New(fmt.Sprintf("could not convert %v to int", x))
		}
		return a.AggregateMaxInt(val, y)
	case Sum + DataTypeInt:
		val, ok := ConvertToInt(x)
		if !ok {
			return errors.New(fmt.Sprintf("could not convert %v to int", x))
		}
		return a.AggregateSumInt(val, y)
	case Max + DataTypeFloat64:
		val, ok := ConvertToFloat64(x)
		if !ok {
			return errors.New(fmt.Sprintf("could not convert %v to float64", x))
		}
		return a.AggregateMaxFloat64(val, y)
	default:
		return errors.New(fmt.Sprintf("could not aggregate on %s", a.DataType))
	}
}

func (a *Aggregation) AggregateSumInt(x int, y IrisRoutingResponseETL) error {
	a.Operator = Sum
	a.SumInt += x
	a.DataSlice = append(a.DataSlice, y) // Append the value if it's equal to the current maximum
	return nil
}

func (a *Aggregation) AggregateMaxFloat64(x float64, y IrisRoutingResponseETL) error {
	a.Operator = Max
	if len(a.DataSlice) == 0 || x >= a.CurrentMaxFloat64 {
		if x > a.CurrentMaxFloat64 {
			a.CurrentMaxFloat64 = x
			a.DataSlice = []IrisRoutingResponseETL{y} // Keep only the new maximum value
		} else {
			a.DataSlice = append(a.DataSlice, y) // Append the value if it's equal to the current maximum
		}
	}
	return nil
}

func (a *Aggregation) AggregateMaxInt(x int, y IrisRoutingResponseETL) error {
	a.Operator = Max
	if len(a.DataSlice) == 0 || x >= a.CurrentMaxInt {
		if x > a.CurrentMaxInt {
			a.CurrentMaxInt = x
			a.DataSlice = []IrisRoutingResponseETL{y} // Keep only the new maximum value
		} else {
			a.DataSlice = append(a.DataSlice, y) // Append the value if it's equal to the current maximum
		}
	}
	return nil
}
