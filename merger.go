package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func MergeBlocklists(category string, fileNames []string) {
	// create a directory for each category
	exPath := filepath.Dir(FindExePath())
	merged_dir := filepath.Join(exPath, "mergedLists", category)
	if !CheckPathExists(merged_dir) {
		err := os.MkdirAll(merged_dir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	// This might not be right :S
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

	// Write the map into the empty file
	for _, v := range merge_map {
		_, err := merged.WriteString(v + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
