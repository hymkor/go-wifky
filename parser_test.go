package wifky

import (
	"testing"
)

func TestVerb(t *testing.T) {
	source := `8<
ahaha
ihihi
>8

<pre>
ahaha
ihihi
</pre>
`

	result := unverb(verbatim(Enc.Replace(source)))
	expect := `
ahaha
ihihi


ahaha
ihihi

`
	if expect != result {
		t.Fatalf("expect '%s' but '%s'", expect, result)
	}
}
