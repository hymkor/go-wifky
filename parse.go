package main

import (
	"errors"
	"strings"
	"unicode/utf8"

	"github.com/nyaosorg/go-windows-mbcs"
)

func unpackH(src string) (string, string, error) {
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

	var result string
	if utf8.Valid(b) {
		result = string(b)
	} else {
		var err error
		result, err = mbcs.AtoU(b, mbcs.ACP)
		if err != nil {
			return "", src[i:], err
		}
	}
	if i%2 == 0 {
		return result, src[i:], nil
	} else {
		return result, src[i:], errors.New("The hexadecimal number is truncated at 4 bits")
	}
}
