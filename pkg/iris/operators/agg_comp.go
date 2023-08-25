package iris_operators

import (
	"errors"
	"fmt"
)

func (a *Aggregation) AggregateLtInt(x IrisRoutingResponseETL) error {
	return a.aggregateCompInt(Lt, x)
}

func (a *Aggregation) AggregateLtEqInt(x IrisRoutingResponseETL) error {
	return a.aggregateCompInt(Lte, x)
}

func (a *Aggregation) AggregateGtInt(x IrisRoutingResponseETL) error {
	return a.aggregateCompInt(Gt, x)
}

func (a *Aggregation) AggregateGtEqInt(x IrisRoutingResponseETL) error {
	return a.aggregateCompInt(Gte, x)
}

func (a *Aggregation) aggregateCompInt(op Op, xt IrisRoutingResponseETL) error {
	a.Comparison.Operator = op
	a.Comparison.DataTypeX = DataTypeInt
	a.Comparison.DataTypeY = DataTypeInt
	a.Comparison.DataTypeZ = DataTypeBool
	y, ok := ConvertToInt(a.Comparison.Y)
	if !ok {
		return errors.New(fmt.Sprintf("could not convert %v to int", a.Comparison.X))
	}
	x, ok := ConvertToInt(xt.Value)
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
		a.DataSlice = append(a.DataSlice, xt)
	}
	if x > a.CurrentMaxInt {
		a.CurrentMaxInt = x
	}
	if x < a.CurrentMinInt {
		a.CurrentMinInt = x
	}
	return nil
}
