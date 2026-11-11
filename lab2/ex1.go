package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/sha3"
)

func hashMD5(s string) string {
	h := md5.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}

func hashSHA1(s string) string {
	h := sha1.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}

func hashSHA256(s string) string {
	h := sha256.Sum256([]byte(s))
	return hex.EncodeToString(h[:])
}

func hashSHA512(s string) string {
	h := sha512.Sum512([]byte(s))
	return hex.EncodeToString(h[:])
}

func hashSHA3(s string) string {
	h := sha3.Sum256([]byte(s))
	return hex.EncodeToString(h[:])
}

func compare(a, b string) string {
	if a == b {
		return "Match!"
	}
	return "No Match!"
}

func proofMe(txt1, txt2 string) {
	fmt.Printf("Hash (MD5)    : Output A / output B => %s\n", compare(hashMD5(txt1), hashMD5(txt2)))
	fmt.Printf("Hash (SHA1)   : Output A / output B => %s\n", compare(hashSHA1(txt1), hashSHA1(txt2)))
	fmt.Printf("Hash (SHA256) : Output A / output B => %s\n", compare(hashSHA256(txt1), hashSHA256(txt2)))
	fmt.Printf("Hash (SHA512) : Output A / output B => %s\n", compare(hashSHA512(txt1), hashSHA512(txt2)))
	fmt.Printf("Hash (SHA3)   : Output A / output B => %s\n", compare(hashSHA3(txt1), hashSHA3(txt2)))
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("======== Name + Hashing Program ========")
	fmt.Print("Please input value 1: ")
	txt1, _ := reader.ReadString('\n')
	txt1 = strings.TrimSpace(txt1)

	fmt.Print("Please input value 2: ")
	txt2, _ := reader.ReadString('\n')
	txt2 = strings.TrimSpace(txt2)

	proofMe(txt1, txt2)
}