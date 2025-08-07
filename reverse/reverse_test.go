package reverse

import (
	"testing"
)

var tests = map[string]struct {
	input string
	want  string
}{
	"empty string": {
		input: "",
		want:  "",
	},
	"single word": {
		input: "hello",
		want:  "olleh",
	},
	"two words": {
		input: "go lang",
		want:  "og gnal",
	},
	"multiple spaces": {
		input: "foo  bar",
		want:  "oof  rab",
	},
	"unicode characters": {
		input: "こんにちは 世界",
		want:  "はちにんこ 界世",
	},
}

func TestReverseCharactersOrderRaw(t *testing.T) {

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := ReverseCharactersOrderRaw(tc.input)
			if got != tc.want {
				t.Fatalf("Wanted: %#v, got: %#v", tc.want, got)
			}
		})
	}
}

func TestReverseCharactersOrder(t *testing.T) {

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := ReverseCharactersOrder(tc.input)
			if got != tc.want {
				t.Fatalf("Wanted: %#v, got: %#v", tc.want, got)
			}
		})
	}
}

//// BENCHMARKS /////

func BenchmarkReverseCharactersOrderRaw(b *testing.B) {

	for b.Loop() {
		ReverseCharactersOrderRaw("Mr. Owl ate my metal worm")
	}
}

func BenchmarkReverseCharactersOrder(b *testing.B) {

	for b.Loop() {
		ReverseCharactersOrder("Mr. Owl ate my metal worm")
	}
}

/*
Benchmarks results:

goos: darwin
goarch: arm64
pkg: github.com/zanpatryk/gocourse/reverse
cpu: Apple M1 Pro
BenchmarkReverseCharactersOrderRaw-8     4805062               243.9 ns/op
BenchmarkReverseCharactersOrder-8        2923836               408.3 ns/op

*/
