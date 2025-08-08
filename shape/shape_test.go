package shape

import (
	"math"
	"testing"
)

func TestCircleArea(t *testing.T) {
	tests := map[string]struct {
		radius float64
		want   float64
	}{
		"zero radius": {radius: 0, want: 0},
		"radius 1":    {radius: 1, want: math.Pi},
		"radius 5:":   {radius: 5, want: math.Pi * 5 * 5},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			c := Circle{Radius: tc.radius}
			got := c.Area()
			if got != tc.want {
				t.Errorf("Wanted: %v, got: %v", tc.want, got)
			}
		})
	}
}

func TestRectangleeArea(t *testing.T) {
	tests := map[string]struct {
		width  float64
		height float64
		want   float64
	}{
		"zero width/height":       {width: 0, height: 0, want: 0},
		"same width/height":       {width: 5, height: 5, want: 25},
		"different width/height:": {width: 10, height: 8, want: 80},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r := Rectangle{Width: tc.width, Height: tc.height}
			got := r.Area()
			if got != tc.want {
				t.Errorf("Wanted: %v, got: %v", tc.want, got)
			}
		})
	}
}
