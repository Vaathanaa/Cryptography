package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"regexp"
)

func getMD5(s string) string {
	h := md5.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}

func main() {
	fmt.Println("=== Lab #4: Mini-CTF (Cat Leak) ===")

	// 1. URL Query Unescape
	encodedQuery := "%62%36%63%63%62%34%65%63%65%35%34%35%34%64%63%61%65%35%31%37%37%38%62%33%65%32%33%39%65%62%63%32"
	leakedHash, _ := url.QueryUnescape(encodedQuery)
	fmt.Printf("[1] Leaked Hash from URL: %s\n", leakedHash)

	// 2. Decode the cat password:
	// ASCII hex \x6d\x65\x6f\x77 = "meow", repeated 2 times = "meowmeow"
	catPassword := "meowmeow"
	fmt.Printf("[2] Cat Password        : %s\n", catPassword)

	// 3. Verify MD5 Hash
	computedHash := getMD5(catPassword)
	fmt.Printf("[3] MD5(\"%s\")    : %s\n", catPassword, computedHash)

	if computedHash == leakedHash {
		fmt.Println("[+] Match confirmed! The developer used MD5 instead of bcrypt.")
	}

	// 4. Verify Flag using the Regex from the slide
	flag := fmt.Sprintf("cryptoCTF{%s}", catPassword)
	regexPattern := `^cryptoCTF\{(?:meow){2}\}$`
	isMatch, _ := regexp.MatchString(regexPattern, flag)

	fmt.Println("\n================ FLAG ================")
	fmt.Printf("Flag       : %s\n", flag)
	fmt.Printf("Regex Valid: %t\n", isMatch)
	fmt.Println("======================================")
}