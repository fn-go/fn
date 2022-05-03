package fnfile

type Fn struct {
	// Inputs are parameters for this Fn, used when called by another Fn.
	//
	Inputs Inputs `json:"inputs,omitempty"`

	// Locals are values only available inside this fn via contexts.Locals
	Locals *Vars `json:"locals,omitempty"`

	// Outputs are values that are available via contexts.Fn
	Outputs *Vars `json:"outputs,omitempty"`

	// Short is the short description shown in the 'help' output.
	Short string `json:"short,omitempty"`

	// Long is the long message shown in the 'help <this-fn>' output.
	Long string `json:"long,omitempty"`

	// Example is examples of how to use the fn.
	Example string `json:"example,omitempty"`

	// Steps are the things that will run in serial
	Do Steps `json:"do,omitempty"`

	// Dir is the directory in which steps will be executed from (default: fnfile.yml directory)
	// This dir is available as a context to sub steps, and thus, those steps can use it to make decisions,
	// if specifying their own directory as relative, it will be relative to this.
	Dir string `json:"dir,omitempty"`

	// Env is a map of environment variables that are available to all steps in the fn.
	// You can also set environment variables for the entire fnfile or an individual step.
	//
	// When more than one environment variable is defined with the same name (in different locations),
	// Fn uses the most specific environment variable.
	//
	// For example, an environment variable defined in a step will override fn and global (fnfile) variables
	// with the same name, while the step executes.
	// An environment variable defined for a fn will override a global (fnfile) variable
	// with the same name, while the fn executes.
	Env Vars `json:"env,omitempty"`

	// SerialGroups is an array of arbitrary label-like strings. Executions of this fn
	// and other fns referencing the same tags will be serialized.
	SerialGroups SerialGroups `json:"serialGroups,omitempty"`

	// Matrix is a step/fn hook to define a matrix of step/fn configurations.
	//
	// A matrix allows you to create multiple steps/fns by performing variable substitution
	// in a single step/fn definition.
	//
	// When defined on a fn, it dynamically creates several fns, suffixing the fn name with the matrix values.
	//
	// Each individual fn can be called. The original fn becomes a "virtual fn" that "needs" the matrix of fns.
	// Allowing you to call all matrix fns from the single "parent" virtual fn.
	//
	// For example, you can use a matrix to create fns for more than one supported version of a programming language,
	// variable values, or tool, etc. A matrix reuses the step/fn's configuration
	// and creates a step/fn for each matrix you configure.
	//
	// Example
	//
	//	fns:
	//	  echo:
	//		matrix:
	//		- env: ["dev", "staging", "prod"]
	//		steps:
	//		- run: "echo {{.Matrix.env}}"
	//
	// fns defined are:
	//  echo
	//  echo:dev
	//  echo:staging
	//  echo:prod
	Matrix Matrix `json:"matrix,omitempty"`

	// Timeout is the bounding time limit (duration) for the fn before signalling for termination via SIGINT.
	Timeout Duration `json:"Timeout,omitempty"`

	// TerminateAfter is the bounding time limit (duration) for this fn before sending subprocesses a SIGKILL.
	// Usually specified as a longer duration that timeout.
	TerminateAfter Duration `json:"terminateAfter,omitempty"`
}
