package iris_operators

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

func (a *Aggregation) AggregateMaxFloat64(x float64) error {
	a.Operator = Max
	if len(a.DataSlice) == 0 || x >= a.CurrentMaxFloat64 {
		if x > a.CurrentMaxFloat64 {
			a.CurrentMaxFloat64 = x
			a.DataSlice = []any{x} // Keep only the new maximum value
		} else {
			a.DataSlice = append(a.DataSlice, x) // Append the value if it's equal to the current maximum
		}
	}
	return nil
}

func (a *Aggregation) AggregateMaxInt(x int) error {
	a.Operator = Max
	if len(a.DataSlice) == 0 || x >= a.CurrentMaxInt {
		if x > a.CurrentMaxInt {
			a.CurrentMaxInt = x
			a.DataSlice = []any{x} // Keep only the new maximum value
		} else {
			a.DataSlice = append(a.DataSlice, x) // Append the value if it's equal to the current maximum
		}
	}
	return nil
}

//func (a *Aggregation) AggregateMaxHexstr(x string) error {
//	return nil
//}
