package sites

import "encoding/json"

type Tags []string

func (tags Tags) MarshalJSON() ([]byte, error) {
	if len(tags) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal([]string(tags))
}
