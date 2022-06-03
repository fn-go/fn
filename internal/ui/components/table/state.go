package table

type state struct {
	height int
	width  int

	items        Items
	headersOrder []string
	headersMap   map[string]Header
}
