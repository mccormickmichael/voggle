package main

import (
	"fmt"
	"tl/voggle/internal/dict"
)

func main() {
	td := dict.NewTreeDict()

	td.Add("spa")
	td.Add("spam")
	td.Add("spend")
	td.Add("speed")
	td.Add("spoon")
	td.Add("egg")
	td.Add("eat")
	td.Add("bacon")
	td.Add("basic")

	words := make(chan *dict.Node)
	go td.Depth(words)

	for w := range words {
		wp := ""
		if w.IsWord() {
			wp = "*"
		}
		fmt.Printf("%s%s\n", w.Reconstitute(), wp)
	}
}
