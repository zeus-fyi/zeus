package iris_operators

import (
	"errors"
	"fmt"
	"reflect"
)

// adding a name will generate this header: fmt.Sprintf("X-Agg-Max-Value-%s", v.Name)
// etc agg operation. only max, gt, gte are supported for now for this feature

type Aggregation struct {
	Name       string                   `json:"name,omitempty"`
	Operator   AggOp                    `json:"operator"`
	Comparison *Operation               `json:"comparison,omitempty"`
	DataType   string                   `json:"dataType"`
	DataSlice  []IrisRoutingResponseETL `json:"dataSlice"`
	//WindowFilter any    `json:"windowFilter,omitempty"`

	SumInt            int     `json:"sumInt,omitempty"`
	CurrentMinInt     int     `json:"currentMinInt,omitempty"`
	CurrentMaxInt     int     `json:"currentMaxInt,omitempty"`
	CurrentMaxFloat64 float64 `json:"CurrentMaxFloat64,omitempty"`
}

type AggOp string

func (a AggOp) Max() string {
	return Max
}

func (a AggOp) Sum() string {
	return Sum
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

// aggregation comparison

const (
	Max = "max"
	Sum = "sum"
)

// AggregateOn order of priority, and will only execute the first valid operator: comparison operator || agg operator
func (a *Aggregation) AggregateOn(x any, y IrisRoutingResponseETL) error {
	if a.Comparison != nil && a.Comparison.X != nil {
		switch a.Comparison.DataTypeX {
		case DataTypeInt:
			// for comparison ops: the x operator gets converted inside the comparison agg functions
			// this is most often a pre-set stored procedure value, but can also be dynamically set
			switch a.Comparison.Operator {
			case Gt:
				return a.AggregateGtInt(y)
			case Gte:
				return a.AggregateGtEqInt(y)
			case Lt:
				return a.AggregateLtInt(y)
			case Lte:
				return a.AggregateLtEqInt(y)
			case Eq:
				return a.AggregateEqInt(y)
			case Neq:
				return a.AggregateNeqInt(y)
			}
		}
	}

	switch string(a.Operator) + a.DataType {
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
			a.CurrentMinInt = x
			a.DataSlice = []IrisRoutingResponseETL{y} // Keep only the new maximum value
		} else {
			a.DataSlice = append(a.DataSlice, y) // Append the value if it's equal to the current maximum
		}
	}
	return nil
}
