package main

import (
	"bufio"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func getSHA512(s string) string {
	h := sha512.Sum512([]byte(s))
	return hex.EncodeToString(h[:])
}

func main() {
	targetHash := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"
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

	fmt.Fprintf(writer, "=== Lab #3: SHA512 Password Cracking ===\n")
	fmt.Fprintf(writer, "Target Hash: %s\n\n", targetHash)

	scanner := bufio.NewScanner(file)
	attempt := 0
	found := false

	// Step 3: Iterate over wordlist using a for loop
	for scanner.Scan() {
		attempt++
		password := strings.TrimSpace(scanner.Text())
		currentHash := getSHA512(password)

		fmt.Fprintf(writer, "[Attempt #%d] Testing: %-20s => %s\n", attempt, password, currentHash)

		if currentHash == targetHash {
			fmt.Fprintf(writer, "\n[+] SUCCESS! SHA512 Password Cracked: %s (after %d attempts)\n", password, attempt)
			found = true
			break
		}
	}

	if !found {
		fmt.Fprintf(writer, "\n[-] Password not found in wordlist.\n")
	}
}