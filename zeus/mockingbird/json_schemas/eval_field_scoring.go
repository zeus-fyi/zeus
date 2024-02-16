package mb_json_schemas

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/rs/zerolog/log"
)

func TransformJSONToEvalScoredMetrics(jsonSchemaDef *JsonSchemaDefinition) error {
	for vi, _ := range jsonSchemaDef.Fields {
		for i, _ := range jsonSchemaDef.Fields[vi].EvalMetrics {
			if jsonSchemaDef.Fields[vi].EvalMetrics[i] == nil {
				jsonSchemaDef.Fields[vi].EvalMetrics[i] = &EvalMetric{}
			}
			jsonSchemaDef.Fields[vi].EvalMetrics[i].EvalMetricResult = &EvalMetricResult{}
			eocr := EvalMetaDataResult{
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
