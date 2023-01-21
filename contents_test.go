package wifky

import (
	"testing"
)

func TestUnpackH(t *testing.T) {
	text, rest, err := unpackH("e69716f637e2d616e6", nil)
	if err != nil {
		t.Fatalf("%s/%s: %s\n", text, rest, err.Error())
		return
	}
	if text != "nyaos.man" || rest != "" {
		t.Fatalf("expect nyaos.man but '%s' and rest is '%s'\n", text, rest)
	}
	// println(text, rest)
}

func TestFilesToPages(t *testing.T) {
	c := newContents()
	c.add("b5f4271636c656d5b534b2b2d502f434349423")
	c.add("b5f4271636c656d5b534b2b2d502f434349423__27566656275627e2478747")

	for key, contents := range c.Pages {
		//println(key)
		//for _, a := range assets {
		//println("  ", a)
		//}
		const expectedKey = "[Oracle][C++] OCCI2"
		if key != expectedKey {
			t.Fatalf("key decrypt: expected:'%s' but '%s'", expectedKey, key)
			return
		}
		const expectedValue = "referer.txt"
		if _, ok := contents.Assets[expectedValue]; !ok {
			t.Fatalf("value decrypt: not found '%s'", expectedValue)
			return
		}
	}
}
