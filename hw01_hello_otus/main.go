package main

import "fmt"

import (
	"golang.org/x/example/stringutil"
)

func main() {
	ReversedString := stringutil.Reverse("Hello, OTUS!")
	fmt.Println(ReversedString)
}
