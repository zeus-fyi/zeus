package mb_json_schemas

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
