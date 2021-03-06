package fnfile

type IfSpec Step

type FileCondition struct {
	// Sources matched files have their contents hashed.
	// If the content hashes change between runs, this fn will be marked as "out-of-date".
	Sources FileGlobs `json:"from,omitempty"`

	// Generates matched files are checked to exist.
	// If these files do not exist, this fn will be marked as "out-of-date".
	// When providing a glob, only 1 match is required to keep this fn "up-to-date".
	Generates FileGlobs `json:"makes,omitempty"`
}
