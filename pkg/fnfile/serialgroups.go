package fnfile

import (
	"encoding/json"
	"fmt"

	"github.com/samber/lo"
)

type SerialGroups []string

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *SerialGroups) UnmarshalJSON(data []byte) error {
	tmp := []string{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return fmt.Errorf("unmarshalling serialgroup: %w", err)
	}

	*t = lo.Uniq(tmp)
	return nil
}
