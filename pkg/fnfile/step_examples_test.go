package fnfile

import (
	"fmt"

	"sigs.k8s.io/yaml"
)

func ExampleStepFromStepJson() {
	example := `
sh:
  run: echo "hello"
`

	stepJson := StepJson{}
	if err := yaml.Unmarshal([]byte(example), &stepJson); err != nil {
		panic(err)
	}

	step, err := StepFromStepJson(stepJson)
	if err != nil {
		panic(err)
	}

	fmt.Println(step)
	// Output:
}
