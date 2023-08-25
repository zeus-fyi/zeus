package iris_operators

import (
	"errors"
	"fmt"
)

func (a *Aggregation) AggregateGtInt(y IrisRoutingResponseETL) error {
	return a.aggregateCompInt(Gt, y)
}

func (a *Aggregation) AggregateGtEqInt(y IrisRoutingResponseETL) error {
	return a.aggregateCompInt(Gte, y)
}

func (a *Aggregation) aggregateCompInt(op Op, yt IrisRoutingResponseETL) error {
	a.Comparison.Operator = op
	a.Comparison.DataTypeX = DataTypeInt
	a.Comparison.DataTypeY = DataTypeBool
	x, ok := ConvertToInt(a.Comparison.X)
	if !ok {
		return errors.New(fmt.Sprintf("could not convert %v to int", a.Comparison.X))
	}
	y, ok := ConvertToInt(yt.Value)
	if !ok {
		return errors.New(fmt.Sprintf("could not convert %v to int", a.Comparison.Y))
	}
	err := a.Comparison.Compute(op, x, y)
	if err != nil {
		return err
	}
	z, ok := a.Comparison.Z.(bool)
	if !ok {
		return errors.New(fmt.Sprintf("could not convert %v to bool", a.Comparison.Z))
	}
	if z {
		if a.DataSlice == nil {
			a.CurrentMinInt = y
			a.CurrentMaxInt = y
			a.DataSlice = []IrisRoutingResponseETL{}
		}
		a.DataSlice = append(a.DataSlice, yt)
	}
	if y > x {
		a.CurrentMaxInt = y
	}
	if y < x {
		a.CurrentMinInt = y
	}
	return nil
}
