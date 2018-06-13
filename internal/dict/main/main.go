package main

import (
	"fmt"
	"tl/dict"
)

func main() {
	d := dict.NewDict()
	d.Add("a")
	fmt.Println(d)
}
