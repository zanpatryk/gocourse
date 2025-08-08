package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/zanpatryk/gocourse/shape"
)

func TestRun_TableDriven(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantExit   int
		wantOutput string
	}{
		{"missing shape", []string{}, 1, "[ERROR] -shape flag is required"},
		{"invalid shape", []string{"-shape", "triangle"}, 1, "[ERROR] Invalid shape!"},
		{"rectangle default", []string{"-shape", "rectangle"}, 0, "Rectangle Area: 1.0000"},
		{"rectangle custom", []string{"-shape", "rectangle", "-width", "3", "-height", "2"}, 0, "Rectangle Area: 6.0000"},
		{"rectangle negative", []string{"-shape", "rectangle", "-width", "-1", "-height", "2"}, 1, "[ERROR] -width and -height cannot be less than 0"},
		{"circle success", []string{"-shape", "circle", "-radius", "2"}, 0, fmt.Sprintf("Circle Area: %.4f", shape.Circle{Radius: 2}.Area())},
		{"circle negative", []string{"-shape", "circle", "-radius", "-2"}, 1, "[ERROR] -radius cannot be less than 0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			got := run(tt.args, &buf)
			out := strings.TrimSpace(buf.String())

			if got != tt.wantExit {
				t.Fatalf("exit = %d; want %d; output=%q", got, tt.wantExit, out)
			}
			if tt.wantOutput != "" && !strings.Contains(out, tt.wantOutput) {
				t.Fatalf("output = %q; want to contain %q", out, tt.wantOutput)
			}
		})
	}
}
