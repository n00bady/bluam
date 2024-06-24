package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var url_link string

func main() {
	fmt.Println("Updating the blocklists...")

	// opens the urls.txt file and takes each line as a url to download the file
	f, err := os.OpenFile("./urls.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fScanner := bufio.NewScanner(f)
	fScanner.Split(bufio.ScanLines)
	for fScanner.Scan() {
		url_link = fScanner.Text()
		if strings.TrimSpace(url_link) == "" {
			continue
		}
		fmt.Println("Downloading ", url_link, "...")
		DownloadBlocklist(url_link)
		fmt.Println("Finished downloading ", url_link)
	}
}

func DownloadBlocklist(url_link string) {
	// checks is ./dl_blocklists exists in not it create it
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)
	dl_path := filepath.Join(exPath, "dl_blocklists")
	if !checkPathExist(dl_path) {
		err = os.Mkdir(dl_path, os.ModePerm)
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

	fmt.Printf("Downloaded %s with size %d", fileName, size)
}

func checkPathExist(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func MergeBlocklists() {
}
