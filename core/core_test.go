package core_test

import (
	"regexp"
	"strconv"
	"testing"

	"github.com/OpenBazaar/openbazaar-go/core"
)

func TestEmojiToHTML(t *testing.T) {
	var (
		expected  string
		container = make(map[string]string)
		rx        = regexp.MustCompile(core.EmojiPattern)
		text      = "a #💩 #and #🍦 #😳"
		i         = -1
		replaced  = rx.ReplaceAllStringFunc(text, func(s string) string {
			i++
			key := "_$" + strconv.Itoa(i) + "_"
			container[key] = s
			return key
		})
	)

	expected = "a #_$0_ #and #_$1_ #_$2_"
	if replaced != expected {
		t.Errorf("expected processed string to be %s, but was %s", expected, replaced)
	}

	htmlEnt := core.ToHtmlEntities(text)

	expected = "a #&#x1F4A9; #and #&#x1F366; #&#x1F633;"
	if htmlEnt != expected {
		t.Errorf("expected processed string to be %s, but was %s", expected, replaced)
	}

	recovered := regexp.MustCompile(`\_\$\d+\_`).ReplaceAllStringFunc(replaced, func(s string) string {
		return container[s]
	})
	if recovered != text {
		t.Errorf("expected processed string to be %s, but was %s", text, recovered)
	}
}