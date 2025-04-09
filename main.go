package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

//export sha256
func Sha256() {
	sha256HelloWorld()
}

//export sha1
func Sha1() {
	sha1HelloWorld()
}

func main() {
	sha1HelloWorld()
	sha256HelloWorld()
}

func sha256HelloWorld() {
	h := sha256.New()
	h.Write([]byte("hello world"))
	hash := hex.EncodeToString(h.Sum([]byte{}))
	fmt.Printf("sha256: %s\n", hash)
}

func sha1HelloWorld() {
	h := sha1.New()
	h.Write([]byte("hello world"))
	hash := hex.EncodeToString(h.Sum([]byte{}))
	fmt.Printf("sha1: %s\n", hash)
}
