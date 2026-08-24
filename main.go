
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: Encryption-Toolkit <file> [file...]")
		os.Exit(1)
	}
	for _, p := range os.Args[1:] {
		b, err := ioutil.ReadFile(p)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read %s: %v\n", p, err)
			os.Exit(1)
		}
		s256 := sha256.Sum256(b)
		s512 := sha512.Sum512(b)
		fmt.Printf("%s  sha256=%s  sha512=%s\n", p, hex.EncodeToString(s256[:]), hex.EncodeToString(s512[:]))
	}
}
