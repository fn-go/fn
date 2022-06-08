package fnfile

import (
	"encoding/json"
	"sort"

	"github.com/samber/lo"

	"github.com/go-fn/fn/pkg/set"
)

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

	// Include enables different matrix "shapes" aside from the expected NxM table produced from KVs.
	// Each key:value pair will be added to each of the matrix combinations, but only if the keys are unique (not part of the original matrix).
	// For any unoriginal keys, new matrix combinations will be created.
	Include KeyValues `json:"include,omitempty"`
}

func (m *Matrix) UnmarshalJSON(data []byte) (err error) {
	*m, err = UnmarshalMatrix(data)
	return
}

func (m Matrix) Handle(w ResponseWriter, c *FnContext) {
	panic("not implemented!")
}

func UnmarshalMatrixStep(data []byte) (Step, error) {
	return UnmarshalMatrix(data)
}

func UnmarshalMatrix(data []byte) (Matrix, error) {
	type MatrixAlias Matrix
	var tmpMatrix MatrixAlias
	err := json.Unmarshal(data, &tmpMatrix)
	return Matrix(tmpMatrix), err
}

type KeyValues map[string]set.Set[string]

// fruit: [apple, pear, orange]
// animal: [cat, dog]
//
// should generate
//
// apple, cat
// apple, dog
// pear, cat
// pear, dog
// orange, cat
// orange, dog
func GenerateCombinations(kvs KeyValues) []set.Set[string] {
	keys := lo.Keys(kvs)
	sort.Strings(keys)

	return nil
}
