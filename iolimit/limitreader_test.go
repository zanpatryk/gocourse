package iolimit

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

func TestLimitReader(t *testing.T) {
	tests := map[string]struct {
		input          string
		limit          int64
		readChunkSizes []int // nil => use io.ReadAll; otherwise read with these chunk sizes
		wantData       string
		wantErr        error  // expected final error: nil or io.EOF
		wantRemaining  string // what's left in underlying bytes.Buffer after reads
	}{
		"limit 0": {
			input:          "hello",
			limit:          0,
			readChunkSizes: nil,
			wantData:       "",
			wantErr:        nil,
			wantRemaining:  "hello",
		},
		"limit < length": {
			input:          "abcdef",
			limit:          3,
			readChunkSizes: nil,
			wantData:       "abc",
			wantErr:        nil,
			wantRemaining:  "def",
		},
		"limit == length": {
			input:          "xyz",
			limit:          3,
			readChunkSizes: nil,
			wantData:       "xyz",
			wantErr:        nil,
			wantRemaining:  "",
		},
		"limit > length": {
			input:          "ok",
			limit:          10,
			readChunkSizes: nil,
			wantData:       "ok",
			wantErr:        nil,
			wantRemaining:  "",
		},
		"partial read small chunks": {
			input:          "012345",
			limit:          4,
			readChunkSizes: []int{1, 1, 1, 1, 1},
			wantData:       "0123",
			wantErr:        io.EOF,
			wantRemaining:  "45",
		},
		"large read slice but shorter limit": {
			input:          "longcontent",
			limit:          4,
			readChunkSizes: []int{100},
			wantData:       "long",
			wantErr:        io.EOF,
			wantRemaining:  "content",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			underlying := bytes.NewBufferString(tc.input)
			lr := LimitReader(underlying, tc.limit)

			var got []byte
			var gotErr error

			if tc.readChunkSizes == nil {
				got, gotErr = io.ReadAll(lr)
			} else {
				var collected bytes.Buffer
				for _, size := range tc.readChunkSizes {
					p := make([]byte, size)
					n, err := lr.Read(p)
					if n > 0 {
						collected.Write(p[:n])
					}
					if err != nil {
						gotErr = err
						break
					}
				}
				if gotErr == nil {
					_, gotErr = lr.Read(make([]byte, 1))
				}
				got = collected.Bytes()
			}

			// verify returned bytes match expectation
			if string(got) != tc.wantData {
				t.Fatalf("[%s] returned data = %q; want %q", name, string(got), tc.wantData)
			}

			// verify error expectations:
			// - for io.ReadAll paths - nil
			// - for chunked-read paths - io.EOF
			if tc.wantErr == nil {
				if gotErr != nil && !errors.Is(gotErr, io.EOF) {
					t.Fatalf("[%s] got error = %v; want nil (or io.EOF accepted for chunk mode)", name, gotErr)
				}
			} else if errors.Is(tc.wantErr, io.EOF) {
				if !errors.Is(gotErr, io.EOF) {
					t.Fatalf("[%s] got error = %v; want io.EOF", name, gotErr)
				}
			}

			// verify the underlying buffer's unread contents equal expected remainder
			if underlying.String() != tc.wantRemaining {
				t.Fatalf("[%s] underlying remaining = %q; want %q", name, underlying.String(), tc.wantRemaining)
			}
		})
	}
}
