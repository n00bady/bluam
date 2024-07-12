package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

// Take a a string (category) and a slice of strings for the filenames
// creates a new file with the name of the category in ./dns/merged_Lists ,
// iterates of the merge_map and writes in that file.
func MergeBlocklists(path string, category string, fileNames []string) {
	if !CheckPathExists(path) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	// read each file line by line and creates a uniq map with each line as key
	// also the toPlainDomain() function cleans most of the prefixes that different
	// blocklists tend to have like # and ! for comments/headers and || for all subdomains
	merge_map := make(map[string]struct{})
	for _, fN := range fileNames {
		f, err := os.Open(filepath.Join("./dns/originals", fN))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		fScanner := bufio.NewScanner(f)
		fScanner.Split(bufio.ScanLines)
		for fScanner.Scan() {
			entry := fScanner.Text()
			merge_map[toPlainDomain(entry)] = struct{}{}
		}
	}

	// Create an empty output file
	fileName := category + ".txt"
	location := filepath.Join(path, fileName)
	merged, err := os.Create(location)
	if err != nil {
		log.Fatal(err)
	}
	defer merged.Close()

	// Write the map into the empty file
	// Should we care about writing them in alphabetical order ???
	for k := range merge_map {
		_, err := merged.WriteString(k + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
