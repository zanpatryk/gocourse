package set

import (
	"reflect"
	"testing"
)

func TestCreateSet(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  []string
	}{
		"empty slice":      {[]string{}, []string{}},
		"no duplicates":    {[]string{"a", "b", "c"}, []string{"a", "b", "c"}},
		"all duplicates":   {[]string{"a", "a", "a"}, []string{"a"}},
		"mixed duplicates": {[]string{"a", "b", "a", "c", "b"}, []string{"a", "b", "c"}},
		"single element":   {[]string{"only"}, []string{"only"}},
		"long list":        {[]string{"one", "two", "three"}, []string{"one", "two", "three"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := CreateSet(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Wanted: %#v, got: %#v", tc.want, got)
			}
		})
	}
}
