package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	flag.Parse()
	d := flag.Args()
	path := d[0]
	env, err := ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	code := RunCmd(d[1:], env)
	fmt.Println(code)
}
