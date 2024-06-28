package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func DownloadBlocklist(category string, url_link string) {
	// checks is ./dl_blocklists exists in not it create it
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)
	dl_path := filepath.Join(exPath, "dl_blocklists", category)
	if !CheckPathExists(dl_path) {
		err = os.MkdirAll(dl_path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	fullURLFile := url_link

	fileURL, err := url.Parse(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}

	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-2] + "_" + segments[len(segments)-1]

	// create the full path with the filname
	// maybe it will better if use an xdg standard location like the Downloads folder ???
	location := filepath.Join(dl_path, fileName)
	file, err := os.Create(location)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := client.Get(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, _ := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("Downloaded %s with size %d\n", fileName, size)
}
