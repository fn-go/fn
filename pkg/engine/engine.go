package engine

type engine struct{}

type Options struct{}

// New returns a new engine
func New(options ...func(engineOptions *Options)) (*engine, error) {
	opts := &Options{}

	for _, o := range options {
		o(opts)
	}

	var eng engine

	return &eng, nil
}
