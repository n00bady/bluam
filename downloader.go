package main

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Maybe I should pass the entire blocklists struct instead ???
// Gets the category as a string and the url as a string and creates
// in ./dns/dl_blocklists/ subdirs for each category. Then donwloads the
// file from the url inside there
func DownloadBlocklist(dl_path string, url_link string) (err error) {
	if !CheckPathExists(dl_path) {
		err := os.MkdirAll(dl_path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// if it's a filepath then it just get copied to the dl_blocklists/
	// else I assume it's a URL and make a request
	fileName := encodeListURLToFileName(url_link)
	filePath := filepath.Join(dl_path, fileName)
	source, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer source.Close()

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := client.Get(url_link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	if strings.Contains(contentType, "text/plain") {
		_, err := io.Copy(source, resp.Body)
		return err
	} else {
		err = errors.New("file content not text/plain")
	}
	return nil
}
