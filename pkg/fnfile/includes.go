package fnfile

type Includes []Include

type PullPolicy string

const (
	Always       PullPolicy = "Always"
	IfNotPresent PullPolicy = "IfNotPresent"
	Never        PullPolicy = "Never"
)

type CheckSumAlgorithm string

const (
	SHA1   CheckSumAlgorithm = "SHA1"
	SHA256 CheckSumAlgorithm = "SHA256"
	SHA512 CheckSumAlgorithm = "SHA512"
)

type Include struct {
	Filename  string `json:"filename"`
	Checksum  string `json:"checksum"`
	Algorithm string `json:"algorithm"`
	// Always, IfNotPresent, Never
	PullPolicy  PullPolicy `json:"pullPolicy"`
	PullTimeout Duration   `json:"pullTimeout"`
	Namespace   string     `json:"namespace"`
}
