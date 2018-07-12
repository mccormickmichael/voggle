package main

import (
	"fmt"
	"tl/voggle/internal/dict"
)

func main() {
	d := dict.EmptyTreeDict()
	d.Add("a")
	fmt.Println(d)
}
