package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func MergeBlocklists(category string, fileNames []string) {
	// create a directory for each category
	// maybe I should just dump them all in the mergedLists dir and not create subdirs
	// exPath := filepath.Dir(FindExePath())
	// merged_dir := filepath.Join(exPath, "mergedLists", category)
	merged_dir := filepath.Join(config.ListPath, "merged_Lists")
	err := os.Mkdir(merged_dir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	if !CheckPathExists(merged_dir) {
		err := os.Mkdir(merged_dir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	// read each file line by line and creates a uniq map with each line as key
	// also the toPlainDomain() function cleans most of the prefixes that different
	// blocklists tend to have like # and ! for comments/headers and || for all subdomains
	merge_map := make(map[string]bool)
	for _, fN := range fileNames {
		f, err := os.Open(fN)
		if err != nil {
			log.Fatal(err)
		}
		fScanner := bufio.NewScanner(f)
		fScanner.Split(bufio.ScanLines)
		for fScanner.Scan() {
			entry := fScanner.Text()
			merge_map[toPlainDomain(entry)] = true
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

	// Write the map into the empty file
	// Should we care about writing them in alphabetical order ???
	for k := range merge_map {
		_, err := merged.WriteString(k + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
