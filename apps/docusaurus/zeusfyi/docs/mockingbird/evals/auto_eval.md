---
sidebar_position: 1
displayed_sidebar: mockingbird
---

# Overview

## Eval States

### Info

This is the non-case specific state, i.e. not using for filtering or error handling. It's used for informational purposes

### Filter

![ScreenshoM](https://github.com/zeus-fyi/zeus/assets/17446735/ad0286bc-3f0f-4d24-bef0-73ed4e788301)

The default behavior is after an eval has completed, it will filter out the object(s) from the final result if the eval
state "result" does not match the expected "pass" or "fail" condition. 
```json
    [
        {
            "msg_id": 0,
            "spam_score": 3
        },
        {
            "msg_id": 1,
            "spam_score": 6
        }
    ]
```
I.e. say you have an array of objects of the object type, and your eval is

    expecting msg_id > 0 to pass
        desired outcome: msg_id > 0

    expecting spam_score > 4 to fail
        desired outcome: spam_score > 4 to fail
        actual outcome: spam_score = 3, so since it did not fail like we expected,
                        this object will be filtered out of the final result.

With those two filters applied, the final result for task result, and for sending to next aggregation, trigger, or other 
modular stages would be:

```json
    [
        {
            "msg_id": 1,
            "spam_score": 6
        }
    ]
```

### This Go function simulates the filtering process:
```go

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

type JsonSchemaDefinition struct {
    SchemaID          int               `db:"schema_id" json:"schemaID"`
    SchemaStrID       string            `db:"-" json:"schemaStrID,omitempty"`
    SchemaName        string            `db:"schema_name" json:"schemaName"`
    SchemaGroup       string            `db:"schema_group" json:"schemaGroup"`
    SchemaDescription string            `db:"schema_description" json:"schemaDescription"`
    IsObjArray        bool              `db:"is_obj_array" json:"isObjArray"`
    Fields            []JsonSchemaField `db:"-" json:"fields"`
    FieldsMap         map[string]*JsonSchemaField
    ScoredEvalMetrics []*EvalMetric `db:"-" json:"totalEvalMetrics,omitempty"`
}

type JsonSchemaField struct {
    FieldID          int    `db:"field_id" json:"fieldID"`
    FieldStrID       string `db:"-" json:"fieldStrID,omitempty"`
    FieldName        string `db:"field_name" json:"fieldName"`
    FieldDescription string `db:"field_description" json:"fieldDescription"`
    DataType         string `db:"data_type" json:"dataType"`
    FieldValue
    EvalMetrics []*EvalMetric `db:"-" json:"evalMetrics,omitempty"`
}

type FieldValue struct {
    IntegerValue      *int      `db:"-" json:"intValue,omitempty"`
    StringValue       *string   `db:"-" json:"stringValue,omitempty"`
    NumberValue       *float64  `db:"-" json:"numberValue,omitempty"`
    BooleanValue      *bool     `db:"-" json:"booleanValue,omitempty"`
    IntegerValueSlice []int     `db:"-" json:"intValueSlice,omitempty"`
    StringValueSlice  []string  `db:"-" json:"stringValueSlice,omitempty"`
    NumberValueSlice  []float64 `db:"-" json:"numberValueSlice,omitempty"`
    BooleanValueSlice []bool    `db:"-" json:"booleanValueSlice,omitempty"`
    IsValidated       bool      `db:"-" json:"isValidated,omitempty"`
}

type JsonResponseGroupsByOutcome struct {
    EvalResultsTriggersOn string                                        `json:"evalResultsTriggerOn"`
    Passed                []artemis_orchestrations.JsonSchemaDefinition `json:"passed"`
    Failed                []artemis_orchestrations.JsonSchemaDefinition `json:"failed"`
}

type JsonResponseGroupsByOutcomeMap map[string]JsonResponseGroupsByOutcome


func FilterPassingEvalPassingResponses(jres []artemis_orchestrations.JsonSchemaDefinition) JsonResponseGroupsByOutcomeMap {
	jro := make(map[string]JsonResponseGroupsByOutcome)
	jro["filter"] = JsonResponseGroupsByOutcome{
		Passed: []artemis_orchestrations.JsonSchemaDefinition{},
		Failed: []artemis_orchestrations.JsonSchemaDefinition{},
	}
	for _, jr := range jres {
		if jr.ScoredEvalMetrics == nil {
			continue
		}
		count := 0
		for _, er := range jr.ScoredEvalMetrics {
			if er.EvalExpectedResultState == "ignore" {
				count += 1
				continue
			}
			if er.EvalState != "filter" {
				continue
			}
			if er.EvalExpectedResultState == "pass" && er.EvalMetricResult != nil && er.EvalMetricResult.EvalResultOutcomeBool != nil && *er.EvalMetricResult.EvalResultOutcomeBool {
				count += 1
			} else if er.EvalExpectedResultState == "fail" && er.EvalMetricResult != nil && er.EvalMetricResult.EvalResultOutcomeBool != nil && !*er.EvalMetricResult.EvalResultOutcomeBool {
				count += 1
			}
		}
		if count == len(jr.ScoredEvalMetrics) && len(jr.ScoredEvalMetrics) > 0 {
			tmp := jro["filter"]
			tmp.Passed = append(tmp.Passed, jr)
			jro["filter"] = tmp
		} else {
			tmp := jro["filter"]
			tmp.Failed = append(tmp.Failed, jr)
			jro["filter"] = tmp
		}
	}
	return jro
}

```


When you set the filter state, you should set all the fields in that object to the filter state as well, otherwise
you'll need to set the other fields to "ignore" the result.


### Error

This is reserved for now, but acts like an info state for now. We're planning to use it for things like terminating the workflow 
if > 50% of the objects in the result are in an error state, or if the error state is set to terminate the workflow. 

If you have any specific use cases for this you like us to include in the feature set please let us know! support@zeus.fyi


## Result States

### Pass

Passing means the result of your operators did match the expected result.

```text
Operator:
    if len(str) > 5
        
        Result:
            len("123456") > 5 == true | Expected: PASS, Actual: PASS
            len("123") > 5 == false   | Expected: PASS, Actual: FAIL
```

### Fail

Failing means the result of your operators did not match the expected result.

```text
Operator:
    if len(str) > 5
        
        Result:
            len("123456") > 5 == true | Expected: FAIL, Actual: PASS
            len("123") > 5 == false   | Expected: FAIL, Actual: FAIL
```

### Ignore

When using filtering via EvalState=='filter' the result will be ignored and not used in the final result of group metric scoring.
I.e. if you have 3 fields in an object, and two of them are set to filtered, only those two fields are used for 
determining the final result of the group metric scoring.

## JSON Data Types

### Output Formats

- JSON Object Array
- JSON Object

### Object Key Data Types

#### Field Values
- Boolean
- Integer
- Number
- String

#### Field Groups
- Array[Integer]
- Array[Number]
- Array[String]
- Array[Boolean]

## Reserved Keywords

### Indexer Usage

- msg_id
- msg_body

When you retrieve data from an indexer source, you can use these reserved keywords to add to an eval request to filter
msg_id > 0, or msg_body contains "spam" for example. 

The msg_body keyword returns the retrieval platform specific content. For example, if you're using the Twitter indexer,
msg_body would return the tweet timestamp and content as msg_id and msg_body respectively.

## Eval Comparison Code

To Simulate Eval Scoring, we use the following code to transform the JSON Schema Definition to Eval Scored Metrics.

```go
func TransformJSONToEvalScoredMetrics(jsonSchemaDef *artemis_orchestrations.JsonSchemaDefinition) error {
	for vi, _ := range jsonSchemaDef.Fields {
		for i, _ := range jsonSchemaDef.Fields[vi].EvalMetrics {
			if jsonSchemaDef.Fields[vi].EvalMetrics[i] == nil {
				jsonSchemaDef.Fields[vi].EvalMetrics[i] = &artemis_orchestrations.EvalMetric{}
			}
			jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricResult = &artemis_orchestrations.EvalMetricResult{}
			eocr := artemis_orchestrations.EvalMetaDataResult{
				EvalOpCtxStr:               "",
				Operator:                   jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalOperator,
				EvalMetricComparisonValues: jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues,
				FieldValue:                 &jsonSchemaDef.Fields[vi].FieldValue,
			}
			switch jsonSchemaDef.Fields[vi].DataType {
			case "integer":
				if jsonSchemaDef.Fields[vi].IntegerValue == nil {
					return fmt.Errorf("no int value for key '%s'", jsonSchemaDef.Fields[vi].FieldName)
				}
				if jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonNumber == nil && jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonInteger == nil {
					return fmt.Errorf("no comparison number for key '%s'", jsonSchemaDef.Fields[vi].FieldName)
				}
				if jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonNumber != nil {
					fv := aws.ToFloat64(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonNumber)
					jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonInteger = aws.Int(int(fv))
					jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonNumber = nil
				}

				if jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonInteger != nil {
					jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricResult.EvalResultOutcomeBool = aws.Bool(GetIntEvalComparisonResult(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalOperator, *jsonSchemaDef.Fields[vi].IntegerValue, aws.ToInt(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonInteger)))
					eocr.EvalOpCtxStr = fmt.Sprintf("%s %d %s %d", jsonSchemaDef.Fields[vi].FieldName, aws.ToInt(jsonSchemaDef.Fields[vi].IntegerValue), jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalOperator, aws.ToInt(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonInteger))
				} else if jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonNumber != nil {
					eocr.EvalOpCtxStr = fmt.Sprintf("%s %d %s %d", jsonSchemaDef.Fields[vi].FieldName, aws.ToInt(jsonSchemaDef.Fields[vi].IntegerValue), jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalOperator, int(aws.ToFloat64(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonNumber)))
					jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricResult.EvalResultOutcomeBool = aws.Bool(GetIntEvalComparisonResult(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalOperator, aws.ToInt(jsonSchemaDef.Fields[vi].IntegerValue), int(aws.ToFloat64(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonNumber))))
				} else {
					return fmt.Errorf("no comparison number for key '%s'", jsonSchemaDef.Fields[vi].FieldName)
				}
			case "number":
				if jsonSchemaDef.Fields[vi].NumberValue == nil {
					return fmt.Errorf("no number value for key '%s'", jsonSchemaDef.Fields[vi].FieldName)
				}
				if jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonNumber == nil {
					return fmt.Errorf("no comparison number for key '%s'", jsonSchemaDef.Fields[vi].FieldName)
				}
				eocr.EvalOpCtxStr = fmt.Sprintf("%s %f %s %f", jsonSchemaDef.Fields[vi].FieldName, aws.ToFloat64(jsonSchemaDef.Fields[vi].NumberValue), jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalOperator, aws.ToFloat64(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonNumber))
				jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricResult.EvalResultOutcomeBool = aws.Bool(GetNumericEvalComparisonResult(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalOperator, aws.ToFloat64(jsonSchemaDef.Fields[vi].NumberValue), aws.ToFloat64(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonNumber)))
			case "string":
				if jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonString == nil {
					return fmt.Errorf("no comparison string for key '%s'", jsonSchemaDef.Fields[vi].FieldName)
				}
				if jsonSchemaDef.Fields[vi].StringValue == nil {
					return fmt.Errorf("no string value for key '%s'", jsonSchemaDef.Fields[vi].FieldName)
				}
				eocr.EvalOpCtxStr = fmt.Sprintf("%s %s %s %s", jsonSchemaDef.Fields[vi].FieldName, aws.ToString(jsonSchemaDef.Fields[vi].StringValue), jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalOperator, aws.ToString(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonString))
				jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricResult.EvalResultOutcomeBool = aws.Bool(GetStringEvalComparisonResult(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalOperator, aws.ToString(jsonSchemaDef.Fields[vi].StringValue), aws.ToString(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonString)))
			case "boolean":
				if jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonBoolean == nil {
					return fmt.Errorf("no comparison boolean for key '%s'", jsonSchemaDef.Fields[vi].FieldName)
				}
				if jsonSchemaDef.Fields[vi].BooleanValue == nil {
					return fmt.Errorf("no boolean value for key '%s'", jsonSchemaDef.Fields[vi].FieldName)
				}

				eocr.EvalOpCtxStr = fmt.Sprintf("%s %t %s %t", jsonSchemaDef.Fields[vi].FieldName, *jsonSchemaDef.Fields[vi].BooleanValue, jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalOperator, *jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonBoolean)
				jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricResult.EvalResultOutcomeBool = aws.Bool(GetBooleanEvalComparisonResult(*jsonSchemaDef.Fields[vi].BooleanValue, *jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonBoolean))
			case "array[integer]":
				if jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonNumber == nil && jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonInteger == nil {
					return fmt.Errorf("no comparison number for key '%s'", jsonSchemaDef.Fields[vi].FieldName)
				}
				if jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonInteger != nil {
					results, rerr := EvaluateIntArray(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalOperator, jsonSchemaDef.Fields[vi].IntegerValueSlice, *jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonInteger)
					if rerr != nil {
						return rerr
					}
					eocr.EvalOpCtxStr = fmt.Sprintf("%s %d %s %d", jsonSchemaDef.Fields[vi].FieldName, jsonSchemaDef.Fields[vi].IntegerValueSlice, jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalOperator, aws.ToInt(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonInteger))
					jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricResult.EvalResultOutcomeBool = aws.Bool(Pass(results))
				} else {
					results, rerr := EvaluateIntArray(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalOperator, jsonSchemaDef.Fields[vi].IntegerValueSlice, int(*jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonNumber))
					if rerr != nil {
						return rerr
					}
					eocr.EvalOpCtxStr = fmt.Sprintf("%s %v %s %f", jsonSchemaDef.Fields[vi].FieldName, jsonSchemaDef.Fields[vi].IntegerValueSlice, jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalOperator, aws.ToFloat64(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonNumber))
					jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricResult.EvalResultOutcomeBool = aws.Bool(Pass(results))
				}
			case "array[number]":
				if jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonNumber == nil {
					return fmt.Errorf("no comparison number for key '%s'", jsonSchemaDef.Fields[vi].FieldName)
				}
				results, rerr := EvaluateNumericArray(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalOperator, jsonSchemaDef.Fields[vi].NumberValueSlice, aws.ToFloat64(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonNumber))
				if rerr != nil {
					return rerr
				}
				eocr.EvalOpCtxStr = fmt.Sprintf("%s %f %s %f", jsonSchemaDef.Fields[vi].FieldName, jsonSchemaDef.Fields[vi].NumberValueSlice, jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalOperator, aws.ToFloat64(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonNumber))
				jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricResult.EvalResultOutcomeBool = aws.Bool(Pass(results))
			case "array[string]":
				results, rerr := EvaluateStringArray(jsonSchemaDef.Fields[vi].StringValueSlice, jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalOperator, aws.ToString(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonString))
				if rerr != nil {
					return rerr
				}
				eocr.EvalOpCtxStr = fmt.Sprintf("%s %s %s %s", jsonSchemaDef.Fields[vi].FieldName, jsonSchemaDef.Fields[vi].StringValueSlice, jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalOperator, aws.ToString(jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonString))
				jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricResult.EvalResultOutcomeBool = aws.Bool(Pass(results))
			case "array[boolean]":
				results, rerr := EvaluateBooleanArray(jsonSchemaDef.Fields[vi].BooleanValueSlice, *jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonBoolean)
				if rerr != nil {
					return rerr
				}
				eocr.EvalOpCtxStr = fmt.Sprintf("%s %t %s %t", jsonSchemaDef.Fields[vi].FieldName, jsonSchemaDef.Fields[vi].BooleanValueSlice, jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalOperator, *jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricComparisonValues.EvalComparisonBoolean)
				jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricResult.EvalResultOutcomeBool = aws.Bool(Pass(results))
			default:
				return fmt.Errorf("unknown data type '%s'", jsonSchemaDef.Fields[vi].DataType)
			}
			b, err := json.Marshal(eocr)
			if err != nil {
				log.Err(err).Msg("failed to marshal eval op ctx")
				return err
			}
			jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricResult.EvalMetadata = b
		}
	}
	return nil
}

func GetBooleanEvalComparisonResult(actual, expected bool) bool {
	return actual == expected
}

func GetIntEvalComparisonResult(operator string, actual, expected int) bool {
	switch operator {
	case "==", "eq":
		return actual == expected
	case "!=", "neq":
		return actual != expected
	case ">", "gt":
		return actual > expected
	case "<", "lt":
		return actual < expected
	case ">=", "gte":
		return actual >= expected
	case "<=", "lte":
		return actual <= expected
	}
	return false
}

func GetNumericEvalComparisonResult(operator string, actual, expected float64) bool {
	switch operator {
	case "==", "eq":
		return actual == expected
	case "!=", "neq":
		return actual != expected
	case ">", "gt":
		return actual > expected
	case "<", "lt":
		return actual < expected
	case ">=", "gte":
		return actual >= expected
	case "<=", "lte":
		return actual <= expected
	}
	return false
}

func EvaluateIntArray(operator string, array []int, expected int) ([]bool, error) {
	var results []bool
	for _, value := range array {
		result := GetIntEvalComparisonResult(operator, value, expected)
		results = append(results, result)
	}
	return results, nil
}

func EvaluateNumericArray(operator string, array []float64, expected float64) ([]bool, error) {
	var results []bool
	for _, value := range array {
		result := GetNumericEvalComparisonResult(operator, value, expected)
		results = append(results, result)
	}
	return results, nil
}

func Pass(results []bool) bool {
	for _, result := range results {
		if !result {
			return false
		}
	}
	return true
}

func EvaluateBooleanArray(array []bool, expected bool) ([]bool, error) {
	var results []bool
	for _, value := range array {
		result := GetBooleanEvalComparisonResult(value, expected)
		results = append(results, result)
	}
	return results, nil
}

func GetStringEvalComparisonResult(operator string, actual, expected string) bool {
switch operator {
case "equals-one-from-list":
acceptable := strings.Split(expected, ",")
for _, a := range acceptable {
if actual == a {
return true
}
}
	case "contains":
		if strings.Contains(actual, expected) {
			return true
		}
	case "has-prefix":
		if strings.HasPrefix(actual, expected) {
			return true
		}
	case "has-suffix":
		if strings.HasSuffix(actual, expected) {
			return true
		}
	case "does-not-start-with-any":
		fs := &strings_filter.FilterOpts{
			DoesNotStartWithThese: strings.Split(expected, ","),
		}
		return strings_filter.FilterStringWithOpts(actual, fs)
	case "does-not-include":
		fs := &strings_filter.FilterOpts{
			DoesNotInclude: strings.Split(expected, ","),
		}
		return strings_filter.FilterStringWithOpts(actual, fs)
	case "equals":
		return actual == expected
	case "length-eq":
		expectedLen := len(expected)
		comparedLengthLimit, err := strconv.Atoi(expected)
		if err == nil {
			expectedLen = comparedLengthLimit
		}
		actualLen := len(actual)
		return actualLen == expectedLen
	case "length-less-than":
		expectedLen := len(expected)
		comparedLengthLimit, err := strconv.Atoi(expected)
		if err == nil {
			expectedLen = comparedLengthLimit
		}
		actualLen := len(actual)
		if actualLen < expectedLen {
			return true
		}
	case "length-less-than-eq":
		expectedLen := len(expected)
		comparedLengthLimit, err := strconv.Atoi(expected)
		if err == nil {
			expectedLen = comparedLengthLimit
		}
		if len(actual) <= expectedLen {
			return true
		}
	case "length-greater-than":
		expectedLen := len(expected)
		comparedLengthLimit, err := strconv.Atoi(expected)
		if err == nil {
			expectedLen = comparedLengthLimit
		}
		if len(actual) > expectedLen {
			return true
		}
	case "length-greater-than-eq":
		expectedLen := len(expected)
		comparedLengthLimit, err := strconv.Atoi(expected)
		if err == nil {
			expectedLen = comparedLengthLimit
		}
		if len(actual) >= expectedLen {
			return true
		}
	}

	return false
}

func EvaluateStringArray(array []string, operator, expected string) ([]bool, error) {
	var results []bool
	seen := make(map[string]bool)
	for _, value := range array {
		if operator == "all-unique-words" {
			_, ok := seen[value]
			if ok {
				results = append(results, false)
			} else {
				results = append(results, true)
			}
		} else {
			result := GetStringEvalComparisonResult(operator, value, expected)
			results = append(results, result)
		}
		seen[value] = true
	}
	return results, nil
}
```
