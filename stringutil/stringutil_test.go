package stringutil

import (
	"strings"
	"testing"
)

var tests = []struct {
	haystack string
	needle   string
	want     bool
}{
	{"this is a test", "is a", true},
	{"this is a test", "IS A", true},
	{"this is a test", "iS a", true},
	{"THIS IS A TEST", "is a", true},
	{"THIS IS A TEST", "IS A", true},
	{"THIS IS A TEST", "iS a", true},

	{"this is a test", "t est", false},
	{"this is a test", "t est", false},
	{"this is a test", "t est", false},
	{"THIS IS A TEST", "t est", false},
	{"THIS IS A TEST", "t est", false},
	{"THIS IS A TEST", "t est", false},

	{"", "t est", false},
	{"THIS IS A TEST", "", false},
	{"", "", false},

	{"hálló", "álló", true},
	{"hálló", "ÁLLÓ", true},
	{"hálló", "álLó", true},

	{"hálló", "allo", false},
	{"hálló", "Allo", false},
	{"hálló", "ALLO", false},
}

func TestContainsIgnoreCase(t *testing.T) {
	for _, test := range tests {
		if res := ContainsIgnoreCase(test.haystack, test.needle); res != test.want {
			t.Errorf("ContainsIgnoreCase(%q, %q) = %t", test.haystack, test.needle, res)
		}
	}
}

func TestContainsIgnoreCaseV2(t *testing.T) {
	for _, test := range tests {
		if res := ContainsIgnoreCaseV2(test.haystack, test.needle); res != test.want {
			t.Errorf("ContainsIgnoreCaseV2(%q, %q) = %t", test.haystack, test.needle, res)
		}
	}
}

func BenchmarkContainsIgnoreCase(b *testing.B) {
	haystack, needle := "this is a test", "iS a"

	for i := 0; i < b.N; i++ {
		ContainsIgnoreCase(haystack, needle)
	}
}

func BenchmarkContainsIgnoreCaseV2(b *testing.B) {
	haystack, needle := "this is a test", "iS a"

	for i := 0; i < b.N; i++ {
		ContainsIgnoreCaseV2(haystack, needle)
	}
}

func BenchmarkStringsContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Contains("this is a test", "is a")
	}
}
