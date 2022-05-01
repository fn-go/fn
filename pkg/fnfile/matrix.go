package fnfile

import (
	"context"
)

// Matrix is a step/fn hook to define a matrix of step/fn configurations.
//
// A matrix allows you to create multiple steps/tasks by performing variable substitution
// in a single step/fn definition.
//
// When defined on a StepCommon, it dynamically creates several steps. Each step is inlined, and are run serially.
// StepCommon ordering is based on the order of each provided matrix.
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
	KVs KeyValues `json:"kvs,omitempty"`

	// Includes lets you add additional configuration options to a build matrix step/task that already exists.
	Includes KeyValues `json:"includes,omitempty"`
}

func (m *Matrix) Visit(ctx context.Context, v StepVisitor) error {
	return v.VisitMatrix(ctx, m)
}

type KeyValues struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}
