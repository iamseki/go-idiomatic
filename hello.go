package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/iamseki/go-idiomatic-001/morestrings"
)

func main() {
	fmt.Println("Hello Idiomatic GO !")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
