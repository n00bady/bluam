package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	dl_dir = "dl_blocklists"
)

func main() {
	fmt.Println("Updating the blocklists...")

	// Should i delete the older stuff in the dl_blocklists if they exist ?
	for category, url_links := range urls_map {
		for _, link := range url_links {
			DownloadBlocklist(category, link)
		}
	}

	// get exec path and from there find the dl_blocklists directory 
	// that all the blocklists are saved in
	exPath := filepath.Dir(FindExePath())
	dl_path := filepath.Join(exPath, dl_dir) 
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
		fmt.Println("\t", cat, "blocklists merged!")
		locations = nil
	}
}
