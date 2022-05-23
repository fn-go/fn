package fnfile

import (
	"encoding/json"
	"fmt"
	"time"
)

// Duration uses time.ParseDuration (see https://pkg.go.dev/time#ParseDuration) for unmarshalling.
type Duration time.Duration

// MarshalJSON implements the json.Marshaler interface.
func (t Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(t).String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *Duration) UnmarshalJSON(data []byte) error {
	var txt string
	err := json.Unmarshal(data, &txt)
	if err != nil {
		return fmt.Errorf("unmarshalling timeout: %w", err)
	}

	d, err := time.ParseDuration(txt)
	if err != nil {
		return err
	}

	*t = Duration(d)
	return nil
}
