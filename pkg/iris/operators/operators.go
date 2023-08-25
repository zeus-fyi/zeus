package iris_operators

import (
	"errors"
	"fmt"
)

const ()

const (
	OperatorGt  = ">"
	OperatorLt  = "<"
	OperatorEq  = "=="
	OperatorNeq = "!="
	OperatorGte = ">="
	OperatorLte = "<="

	Gt  = "gt"
	Gte = "gte"
	Lt  = "lt"
	Lte = "lte"
	Eq  = "eq"
	Neq = "neq"

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

func (o Op) Lte() string {
	return OperatorLte
}

func GetOp(op string) (string, error) {
	var o Op
	switch op {
	case Gt:
		return o.Gt(), nil
	case Gte:
		return o.Gte(), nil
	case Lt:
		return o.Lt(), nil
	case Lte:
		return o.Lte(), nil
	case Eq:
		return o.Eq(), nil
	case Neq:
		return o.Neq(), nil
	}
	return "", errors.New(fmt.Sprintf("operator %s not supported", op))
}

type Operation struct {
	Operator  Op     `json:"operator"`
	DataTypeX string `json:"dataTypeX"`
	DataTypeY string `json:"dataTypeY"`
	DataTypeZ string `json:"dataTypeZ"`
	X         any    `json:"dataInX,"`
	Y         any    `json:"dataInY,omitempty"`
	Z         any    `json:"-,omitempty"`
}

func (o *Operation) Compute(operator Op, x any, y any) error {
	o.X = x
	o.Y = y
	o.Operator = operator
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
		o.Z, err = compareIntInt(string(o.Operator), xt, yt)
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
		o.Z, err = compareFloat64Float64(string(o.Operator), xt, yt)
		return err
	default:
		return errors.New(fmt.Sprintf("unregistered data types: %s, %s", o.DataTypeX, o.DataTypeY))
	}
}
