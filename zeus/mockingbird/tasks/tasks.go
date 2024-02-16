package mb_tasks

import (
	mb_evals "github.com/zeus-fyi/zeus/zeus/mockingbird/evals"
	mb_json_schemas "github.com/zeus-fyi/zeus/zeus/mockingbird/json_schemas"
	mb_search "github.com/zeus-fyi/zeus/zeus/mockingbird/search"
)

type AITaskLibrary struct {
	TaskStrID             string                                           `db:"-" json:"taskStrID,omitempty"`
	TaskID                int                                              `db:"task_id" json:"taskID,omitempty"`
	MaxTokensPerTask      int                                              `db:"max_tokens_per_task" json:"maxTokensPerTask"`
	TaskType              string                                           `db:"task_type" json:"taskType"`
	TaskName              string                                           `db:"task_name" json:"taskName"`
	TaskGroup             string                                           `db:"task_group" json:"taskGroup"`
	TokenOverflowStrategy string                                           `db:"token_overflow_strategy" json:"tokenOverflowStrategy"`
	Model                 string                                           `db:"model" json:"model"`
	Temperature           float64                                          `db:"temperature" json:"temperature"`
	MarginBuffer          float64                                          `db:"margin_buffer" json:"marginBuffer"`
	Prompt                string                                           `db:"prompt" json:"prompt"`
	Schemas               []*mb_json_schemas.JsonSchemaDefinition          `json:"schemas,omitempty"`
	SchemasMap            map[string]*mb_json_schemas.JsonSchemaDefinition `json:"schemasMap,omitempty"`
	ResponseFormat        string                                           `db:"response_format" json:"responseFormat"`
	CycleCount            int                                              `db:"cycle_count" json:"cycleCount,omitempty"`
	RetrievalDependencies []mb_search.RetrievalItem                        `json:"retrievalDependencies,omitempty"`
	EvalFns               []mb_evals.EvalFn                                `json:"evalFns,omitempty"`
}
