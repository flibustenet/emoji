package emoji

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log/slog"
)

//go:embed emojis.json
var emojison []byte

type alias []string

type uncd string

func (u *uncd) UnmarshalJSON(b []byte) error {
	if string(b) == "false" {
		return nil
	}
	s := ""
	err := json.Unmarshal(b, &s)
	if err != nil {
		return fmt.Errorf("unmarshal unicode %s to string: %v", b, err)
	}
	*u = uncd(s)
	return nil
}
func (a *alias) UnmarshalJSON(b []byte) error {
	//log.Printf("b0=%c alias='%s'", b[0], b)
	if b[0] == '[' {
		array := []string{}
		err := json.Unmarshal(b, &array)
		if err != nil {
			return fmt.Errorf("unmarshal %s to []string: %v", b, err)
		}
		*a = array
		return nil
	}
	s := ""
	err := json.Unmarshal(b, &s)
	if err != nil {
		return fmt.Errorf("unmarshal %s to string: %v", b, err)
	}
	*a = []string{s}
	return nil
}

type emj struct {
	Alias   alias `json:"alias"`
	Unicode uncd  `json:"unicode"`
}

var Emojis map[string]string

func init() {
	Emojis = map[string]string{}
	emjs := []*emj{}
	err := json.Unmarshal(emojison, &emjs)
	if err != nil {
		slog.Error(fmt.Sprintf("unmarshall emojis: %v", err))
		return
	}
	for _, e := range emjs {
		for _, alias := range e.Alias {
			if e.Unicode == "" {
				continue
			}
			Emojis[alias] = string(e.Unicode)
		}
	}
}

func Emoji(s string) string {
	u, ok := Emojis[s]
	if !ok {
		return ":" + s + ":"
	}
	return u
}
