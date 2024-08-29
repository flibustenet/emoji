package emoji

import "testing"

func TestEmojis(t *testing.T) {
	for _, emj := range []struct {
		alias   string
		unicode string
	}{
		{"thumbsup", "👍"},
		{"alarm_clock", "⏰"},
		{"_no_", ":_no_:"},
	} {
		if eu := Emoji(emj.alias); emj.unicode != eu {
			t.Errorf("alias %s is %s instead of %s", emj.alias, emj.unicode, eu)
		}
	}
}
