package mb_json_schemas

import "encoding/json"

type EvalMetric struct {
	EvalMetricStrID            *string                     `json:"evalMetricStrID,omitempty"`
	EvalMetricID               *int                        `json:"evalMetricID"`
	EvalField                  *JsonSchemaField            `json:"evalField,omitempty"`
	EvalName                   *string                     `json:"evalName,omitempty"`
	EvalMetricResult           *EvalMetricResult           `json:"evalMetricResult,omitempty"`
	EvalOperator               string                      `json:"evalOperator"`
	EvalState                  string                      `json:"evalState"`
	EvalExpectedResultState    string                      `json:"evalExpectedResultState"`
	EvalMetricComparisonValues *EvalMetricComparisonValues `json:"evalMetricComparisonValues,omitempty"`
}

type EvalMetricComparisonValues struct {
	EvalComparisonBoolean *bool    `json:"evalComparisonBoolean,omitempty"`
	EvalComparisonNumber  *float64 `json:"evalComparisonNumber,omitempty"`
	EvalComparisonString  *string  `json:"evalComparisonString,omitempty"`
	EvalComparisonInteger *int     `json:"evalComparisonInteger,omitempty"`
}

type EvalMetricResult struct {
	EvalMetricResultStrID     *string `json:"evalMetricsResultStrID"`
	EvalMetricResultID        *int    `json:"evalMetricsResultID"`
	EvalResultOutcomeBool     *bool   `json:"evalResultOutcomeBool,omitempty"`     // true if eval passed, false if eval failed
	EvalResultOutcomeStateStr *string `json:"evalResultOutcomeStateStr,omitempty"` // true if eval passed, false if eval failed
	RunningCycleNumber        *int    `json:"runningCycleNumber,omitempty"`
	EvalIterationCount        *int    `json:"evalIterationCount,omitempty"`

	SearchWindowUnixStart *int            `json:"searchWindowUnixStart,omitempty"`
	SearchWindowUnixEnd   *int            `json:"searchWindowUnixEnd,omitempty"`
	EvalMetadata          json.RawMessage `json:"evalMetadata,omitempty"`
}

type EvalMetaDataResult struct {
	EvalOpCtxStr string `json:"evalOpCtxStr"`
	Operator     string `json:"operator"`
	*EvalMetricComparisonValues
	*FieldValue
}
