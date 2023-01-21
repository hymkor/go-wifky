package wifky

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var Enc = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",
	"\"", "&quot;",
	"'", "&#39;",
	"\r", "",
	"\a", "",
	"\b", "")

var rxVerbatim = regexp.MustCompile(
	"(?sm)^\\s*\\&lt;pre&gt;.*?\n\\s*\\&lt;/pre&gt;|" +
		"^\\s*8\\&lt;.*?\n\\s*\\&gt;8|" +
		"`.`.*?`.`")

var verbList = map[rune]string{}

func verb(text string) string {
	ch := rune(0xE000 + len(verbList))
	verbList[ch] = text
	return string(ch)
}

func verbatim(text string) string {
	return rxVerbatim.ReplaceAllStringFunc(text, func(s string) string {
		s = strings.TrimSpace(s)
		switch s[0] {
		case '`': // ``` ```
			s = s[3 : len(s)-3]
		case '8': // 8< >8 , 8&lt; &gt;8
			s = s[5 : len(s)-5]
		case '&': // <pre> </pre>  , &lt;pre&gt; &lt;/pre&gt;
			s = s[11 : len(s)-12]
		default:
			panic("verbatim: " + s)
		}
		return verb(s)
	})
}

var rxUnverb = regexp.MustCompile("[\uE000-\uF8FF]")

func unverb(text string) string {
	return rxUnverb.ReplaceAllStringFunc(text, func(s string) string {
		r, _ := utf8.DecodeRuneInString(s)
		result, ok := verbList[r]
		if !ok {
			panic("unknown code:" + s)
		}
		return result
	})
}
