package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

var rxEmail = regexp.MustCompile(`[\+\-\.0-9A-Z_a-z]+@[\-\.0-9A-Z_a-z]+[\.][\-\.0-9A-Z_a-z]+`)

func main() {
	fileNames := []string{
		"sample/01.html",
		"sample/02.html",
		"sample/03.html",
		"sample/04.html",
	}

	for _, fileName := range fileNames {
		data := openFile(fileName)
		stdGoMatch, stdGoDuration := measureGoStdRegex(data, rxEmail)
		re2CustomMatch, re2CustomDuration := measureRe2Go(data, re2EmailCustomCheck)
		re2SentinelMatch, re2SentinelDuration := measureRe2Go(data, re2EmailSentinel)

		fmt.Printf("File %q\n", fileName)
		fmt.Printf("  Go std regex        : found %d matches in %d ms\n", stdGoMatch, stdGoDuration)
		fmt.Printf("  re2go custom check  : found %d matches in %d ms (%dx slower)\n", re2CustomMatch, re2CustomDuration, re2CustomDuration/stdGoDuration)
		fmt.Printf("  re2go sentinel bound: found %d matches in %d ms (%dx slower)\n", re2SentinelMatch, re2SentinelDuration, re2SentinelDuration/stdGoDuration)
		fmt.Println()
	}
}

func openFile(srcPath string) []byte {
	data, err := os.ReadFile(srcPath)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func measureRe2Go(data []byte, finder func([]byte) int) (int, int64) {
	start := time.Now()
	nMatches := finder(data)
	elapsed := time.Since(start)
	return nMatches, elapsed.Milliseconds()
}

func measureGoStdRegex(data []byte, rx *regexp.Regexp) (int, int64) {
	start := time.Now()
	matches := rx.FindAll(data, -1)
	elapsed := time.Since(start)
	return len(matches), elapsed.Milliseconds()
}
