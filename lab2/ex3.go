package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func getMD5(s string) string {
	h := md5.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}

func main() {
	targetHash := "6a85dfd77d9cb35770c9dc6728d73d3f"
	wordlistPath := "nord_vpn.txt"

	// Open wordlist
	file, err := os.Open(wordlistPath)
	if err != nil {
		fmt.Println("Error opening wordlist:", err)
		return
	}
	defer file.Close()

	// Step 5: Save output to verbose.txt
	logFile, err := os.Create("verbose.txt")
	if err != nil {
		fmt.Println("Error creating verbose.txt:", err)
		return
	}
	defer logFile.Close()

	// MultiWriter prints to terminal and writes to verbose.txt
	writer := io.MultiWriter(os.Stdout, logFile)

	fmt.Fprintf(writer, "=== Lab #1: MD5 Password Cracking ===\n")
	fmt.Fprintf(writer, "Target Hash: %s\n\n", targetHash)

	scanner := bufio.NewScanner(file)
	attempt := 0
	found := false

	// Step 3: Iterate over wordlist using a for loop
	for scanner.Scan() {
		attempt++
		password := strings.TrimSpace(scanner.Text())
		currentHash := getMD5(password)

		fmt.Fprintf(writer, "[Attempt #%d] Testing: %-20s => %s\n", attempt, password, currentHash)

		if currentHash == targetHash {
			fmt.Fprintf(writer, "\n[+] SUCCESS! MD5 Password Cracked: %s (after %d attempts)\n", password, attempt)
			found = true
			break
		}
	}

	if !found {
		fmt.Fprintf(writer, "\n[-] Password not found in wordlist.\n")
	}
}