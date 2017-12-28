package jsontypes

import (
	"bytes"
	"encoding/json"
	"errors"
)

var (
	ErrInvalidValue = errors.New("invalid JSON value")
)

type Strings []string

func (a Strings) Contains(s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

func (a Strings) MarshalJSON() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Strings) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewBuffer(b))
	t, err := d.Token()
	if err != nil {
		return err
	}
	v := []string(*a)
	if t == json.Delim('[') {
		var s []string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		v = append(v, s...)
	} else if s, ok := t.(string); ok {
		v = append(v, s)
	} else if t == nil {
		v = nil
	} else if !ok {
		return ErrInvalidValue
	}
	*a = Strings(v)
	return nil
}
