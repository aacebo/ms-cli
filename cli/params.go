package cli

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Param interface {
	Type() Primitive
	GetDescription() string
	IsRequired() bool
	Coerce(value string) (any, error)
	Validate(value string) error
}

type Params map[string]Param

func (self Params) Validate(value ...string) (Args, error) {
	args := Args{}
	errs := []error{}

	for i := 0; i < len(value); i++ {
		if !strings.HasPrefix(value[i], "-") {
			errs = append(errs, errors.New(fmt.Sprintf(
				"invalid argument name `%s`",
				value[i],
			)))

			continue
		}

		key := ""
		val := ""
		alias := !strings.HasPrefix(value[i], "--")

		if strings.Contains(value[i], "=") {
			parts := strings.Split(value[i], "=")
			key = parts[0]
			val = parts[1]
		} else {
			key = value[i]

			if i < len(value)-1 && !strings.HasPrefix(value[i+1], "-") {
				val = value[i+1]
				i++
			}
		}

		if alias {
			key = key[1:]
		} else {
			key = key[2:]
		}

		param, ok := self[key]

		if !ok {
			errs = append(errs, errors.New(fmt.Sprintf(
				"parameter `%s` not found",
				key,
			)))

			continue
		}

		coerced, err := param.Coerce(val)

		if err != nil {
			errs = append(errs, err)
			continue
		}

		args[key] = reflect.ValueOf(coerced)
		errs = append(errs, param.Validate(val))
	}

	for key, param := range self {
		_, ok := args[key]

		if !ok && param.IsRequired() {
			errs = append(errs, errors.New("required"))
		}
	}

	return args, errors.Join(errs...)
}
