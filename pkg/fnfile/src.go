package fnfile

// FileGlobs is a series patterns to match for files
type FileGlobs []string

type Src struct {
	Files FileGlobs `json:"files"`
}
