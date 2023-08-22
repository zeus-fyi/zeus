package iris_operators

import (
	"errors"
	"fmt"
)

const (
	OperatorGt  = ">"
	OperatorLt  = "<"
	OperatorEq  = "=="
	OperatorNeq = "!="
	OperatorGte = ">="
	OperatorLte = "<="

	DataTypeBool    = "bool"
	DataTypeInt     = "int"
	DataTypeStr     = "string"
	DataTypeFloat64 = "float64"
)

type Op string

func (o Op) Gt() string {
	return OperatorGt
}

func (o Op) Lt() string {
	return OperatorLt
}

func (o Op) Eq() string {
	return OperatorEq
}

func (o Op) Neq() string {
	return OperatorNeq
}

func (o Op) Gte() string {
	return OperatorGte
}

type Operation struct {
	Operator  string `json:"operator"`
	DataTypeX string `json:"dataTypeX"`
	DataTypeY string `json:"dataTypeY"`
	DataTypeZ string `json:"dataTypeZ"`
	X         any    `json:"dataInX"`
	Y         any    `json:"dataInY"`
	Z         any    `json:"dataOutZ"`
}

func (o *Operation) Compute(x any, y any, operator Op) error {
	o.X = x
	o.Y = y
	o.Operator = string(operator)
	switch o.DataTypeX + o.DataTypeY + o.DataTypeZ {
	case DataTypeInt + DataTypeInt + DataTypeBool:
		xt, ok := ConvertToInt(x)
		if !ok {
			return errors.New(fmt.Sprintf("could not convert %s to int", x))
		}
		yt, ok := ConvertToInt(y)
		if !ok {
			return errors.New(fmt.Sprintf("could not convert %s to int", y))
		}
		var err error
		o.Z, err = compareIntInt(o.Operator, xt, yt)
		return err
	case DataTypeFloat64 + DataTypeFloat64 + DataTypeBool:
		xt, ok := ConvertToFloat64(x)
		if !ok {
			return errors.New(fmt.Sprintf("could not convert %s to float64", x))
		}
		yt, ok := ConvertToFloat64(y)
		if !ok {
			return errors.New(fmt.Sprintf("could not convert %s to float64", y))
		}
		var err error
		o.Z, err = compareFloat64Float64(o.Operator, xt, yt)
		return err
	default:
		return errors.New(fmt.Sprintf("unregistered data types: %s, %s", o.DataTypeX, o.DataTypeY))
	}
}
