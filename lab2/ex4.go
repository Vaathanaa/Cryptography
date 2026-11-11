package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func getSHA1(s string) string {
	h := sha1.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}

func main() {
	targetHash := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"
	wordlistPath := "nord_vpn.txt"

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

	writer := io.MultiWriter(os.Stdout, logFile)

	fmt.Fprintf(writer, "=== Lab #2: SHA1 Password Cracking ===\n")
	fmt.Fprintf(writer, "Target Hash: %s\n\n", targetHash)

	scanner := bufio.NewScanner(file)
	attempt := 0
	found := false

	// Step 3: Iterate over wordlist using a for loop
	for scanner.Scan() {
		attempt++
		password := strings.TrimSpace(scanner.Text())
		currentHash := getSHA1(password)

		fmt.Fprintf(writer, "[Attempt #%d] Testing: %-20s => %s\n", attempt, password, currentHash)

		if currentHash == targetHash {
			fmt.Fprintf(writer, "\n[+] SUCCESS! SHA1 Password Cracked: %s (after %d attempts)\n", password, attempt)
			found = true
			break
		}
	}

	if !found {
		fmt.Fprintf(writer, "\n[-] Password not found in wordlist.\n")
	}
}