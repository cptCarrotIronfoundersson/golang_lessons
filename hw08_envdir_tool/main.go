package main

import (
	"flag"
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

	RunCmd(d[1:], env)

}
