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

func printDetail(name, outA, outB string) {
	fmt.Printf("Hash (%s):\n", name)
	fmt.Printf("  Output A = %s\n", outA)
	fmt.Printf("  Output B = %s\n", outB)
	if outA == outB {
		fmt.Println("  => Match!\n")
	} else {
		fmt.Println("  => No Match!\n")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please input value 1: ")
	txt1, _ := reader.ReadString('\n')
	txt1 = strings.TrimSpace(txt1)

	fmt.Print("Please input value 2: ")
	txt2, _ := reader.ReadString('\n')
	txt2 = strings.TrimSpace(txt2)

	fmt.Println("\n================ Result ================")
	printDetail("MD5", hashMD5(txt1), hashMD5(txt2))
	printDetail("SHA1", hashSHA1(txt1), hashSHA1(txt2))
	printDetail("SHA256", hashSHA256(txt1), hashSHA256(txt2))
	printDetail("SHA512", hashSHA512(txt1), hashSHA512(txt2))
	printDetail("SHA3-256", hashSHA3(txt1), hashSHA3(txt2))
}