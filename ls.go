//go:build ignore

package main

import (
	"fmt"
	"os"

	"github.com/hymkor/go-wifky"
)

func main() {
	for _, s := range os.Args[1:] {
		contents, err := wifky.ReadDir(s)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for pageName, contents := range contents.Pages {
			fmt.Printf("\"%s\" -> %s\n", wifky.SafeName(pageName), contents.Source)
			if contents.Assets != nil {
				for assetName, src := range contents.Assets {
					fmt.Printf("  \"%s\" -> %s\n", wifky.SafeName(assetName), src)
				}
			}
		}
		for _, fn := range contents.Misc {
			fmt.Println("passed:", fn)
		}
	}
}
