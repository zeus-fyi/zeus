package mb_evals

import (
	mb_json_schemas "github.com/zeus-fyi/zeus/zeus/mockingbird/json_schemas"
)

type EvalFn struct {
	EvalStrID      *string                                          `json:"evalStrID,omitempty"`
	EvalID         *int                                             `json:"evalID,omitempty"`
	OrgID          int                                              `json:"orgID,omitempty"`
	UserID         int                                              `json:"userID,omitempty"`
	EvalName       string                                           `json:"evalName"`
	EvalType       string                                           `json:"evalType"`
	EvalGroupName  string                                           `json:"evalGroupName"`
	EvalModel      *string                                          `json:"evalModel,omitempty"`
	EvalFormat     string                                           `json:"evalFormat"`
	EvalCycleCount int                                              `json:"evalCycleCount,omitempty"`
	Schemas        []*mb_json_schemas.JsonSchemaDefinition          `json:"schemas,omitempty"`
	SchemasMap     map[string]*mb_json_schemas.JsonSchemaDefinition `json:"schemaMap"`
}
