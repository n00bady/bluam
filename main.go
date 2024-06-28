package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

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
	fmt.Println("Merging blocklists...")
	for _, d := range subDirs {
		cat := d.Name()
		var locations []string
		if d.IsDir() {
			subItems, _ := os.ReadDir(filepath.Join(dl_path, d.Name()))
			for _, fN := range subItems {
				locations = append(locations, filepath.Join(dl_path, d.Name(), fN.Name()))
			}
		}
		MergeBlocklists(cat, locations)
		fmt.Println(cat, "blocklists merged!")
		locations = nil
	}
}
