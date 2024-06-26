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

// add more []string variables and expand the urls_map with those for more categories and/or
// expand the already []string variables for more sources
// probably not the most convenient way but seems to work fine...
var adblocklists = []string{
	"https://github.com/badmojr/1Hosts/releases/download/latest/1hosts-Pro_adblock.txt",
	"https://cdn.jsdelivr.net/gh/hagezi/dns-blocklists@latest/adblock/light.txt",
}

var domainsblocklists = []string{
	"https://github.com/badmojr/1Hosts/releases/download/latest/1hosts-Pro_domains.txt",
	"https://cdn.jsdelivr.net/gh/hagezi/dns-blocklists@latest/domains/light.txt",
}

var hostsblocklists = []string{
	"https://github.com/badmojr/1Hosts/releases/download/latest/1hosts-Pro_hosts.txt",
	"https://cdn.jsdelivr.net/gh/hagezi/dns-blocklists@latest/hosts/light.txt",
}

var urls_map = map[string][]string{
	"adblock": adblocklists,
	"domains": domainsblocklists,
	"hosts":   hostsblocklists,
}

func main() {
	fmt.Println("Updating the blocklists...")

	for category, url_links := range urls_map {
		for _, link := range url_links {
			DownloadBlocklist(category, link)
		}
	}

	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)
	dl_path := filepath.Join(exPath, "dl_blocklists")
	subDirs, _ := os.ReadDir(dl_path)
	for _, d := range subDirs {
		fmt.Println("d = ", d)
		cat := d.Name()
		var locations []string
		if d.IsDir() {
			subItems, _ := os.ReadDir(filepath.Join(dl_path, d.Name()))
			for _, fN := range subItems {
				locations = append(locations, filepath.Join(dl_path, d.Name(), fN.Name()))
			}
		}
		MergeBlocklists(cat, locations)
		locations = nil
	}
}

func DownloadBlocklist(category string, url_link string) {
	// checks is ./dl_blocklists exists in not it create it
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)
	dl_path := filepath.Join(exPath, "dl_blocklists", category)
	if !checkPathExist(dl_path) {
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

func MergeBlocklists(category string, fileNames []string) {
	fmt.Println(fileNames)
	// create a directory for each category
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)
	merged_dir := filepath.Join(exPath, "mergedLists", category)
	if !checkPathExist(merged_dir) {
		err = os.MkdirAll(merged_dir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	merge_map := make(map[int]string)
	outIndex := 0
	for _, fN := range fileNames {
		// read files
		line := 0
		f, err := os.Open(fN)
		if err != nil {
			log.Fatal(err)
		}
		fScanner := bufio.NewScanner(f)
		fScanner.Split(bufio.ScanLines)
		for fScanner.Scan() {
			entry := fScanner.Text()
			if merge_map[line] == entry {
				continue
			} else {
				merge_map[outIndex] = entry
			}
			line++
			outIndex++
		}
		defer f.Close()
	}

	// Create an empty output file
	fileName := category + ".txt"
	location := filepath.Join(merged_dir, fileName)
	merged, err := os.Create(location)
	if err != nil {
		log.Fatal(err)
	}
	defer merged.Close()

	for _, v := range merge_map {
		_, err := merged.WriteString(v + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func checkPathExist(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}
