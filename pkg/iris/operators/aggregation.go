package iris_operators

import (
	"errors"
	"fmt"
)

type Aggregation struct {
	Operator  string `json:"operator"`
	DataType  string `json:"dataType"`
	DataSlice []any  `json:"dataSlice"`
	//WindowFilter any    `json:"windowFilter,omitempty"`

	CurrentMaxInt     int     `json:"currentMaxInt,omitempty"`
	CurrentMaxFloat64 float64 `json:"CurrentMaxFloat64,omitempty"`
}

const (
	Max = "max"
)

func (a *Aggregation) AggregateOn(x any, y any) error {
	switch a.Operator + a.DataType {
	case Max + DataTypeInt:
		val, ok := ConvertToInt(x)
		if !ok {
			return errors.New(fmt.Sprintf("could not convert %v to int", x))
		}
		return a.AggregateMaxInt(val, y)
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

func (a *Aggregation) AggregateMaxFloat64(x float64, y any) error {
	a.Operator = Max
	if len(a.DataSlice) == 0 || x >= a.CurrentMaxFloat64 {
		if x > a.CurrentMaxFloat64 {
			a.CurrentMaxFloat64 = x
			a.DataSlice = []any{y} // Keep only the new maximum value
		} else {
			a.DataSlice = append(a.DataSlice, y) // Append the value if it's equal to the current maximum
		}
	}
	return nil
}

func (a *Aggregation) AggregateMaxInt(x int, y any) error {
	a.Operator = Max
	if len(a.DataSlice) == 0 || x >= a.CurrentMaxInt {
		if x > a.CurrentMaxInt {
			a.CurrentMaxInt = x
			a.DataSlice = []any{y} // Keep only the new maximum value
		} else {
			a.DataSlice = append(a.DataSlice, y) // Append the value if it's equal to the current maximum
		}
	}
	return nil
}

//func (a *Aggregation) AggregateMaxHexstr(x string) error {
//	return nil
//}
