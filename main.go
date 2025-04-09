package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

//export hash
func Hash() {
	hashHelloWorld()
}

func main() {
	hashHelloWorld()
}

func hashHelloWorld() {
	h := sha256.New()
	h.Write([]byte("hello world"))
	hash := hex.EncodeToString(h.Sum([]byte{}))
	fmt.Printf("hash: %s", hash)
}
