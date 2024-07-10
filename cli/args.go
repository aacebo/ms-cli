package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type Args map[string]reflect.Value

func (self Args) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}

func (self Args) MarshalJSON() ([]byte, error) {
	data := map[string]any{}

	for key, value := range self {
		if !value.CanInterface() {
			return nil, errors.New(fmt.Sprintf(
				"invalid value `%s`",
				value.String(),
			))
		}

		data[key] = value.Interface()
	}

	return json.Marshal(data)
}
