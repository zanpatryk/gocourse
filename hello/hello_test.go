package hello

import (
	"reflect"
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World!"

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Expected: %v, got: %v", want, got)
	}
}
