package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type DNSConfig struct {
	LogBlocked bool        `json:"LogBlocked"`
	LogAll     bool        `json:"LogAll"`
	IP         string      `json:"IP"`
	Port       int         `json:"Port"`
	DNSServers []string    `json:"DNSServers"`
	ListPath   string      `json:"ListPath"`
	Lists      []blocklist `json:"Categories"`
}

type blocklist struct {
	Tag            string   `json:"Tag"`
	Sources        []string `json:"Sources"`
	MergedLocation string   `json:"MergedLocation"`
	Enabled        bool     `json:"Enabled"`
}

func LoadConfig(path string) *DNSConfig {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error reading config file: ", err)
	}

	var config DNSConfig
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatal("Error unmarshaling json: ", err)
	}

	return &config
}

func UpdateListsAndMergeTags(config *DNSConfig) {
	// imo it's better to overwrite the old blocklists
	// rather than delete and then re-download
	// because if one of the list sources is temporary
	// or permantly unavailable we can still use the last version
	for _, l := range config.Lists {
		for _, s := range l.Sources {
			DownloadBlocklist(l.Tag, s)
		}
	}

	// reads all the the files in each category in downloaded blocklists
	// directory and merge them
	dl_path := filepath.Join(config.ListPath, "dl_blocklists")
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

func DisableList(config *DNSConfig) {
}

func EnableList(config *DNSConfig) {
}

// keyb1nd makes this
func LoadLists(config *DNSConfig) {
}
