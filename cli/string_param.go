package cli

import (
	"encoding/json"
	"errors"
	"fmt"
)

type StringParam struct {
	description string
	required    bool

	minLength *int
	maxLength *int
}

func String() StringParam {
	return StringParam{}
}

func (self StringParam) Type() Primitive {
	return STRING
}

func (self StringParam) IsRequired() bool {
	return self.required
}

func (self StringParam) GetDescription() string {
	return self.description
}

func (self StringParam) Description(description string) StringParam {
	self.description = description
	return self
}

func (self StringParam) Required() StringParam {
	self.required = true
	return self
}

func (self StringParam) MinLength(minLength int) StringParam {
	self.minLength = &minLength
	return self
}

func (self StringParam) MaxLength(maxLength int) StringParam {
	self.maxLength = &maxLength
	return self
}

func (self StringParam) Coerce(value string) (any, error) {
	return value, nil
}

func (self StringParam) Validate(value string) error {
	errs := []error{}

	if self.required && value == "" {
		errs = append(errs, errors.New("required"))
	}

	if self.minLength != nil && len(value) < *self.minLength {
		errs = append(errs, errors.New(fmt.Sprintf(
			"length `%d` is less than min_length `%d`",
			len(value),
			*self.minLength,
		)))
	}

	if self.maxLength != nil && len(value) > *self.maxLength {
		errs = append(errs, errors.New(fmt.Sprintf(
			"length `%d` is greater than max_length `%d`",
			len(value),
			*self.maxLength,
		)))
	}

	return errors.Join(errs...)
}

func (self StringParam) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}

func (self StringParam) MarshalJSON() ([]byte, error) {
	data := map[string]any{
		"type":        self.Type(),
		"description": self.description,
		"required":    self.required,
	}

	if self.minLength != nil {
		data["min_length"] = self.minLength
	}

	if self.maxLength != nil {
		data["max_length"] = self.maxLength
	}

	return json.Marshal(data)
}

func (self *StringParam) UnmarshalJSON(b []byte) error {
	data := map[string]any{}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	description, ok := data["description"].(string)

	if !ok {
		return errors.New("`description` is a required string")
	}

	required, ok := data["required"].(bool)

	if !ok {
		return errors.New("`required` is a required boolean")
	}

	var minLength *int

	if v, ok := data["min_length"]; ok {
		value, ok := v.(float64)

		if !ok {
			return errors.New("`min_length` must be ")
		}

		i := int(value)
		minLength = &i
	}

	var maxLength *int

	if v, ok := data["max_length"]; ok {
		value, ok := v.(float64)

		if !ok {
			return errors.New("`max_length` must be ")
		}

		i := int(value)
		maxLength = &i
	}

	if minLength != nil && maxLength != nil {
		if *minLength > *maxLength {
			return errors.New("`min_length` must be less than or equal to `max_length`")
		}
	}

	self.description = description
	self.required = required
	self.minLength = minLength
	self.maxLength = maxLength
	return nil
}
