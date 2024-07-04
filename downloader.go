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

// Maybe I should pass the entire blocklists struct instead ???
// Gets the category as a string and the url as a string and creates
// in ./dns/dl_blocklists/ subdirs for each category. Then donwloads the
// file from the url inside there
func DownloadBlocklist(category string, url_link string) {
	dl_path := filepath.Join(config.ListPath, "dl_blocklists", category)
	if !CheckPathExists(dl_path) {
		err := os.MkdirAll(dl_path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	// if it's a filepath then it just get copied to the dl_blocklists/
	// else I assume it's a URL and make a request
	if CheckPathExists(url_link) {
		fmt.Println("It's a local file!")
		source, err := os.Open(url_link)
		if err != nil {
			log.Fatal(err)
		}
		defer source.Close()

		dest, err := os.Create(filepath.Join(dl_path, source.Name()))
		if err != nil {
			log.Fatal(err)
		}
		defer dest.Close()

		size, err := io.Copy(dest, source)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Copied %s in %s with size %d\n", url_link, dl_path+source.Name(), size)
	} else {
		fmt.Println("It's a URL!")
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
		defer file.Close()

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

		// check if it's a text file and if not skip it
		// Is there a better way to check it ???
		contentType := resp.Header.Get("Content-Type")
		fmt.Println("Content-Type: ", contentType)
		if strings.Contains(contentType, "text/plain") {
			size, _ := io.Copy(file, resp.Body)
			fmt.Printf("Downloaded %s with size %d\n\n", fileName, size)
		} else {
			fmt.Printf("Skiping %s not a text file!\n\n", url_link)
		}
	}
}
