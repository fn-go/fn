module github.com/go-fn/fn

go 1.18

require (
	github.com/charmbracelet/bubbles v0.10.3
	github.com/charmbracelet/bubbletea v0.20.1-0.20220516164627-a5f28a3a04bb
	github.com/charmbracelet/glamour v0.5.0
	github.com/charmbracelet/lipgloss v0.5.0
	github.com/ghostsquad/go-timejumper v0.1.2
	github.com/hack-pad/hackpadfs v0.1.2
	github.com/hashicorp/go-multierror v1.1.1
	github.com/kr/pretty v0.3.0
	github.com/oklog/run v1.1.0
	github.com/samber/lo v1.18.0
	github.com/spf13/cobra v1.4.0
	github.com/stretchr/testify v1.7.1
	golang.org/x/exp v0.0.0-20220303212507-bbda1eaf7a17
	mvdan.cc/sh/v3 v3.5.0
	sigs.k8s.io/yaml v1.3.0
)

require (
	github.com/alecthomas/chroma v0.10.0 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/containerd/console v1.0.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/gorilla/css v1.0.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/microcosm-cc/bluemonday v1.0.17 // indirect
	github.com/muesli/ansi v0.0.0-20211031195517-c9f0611b6c70 // indirect
	github.com/muesli/cancelreader v0.2.0 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.11.1-0.20220212125758-44cd13922739 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/yuin/goldmark v1.4.4 // indirect
	github.com/yuin/goldmark-emoji v1.0.1 // indirect
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/term v0.0.0-20220411215600-e5f449aeb171 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

// https://github.com/cornfeedhobo/pflag
// A fork of pflag that is more maintained
replace github.com/spf13/pflag => github.com/cornfeedhobo/pflag v1.1.0
