package fnfile

type Inputs map[string]Input

type Input struct {
	name        string
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Default     string `json:"default,omitempty"`
}
