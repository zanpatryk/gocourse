package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	ErrInvalidURL       = errors.New("invalid url")
	ErrConnectionFailed = errors.New("connection failed")
	ErrDownloadFailed   = errors.New("download failed")
	ErrFileNotFound     = errors.New("file not found")
)

var reqTimeout = 30 * time.Second

func main() {
	urlPtr := flag.String("url", "", "url to file")
	outputPtr := flag.String("output", "", "name of file")

	flag.Parse()

	if *urlPtr == "" || *outputPtr == "" {
		flag.Usage()
		os.Exit(2)
	}

	err := downloadFile(*urlPtr, *outputPtr)
	if err == nil {
		os.Exit(0)
	}

	switch {
	case errors.Is(err, ErrInvalidURL):
		fmt.Printf("Error: invalid url. Please provide a valid http or https URL")
		os.Exit(1)
	case errors.Is(err, ErrConnectionFailed):
		fmt.Println("Erorr: connection failed. Retry or check your connection settings.")
		os.Exit(1)
	case errors.Is(err, ErrDownloadFailed):
		fmt.Println("Error: download failed. Check the availability of the file.")
		os.Exit(1)
	case errors.Is(err, ErrFileNotFound):
		fmt.Println("Error: file not found (404). Verify the URL points to an existing file.")
		os.Exit(1)
	default:
		fmt.Println("Unexpected error:", err)
		os.Exit(1)
	}

}

func downloadFile(fileURL, fileName string) error {
	parsed, err := url.ParseRequestURI(fileURL)

	if err != nil || !(parsed.Scheme == "http" || parsed.Scheme == "https") {
		return ErrInvalidURL
	}

	ctx, cancel := context.WithTimeout(context.Background(), reqTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fileURL, nil)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidURL, err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return ErrConnectionFailed
		}

		return fmt.Errorf("%w: %v", ErrConnectionFailed, err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusNotFound:
		return ErrFileNotFound
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return ErrDownloadFailed
	}

	out, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("Error creating file: %v", err)
	}

	defer func() {
		err = out.Close()
		if err != nil {
			fmt.Printf("Error while closing file: %v", err)
		}
	}()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrDownloadFailed, err)
	}

	return nil
}
