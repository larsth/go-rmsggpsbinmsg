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
	const step = 8
	var (
		size uint32
		p    = make([]byte, sha256.Size224)
		line []byte
		rest []byte
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

	//Pretty print the output
	fmt.Printf("%s\n", "exampleSaltSecret = []byte{")
	rest = p
	for {
		if len(rest) == 0 {
			break
		}
		if len(rest) > step {
			line = rest[:step]
			rest = rest[step:]
			for _, v := range line {
				fmt.Printf("\t0x%X, ", v)
			}
			fmt.Println("")
		} else {
			line = rest
			rest = nil
		}
	}

	//'len(rest) < step' is true:
	for i, v := range line {
		if (len(line) - 1) != i {
			fmt.Printf("\t0x%X, ", v)
		} else {
			fmt.Printf("\t0x%X,\n}\n", v)
		}

	}
	os.Exit(0)
}
