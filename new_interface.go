package main

import (
	"encoding/json"
	"fmt"
	"log"
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

func UpdateConfigAndMergeTags(config *DNSConfig, path string) {
	// do re-merge and config update without download
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
	}

	return nil
}

// func UpdateListsAndMergeTags(config *DNSConfig, path string) {
// 	// imo it's better to overwrite the old blocklists
// 	// rather than delete and then re-download
// 	// because if one of the list sources is temporary
// 	// or permantly unavailable we can still use the last version
// 	dlPath := filepath.Join(path, "originals")
// 	for _, l := range config.Sources {
// 		DownloadBlocklist(dlPath, l.Source)
// 	}
//
// 	// reads all the the files in each category in downloaded blocklists
// 	// directory and merge them
// 	fmt.Println("Merging blocklists...")
// 	categoryMap := make(map[string][]string)
// 	for _, v := range config.Sources {
// 		sourcePath := encodeListURLToFileName(v.Source)
// 		categoryMap[v.Category] = append(categoryMap[v.Category], sourcePath)
// 	}
//
// 	mergePath := filepath.Join(path, "merged")
// 	for i, v := range categoryMap {
// 		MergeBlocklists(mergePath, i, v)
// 		fmt.Println("\t", i, "blocklists merged!")
// 	}
// }

func DisableList(config *DNSConfig, categories []string) {
	for _, c := range categories {
		for i, l := range config.Sources {
			if l.Category == c {
				fmt.Println("Found category: ", c, "Disabling...")
				config.Sources[i].Enabled = false
				bytes, err := json.MarshalIndent(config, "", "\t")
				if err != nil {
					log.Fatal(err)
				}

				f, err := os.Create(configPath)
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()
				_, err = f.Write(bytes)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(c, "disabled succefully!")
			}
		}
	}
}

func EnableList(config *DNSConfig, categories []string) {
	for _, c := range categories {
		for i, l := range config.Sources {
			if l.Category == c {
				fmt.Println("Found category: ", c, "Enabling...")
				config.Sources[i].Enabled = true
				bytes, err := json.MarshalIndent(config, "", "\t")
				if err != nil {
					log.Fatal(err)
				}

				f, err := os.Create(configPath)
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()
				_, err = f.Write(bytes)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(c, "enabled succefully!")
			}
		}
	}
}

// keyb1nd makes this
func LoadLists(config *DNSConfig) {
}
