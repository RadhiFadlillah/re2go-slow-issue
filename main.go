package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"golang.org/x/text/message"
)

var (
	msgp    = message.NewPrinter(message.MatchLanguage("en"))
	rxEmail = regexp.MustCompile(`[\+\-\.0-9A-Z_a-z]+@[\-\.0-9A-Z_a-z]+[\.][\-\.0-9A-Z_a-z]+`)
)

func main() {
	fileNames := []string{
		"sample/01.html",
		"sample/02.html",
		"sample/03.html",
		"sample/04.html",
	}

	for _, fileName := range fileNames {
		data := openFile(fileName)
		_, speed0 := measureGo(data, rxEmail)                                     // Std go library
		_, speed1, comparison1 := measureRe2Go(data, re2EmailCustomCheck, speed0) // re2go with custom bound check
		_, speed2, comparison2 := measureRe2Go(data, re2EmailSentinel, speed0)    // re2go with sentinel bound
		_, speed3, comparison3 := measureRe2Go(data, re2EmailSentinelBV, speed0)  // re2go with sentinel bound + bit vectors
		_, speed4, comparison4 := measureRe2Go(data, re2Email2Rule, speed0)       // re2go with 2 rules + sentinel bound
		_, speed5, comparison5 := measureRe2Go(data, re2Email2RuleBV, speed0)     // re2go with 2 rules + sentinel bound + bit vectors
		_, speed6, comparison6 := measureRe2Go(data, re2Email3Rule, speed0)       // re2go with 3 rules + sentinel bound
		_, speed7, comparison7 := measureRe2Go(data, re2Email3RuleBV, speed0)     // re2go with 3 rules + sentinel bound + bit vectors

		fmt.Printf("File %q\n", fileName)
		fmt.Printf("  Go std regex     : %s µs\n", mcs(speed0))
		fmt.Printf("  re2go custom     : %s µs (%s)\n", mcs(speed1), comparison1)
		fmt.Printf("  re2go sentinel   : %s µs (%s)\n", mcs(speed2), comparison2)
		fmt.Printf("  re2go sentinel BV: %s µs (%s)\n", mcs(speed3), comparison3)
		fmt.Printf("  re2go 2 rules    : %s µs (%s)\n", mcs(speed4), comparison4)
		fmt.Printf("  re2go 2 rules BV : %s µs (%s)\n", mcs(speed5), comparison5)
		fmt.Printf("  re2go 3 rules    : %s µs (%s)\n", mcs(speed6), comparison6)
		fmt.Printf("  re2go 3 rules BV : %s µs (%s)\n", mcs(speed7), comparison7)
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

func measureRe2Go(data []byte, finder func([]byte) int, goDuration time.Duration) (int, time.Duration, string) {
	start := time.Now()
	nMatches := finder(data)
	elapsed := time.Since(start)
	return nMatches, elapsed, comparison(goDuration, elapsed)
}

func measureGo(data []byte, rx *regexp.Regexp) (int, time.Duration) {
	start := time.Now()
	matches := rx.FindAll(data, -1)
	elapsed := time.Since(start)
	return len(matches), elapsed
}

func mcs(duration time.Duration) string {
	return msgp.Sprint(duration.Microseconds())
}

func comparison(goDuration, re2goDuration time.Duration) string {
	if goDuration < re2goDuration {
		return fmt.Sprintf("%dx slower", re2goDuration/goDuration)
	} else {
		return fmt.Sprintf("%dx faster", goDuration/re2goDuration)
	}
}
