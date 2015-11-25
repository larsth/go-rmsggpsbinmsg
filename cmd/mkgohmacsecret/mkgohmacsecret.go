package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	var (
		p = make([]byte, sha256.Size224)
	)
	if _, err := rand.Read(p); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	fmt.Fprintf(os.Stdout, "%#v", p)
	fmt.Println("")
	fmt.Println("")
	os.Exit(0)
}
