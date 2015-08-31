package dockerclient

import (
	"encoding/json"
	"strings"
)

// Entrypoint encapsulates the container entrypoint.
// It might be represented as a string or an array of strings.
// We need to override the json decoder to accept both options.
// The JSON decoder will fail if the api sends an string and
//  we try to decode it into an array of string.
type Entrypoint struct {
	parts []string
}

func (e *Entrypoint) MarshalJSON() ([]byte, error) {
	if e == nil {
		return []byte{}, nil
	}
	return json.Marshal(e.Slice())
}

// UnmarshalJSON decoded the entrypoint whether it's a string or an array of strings.
func (e *Entrypoint) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	p := make([]string, 0, 1)
	if err := json.Unmarshal(b, &p); err != nil {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		p = append(p, s)
	}
	e.parts = p
	return nil
}

func (e *Entrypoint) Len() int {
	if e == nil {
		return 0
	}
	return len(e.parts)
}

func (e *Entrypoint) Slice() []string {
	if e == nil {
		return nil
	}
	return e.parts
}

func NewEntrypoint(parts ...string) *Entrypoint {
	return &Entrypoint{parts}
}

type Command struct {
	parts []string
}

func (e *Command) ToString() string {
	return strings.Join(e.parts, " ")
}

func (e *Command) MarshalJSON() ([]byte, error) {
	if e == nil {
		return []byte{}, nil
	}
	return json.Marshal(e.Slice())
}

// UnmarshalJSON decoded the entrypoint whether it's a string or an array of strings.
func (e *Command) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	p := make([]string, 0, 1)
	if err := json.Unmarshal(b, &p); err != nil {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		p = append(p, s)
	}
	e.parts = p
	return nil
}

func (e *Command) Len() int {
	if e == nil {
		return 0
	}
	return len(e.parts)
}

func (e *Command) Slice() []string {
	if e == nil {
		return nil
	}
	return e.parts
}

func NewCommand(parts ...string) *Command {
	return &Command{parts}
}
