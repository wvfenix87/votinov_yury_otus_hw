package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	helloPhrase := "Hello, OTUS!"
	helloPhrase = stringutil.Reverse(helloPhrase)
	fmt.Println(helloPhrase)
}
