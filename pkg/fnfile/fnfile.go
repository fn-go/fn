package fnfile

type Fnfile struct {
	// Version is the version of the Fn schema
	// It is the first thing that is parsed by Task in order to inform how the rest of the file should be parsed.
	Version string `json:"version,omitempty"`

	// Features are flags for fn
	// It enables users to change the behavior of fn by enabling/disabling features of the engine
	Features Features `json:"features,omitempty"`

	// Globals are variables that are available from any fn
	Globals Vars `json:"globals,omitempty"`

	// Includes allows other Fnfiles to be included in this one. Each other include is becomes "namespaced".
	// More on namespacing later...
	Includes Includes `json:"includes,omitempty"`

	// Fns is a set of function definitions
	Fns map[string]Fn `json:"tasks,omitempty"`
}
