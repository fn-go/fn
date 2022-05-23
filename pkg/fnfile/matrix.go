package fnfile

// Matrix is a step/fn hook to define a matrix of step/fn configurations.
//
// A matrix allows you to create multiple steps/tasks by performing variable substitution
// in a single step/fn definition.
//
// When defined on a StepMeta, it dynamically creates several steps. Each step is inlined, and are run serially.
// StepMeta ordering is based on the order of each provided matrix.
// Example
//
//	task:
//	  echo:
//		matrix:
//		- env: ["dev", "staging", "prod"]
//      - ham: ["bacon", "eggs"]
//		steps:
//		- run: "echo {{.Matrix.env}} {{.Matrix.ham}}"
//
// Steps created defined are:
//  - echo "dev bacon"
//  - echo "dev eggs"
//  - echo "staging bacon"
//  - echo "staging eggs"
//  - echo "prod bacon"
//  - echo "prod eggs"
type Matrix struct {
	StepMeta

	KVs KeyValues `json:"kvs,omitempty"`

	// Includes lets you add additional configuration options to a build matrix step/task that already exists.
	Includes KeyValues `json:"includes,omitempty"`
}

func (m *Matrix) Exec(w ResponseWriter, c *CallInfo) {
	validateHandlerParams(w, c)
}

type KeyValues struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}
