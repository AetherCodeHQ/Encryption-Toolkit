package main

// Encryption-Toolkit: AES-256-GCM ve XOR encryption
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: encryption-toolkit <encrypt|decrypt> <file> <key-hex>")
		os.Exit(1)
	}
	mode := os.Args[1]
	path := os.Args[2]
	keyHex := os.Args[3]

	key, err := hex.DecodeString(keyHex)
	if err != nil || len(key) != 32 {
		fmt.Fprintln(os.Stderr, "key must be 32 bytes hex (64 chars)")
		os.Exit(1)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch mode {
	case "encrypt":
		nonce := make([]byte, 12)
		io.ReadFull(rand.Reader, nonce)
		gcm, _ := cipher.NewGCM(block)
		ciphertext := gcm.Seal(nil, nonce, data, nil)
		out := append(nonce, ciphertext...)
		outPath := path + ".enc"
		os.WriteFile(outPath, out, 0644)
		fmt.Printf("encrypted %d -> %d bytes -> %s\n", len(data), len(out), outPath)
	case "decrypt":
		if len(data) < 12 {
			fmt.Fprintln(os.Stderr, "ciphertext too short")
			os.Exit(1)
		}
		nonce := data[:12]
		ciphertext := data[12:]
		gcm, _ := cipher.NewGCM(block)
		plain, err := gcm.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			fmt.Fprintln(os.Stderr, "decrypt:", err)
			os.Exit(1)
		}
		outPath := path + ".dec"
		os.WriteFile(outPath, plain, 0644)
		fmt.Printf("decrypted %d -> %d bytes -> %s\n", len(data), len(plain), outPath)
	default:
		fmt.Println("unknown mode:", mode)
		os.Exit(1)
	}
}
