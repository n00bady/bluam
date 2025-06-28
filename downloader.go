package main

import (
	"io"
	"net/http"
)

// Takes a URL and returns as string the body content
func DownloadBlocklist(url string) (string, error) {
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			req.URL.Opaque = req.URL.Path
			return nil
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// contentType := resp.Header.Get("Content-Type")
	// if !strings.Contains(contentType, "text/plain") {
	// return "", errors.New("file content not text/plain")
	// }

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
