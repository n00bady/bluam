package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type DNSConfig struct {
	LogBlocked         bool     `json:"LogBlocked"`
	LogAll             bool     `json:"LogAll"`
	IP                 string   `json:"IP"`
	Port               int      `json:"Port"`
	DNSServers         []string `json:"DNSServers"`
	DisabledCategories []string `json:"DisabledCategories"`
	AutoUpdate         bool     `json:"AutoUpdate"`
	Sources            []Source `json:"Sources"`
}

type Source struct {
	Updated  time.Time `json:"Updated"`
	Category string    `json:"Category"`
	Enabled  bool      `json:"Enabled"`
	Source   string    `json:"Source"`
}

func LoadConfig(path string) (*DNSConfig, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config DNSConfig
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func UpdateListsAndMergeTags(config *DNSConfig, path string) error {
	categoryMap := make(map[string]map[string]struct{})
	categories := []string{"adult", "crypto", "socialmedia", "surveillance", "ads", "drugs", "fakenews", "fraud", "gambling", "malware"}

	for _, c := range categories {
		categoryMap[c] = make(map[string]struct{})
	}

	for _, s := range config.Sources {
		if s.Source == "" {
			continue
		}
		fmt.Println("Downloading: ", s.Source)
		responseBody, err := DownloadBlocklist(s.Source)
		if err != nil {
			return err
		}

		lines := strings.Split(responseBody, "\n")
		for _, l := range lines {
			categoryMap[s.Category][toPlainDomain(l)] = struct{}{}
		}
	}

	err := os.MkdirAll("./dns/merged", os.ModePerm)
	if err != nil {
		return err
	}

	// Ok this was harder than it had to be :S
	// Iterate over the map and inner map then create a slice of strings and
	// iterate over the inner map and append them then your sort.Strings() to sort them
	// create the file to be saved and open it for each category and then write it line by line
	for cat, inMap := range categoryMap {
		fmt.Printf("Merging %s...", cat)
		domains := make([]string, 0, len(inMap))
		for domain := range inMap {
			domains = append(domains, domain)
		}
		sort.Strings(domains)

		fileName := cat + ".txt"
		location := filepath.Join("./dns/merged", fileName)
		f, err := os.Create(location)
		if err != nil {
			return err
		}
		defer f.Close()

		for _, d := range domains {
			_, err := f.WriteString(d + "\n")
			if err != nil {
				return err
			}
		}
		fmt.Println("...done!")
	}

	err = gitAddCommitPushLists()
	if err != nil {
		return err
	}

	return nil
}

// assumes you already have git configured properly for the repo
func gitAddCommitPushLists() error {
	changed, err := blocklistsChanged()
	fmt.Println("Changed: ", changed)
	if err != nil {
		return err
	}
	if !changed {
		return errors.New("no changes detected in dns directory, aborting git commit and push")
	}

	err = runCmd("git", "add", "dns")
	if err != nil {
		return err
	}

	err = runCmd("git", "commit", "-m", "Blocklists Update.")
	if err != nil {
		return err
	}

	err = runCmd("git", "push")
	if err != nil {
		return err
	}

	return nil
}
