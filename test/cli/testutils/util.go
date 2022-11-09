package testutils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func SplitLines(s string) []string {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func MustOpen(name string) *os.File {
	f, err := os.Open(name)
	if err != nil {
		log.Panicf("opening %s: %s", name, err)
	}
	return f
}

// StrConcat takes a bunch of strings or string slices
// and concats them all together into one string slice.
// If an arg is not one of those types, this panics.
func StrConcat(args ...interface{}) []string {
	res := make([]string, 0)
	for _, a := range args {
		if s, ok := a.(string); ok {
			res = append(res, s)
			continue
		}
		if ss, ok := a.([]string); ok {
			res = append(res, ss...)
			continue
		}
		panic(fmt.Sprintf("arg '%v' must be a string or string slice, but is '%T'", a, a))
	}
	return res
}

// PreviewStr returns a preview of s, which is a prefix for logging that avoids dumping a huge string to logs.
func PreviewStr(s string) string {
	suffix := "..."
	previewLength := 10
	if len(s) < previewLength {
		previewLength = len(s)
		suffix = ""
	}
	return s[0:previewLength] + suffix
}