package palindrome

import (
	"testing"
)

var tests = map[string]struct {
	input string
	want  bool
}{
	"simple":   {"civic", true},
	"fail":     {"fail", false},
	"numbers":  {"22\\2\\22", true},
	"sentence": {"Mr. Owl ate my metal worm", true},
}

func TestCheckWithStack(t *testing.T) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := CheckWithStack(tc.input)
			if got != tc.want {
				t.Fatalf("Wanted: %t, got: %t", tc.want, got)
			}
		})
	}
}

func TestCheckWithDoublePointer(t *testing.T) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := CheckWithDoublePointer(tc.input)
			if got != tc.want {
				t.Fatalf("Wanted: %t, got: %t", tc.want, got)
			}
		})
	}
}

func TestCheckWithReversedString(t *testing.T) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := CheckWithReversedString(tc.input)
			if got != tc.want {
				t.Fatalf("Wanted: %t, got: %t", tc.want, got)
			}
		})
	}
}

///// BENCHMARKS /////

func BenchmarkCheckWithStack(b *testing.B) {
	for b.Loop() {
		CheckWithStack("Mr. Owl ate my metal worm")
	}
}

func BenchmarkCheckWithDoublePointer(b *testing.B) {
	for b.Loop() {
		CheckWithDoublePointer("Mr. Owl ate my metal worm")
	}
}

func BenchmarkCheckWithReversedString(b *testing.B) {
	for b.Loop() {
		CheckWithReversedString("Mr. Owl ate my metal worm")
	}
}

/*
Benchmark results:

goos: darwin
goarch: arm64
pkg: github.com/zanpatryk/gocourse/palindrome
cpu: Apple M1 Pro
BenchmarkCheckWithStack-8                 332598              3520 ns/op
BenchmarkCheckWithDoublePointer-8        1174317              1021 ns/op
BenchmarkCheckWithReversedString-8       1000000              1112 ns/op

*/
