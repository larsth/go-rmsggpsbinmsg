package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"os"

	"github.com/spf13/pflag"
)

func main() {
	var (
		size uint32
		p    = make([]byte, sha256.Size224)
	)
	pflag.CommandLine.Uint32VarP(&size, "size", "s", 8*1024,
		"Size of the salt.")
	pflag.Parse()
	if size == 0 {
		log.Fatalln("The size of the salt is zero(0) bytes, nothing to do, exiting ...")
	}
	p = make([]byte, size)
	if _, err := rand.Read(p); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	if _, err := os.Stdout.Write(p); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	os.Exit(0)
}
