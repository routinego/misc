package stringutil

import "regexp"
import "strings"

// ContainsIgnoreCase returns true if needle is found in
// haystack in an ignore-case fashion
func ContainsIgnoreCase(haystack, needle string) bool {
	if haystack == "" || needle == "" {
		return false
	}

	m := regexp.MustCompile("(?i)" + needle).FindString(haystack)

	return m != ""
}

// ContainsIgnoreCaseV2 uses a slightly different approach
// and is written to check if there's a speed improvement
func ContainsIgnoreCaseV2(haystack, needle string) bool {
	if haystack == "" || needle == "" {
		return false
	}

	haystack, needle = strings.ToLower(haystack), strings.ToLower(needle)

	return strings.Contains(haystack, needle)
}
