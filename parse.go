package wifky

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
)

func unpackH(src string, decoder *encoding.Decoder) (string, string, error) {
	b := make([]byte, 0, len(src)/2+1)
	lastValue := 0
	var i int
	for i = 0; i < len(src); i++ {
		index := strings.IndexByte("00112233445566778899aAbBcCdDeEfF", src[i])
		if index < 0 {
			break
		}
		index >>= 1
		if i%2 == 0 {
			lastValue = index
		} else {
			b = append(b, byte(lastValue|(index<<4)))
		}
	}

	if decoder != nil {
		var err error
		b, err = decoder.Bytes(b)
		if err != nil {
			return "", "", err
		}
	}
	result := string(b)
	if i%2 == 0 {
		return result, src[i:], nil
	} else {
		return result, src[i:], errors.New("The hexadecimal number is truncated at 4 bits")
	}
}

type Page struct {
	Source string
	Assets map[string]string
}

type Contents struct {
	Pages   map[string]*Page
	Misc    []string
	decoder *encoding.Decoder
}

func (c *Contents) add(s string) {
	name, rest, err := unpackH(s, c.decoder)
	if err != nil {
		c.Misc = append(c.Misc, s)
		return
	}
	if rest != "" {
		if len(rest) < 2 && rest[:2] != "__" {
			c.Misc = append(c.Misc, s)
			return
		}
		var asset string
		asset, rest, err = unpackH(rest[2:], c.decoder)
		if err != nil {
			c.Misc = append(c.Misc, s)
			return
		}
		if page, ok := c.Pages[name]; ok {
			if page.Assets == nil {
				page.Assets = map[string]string{
					asset: s,
				}
			} else {
				page.Assets[asset] = s
			}
		} else {
			c.Pages[name] = &Page{
				Assets: map[string]string{
					asset: s,
				},
			}
		}
	} else {
		if page, ok := c.Pages[name]; ok {
			page.Source = s
		} else {
			c.Pages[name] = &Page{Source: s}
		}
	}
}

func newContents() *Contents {
	return &Contents{
		Pages: make(map[string]*Page),
	}
}

func (c *Contents) setEucJP() {
	c.decoder = japanese.EUCJP.NewDecoder()
}

func (c *Contents) readDir(dirName string) error {
	dirName = strings.TrimRight(dirName, `/\`)
	if strings.HasSuffix(dirName, ".dat") {
		c.setEucJP()
	}
	dirEntry, err := os.ReadDir(dirName)
	if err != nil {
		return err
	}
	for _, entry := range dirEntry {
		if !entry.IsDir() {
			c.add(entry.Name())
		}
	}
	return nil
}

func ReadDir(dirName string) (*Contents, error) {
	c := newContents()
	return c, c.readDir(dirName)
}

const forbiddenChars = `\/:*?"<>|%`

func SafeName(s string) string {
	var buffer strings.Builder
	for _, c := range s {
		if c > 0xFF {
			buffer.WriteRune(c)
		} else {
			index := strings.IndexByte(forbiddenChars, byte(c))
			if index >= 0 || c < ' ' {
				fmt.Fprintf(&buffer, "%%%02X", c)
			} else {
				buffer.WriteByte(byte(c))
			}
		}
	}
	return buffer.String()
}
