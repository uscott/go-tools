package main

import (
	"flag"
	"fmt"

	"github.com/uscott/go-tools/tgz"
)

func main() {
	var (
		path string
	)
	flag.StringVar(&path, "p", ".", "Path to use")
	flag.Parse()
	if err := tgz.TargzToGzDir(path); err != nil {
		fmt.Println(err)
	}
}
