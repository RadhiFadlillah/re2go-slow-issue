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
		re2goMatch, re2goDuration := measureRe2Go(data, re2Email)
		stdGoMatch, stdGoDuration := measureGoStdRegex(data, rxEmail)

		fmt.Printf("File %q\n", fileName)
		fmt.Printf("  re2go: found %d matches in %d ms\n", re2goMatch, re2goDuration)
		fmt.Printf("  std  : found %d matches in %d ms\n", stdGoMatch, stdGoDuration)
		fmt.Printf("  re2go is %dx slower than std\n", re2goDuration/stdGoDuration)
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
