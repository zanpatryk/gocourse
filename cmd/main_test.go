package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestDownloadFile(t *testing.T) {
	ogTimeout := reqTimeout
	defer func() {
		reqTimeout = ogTimeout
	}()

	tests := map[string]struct {
		handler     http.HandlerFunc
		modTimeout  time.Duration
		wantErr     error
		wantContent string
	}{
		"success 200": {
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				_, _ = w.Write([]byte("hello world"))
			},
			wantErr:     nil,
			wantContent: "hello world",
		},

		"not found 404": {

			handler: func(w http.ResponseWriter, r *http.Request) {
				http.NotFound(w, r)
			},
			wantErr: ErrFileNotFound,
		},

		"server error 500": {

			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(500)
				_, _ = w.Write([]byte("boom"))
			},
			wantErr: ErrDownloadFailed,
		},

		"timeout": {

			handler: func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(200 * time.Millisecond)
				w.WriteHeader(200)
				_, _ = w.Write([]byte("late"))
			},
			modTimeout: 100 * time.Millisecond,
			wantErr:    ErrConnectionFailed,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			srv := httptest.NewServer(tc.handler)
			defer srv.Close()

			if tc.modTimeout != 0 {
				reqTimeout = tc.modTimeout
			} else {
				reqTimeout = ogTimeout
			}

			outFile := filepath.Join(t.TempDir(), "out.txt")

			err := downloadFile(srv.URL+"/file", outFile)

			if tc.wantErr == nil {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				data, err := os.ReadFile(outFile)
				if err != nil {
					t.Fatalf("reading output file failed: %v", err)
				}
				if string(data) != tc.wantContent {
					t.Fatalf("content mismatch: got %q want %q", string(data), tc.wantContent)
				}
			} else {
				if err == nil {
					t.Fatalf("expected error %v, got nil", tc.wantErr)
				}
				if !errors.Is(err, tc.wantErr) {
					t.Fatalf("expected error %v, got %v", tc.wantErr, err)
				}
				if _, statErr := os.Stat(outFile); statErr == nil {
					t.Fatalf("expected no output file on error, but file exists: %s", outFile)
				}
			}
		})
	}

}

func TestInvalidUrl(t *testing.T) {
	outFile := filepath.Join(t.TempDir(), "out.txt")
	err := downloadFile("://bad-url", outFile)

	if err == nil {
		t.Fatalf("expected error for invalid url, got nil")
	}
	if !errors.Is(err, ErrInvalidURL) {
		t.Fatalf("expected ErrInvalidURL, got %v", err)
	}
}
