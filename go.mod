module github.com/go-fn/fn

go 1.18

require (
	github.com/ghostsquad/go-timejumper v0.1.2
	github.com/hack-pad/hackpadfs v0.1.2
	github.com/hashicorp/go-multierror v1.1.1
	github.com/mitchellh/go-wordwrap v1.0.1
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6
	github.com/oklog/run v1.1.0
	github.com/samber/lo v1.18.0
	github.com/spf13/cobra v1.4.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.1
	golang.org/x/exp v0.0.0-20220303212507-bbda1eaf7a17
	sigs.k8s.io/yaml v1.3.0
)

require (
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	mvdan.cc/sh/v3 v3.5.0 // indirect
)

// https://github.com/cornfeedhobo/pflag
// A fork of pflag that is more maintained
replace github.com/spf13/pflag => github.com/cornfeedhobo/pflag v1.1.0
