package main

import (
	"fmt"
	"tl/voggle/internal/dict"
)

func main() {
	d := dict.NewTreeDict()
	d.Add("a")
	fmt.Println(d)
}
