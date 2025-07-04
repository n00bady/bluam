package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"strings"
	"time"
)

type DNSConfig struct {
	Sources []Source `json:"Sources"`
}

type Source struct {
	Updated  time.Time `json:"Updated"`
	Category string    `json:"Category"`
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

func UpdateListsAndMergeTags(config *DNSConfig, path string) (err error) {
	defer func() {
		r := recover()
		if r != nil {
			log.Println(r, string(debug.Stack()))
		}
		if err != nil && WEBHOOK != "" {
			SEND_ADMIN_ALERT("UpdateListsAndMerge > err: " + err.Error())
		}
	}()
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
		var rb string
		rb, err = DownloadBlocklist(s.Source)
		if err != nil {
			fmt.Println(err)
			continue
		}

		lines := strings.Split(rb, "\n")
		for _, l := range lines {
			categoryMap[s.Category][toPlainDomain(l)] = struct{}{}
		}
	}

	cmd := exec.Command("rm", "-R", "./dns/merged.bak")
	var out []byte
	out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("RMOUT:", string(out))
		// return err
	}

	cmd = exec.Command("cp", "-R", "./dns/merged", "./dns/merged.bak")
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("CPOUT:", string(out))
		// return err
	}

	err = os.MkdirAll("./dns/merged", os.ModePerm)
	if err != nil {
		return err
	}

	// filter uniques across categories

	fmt.Println("starting de-dupping.. this will take some time..")
	start := time.Now()
	dupeCount := 0
	un := make(map[string]struct{})
	for _, inMap := range categoryMap {
		for d := range inMap {
			_, ok := un[d]
			if ok {
				dupeCount++
				delete(inMap, d)
			} else {
				un[d] = struct{}{}
			}
		}
	}
	fmt.Println("FILTERING TIME:", time.Since(start).Minutes(), "DUPE count:", dupeCount)

	// Ok this was harder than it had to be :S
	// Iterate over the map and inner map then create a slice of strings and
	// iterate over the inner map and append them then your sort.Strings() to sort them
	// create the file to be saved and open it for each category and then write it line by line
	for cat, inMap := range categoryMap {
		fmt.Printf("Merging %s...", cat)
		// domains := make([]string, 0, len(inMap))
		// for domain := range inMap {
		// 	domains = append(domains, domain)
		// }
		// sort.Strings(domains)

		fileName := cat + ".txt"
		location := filepath.Join("./dns/merged", fileName)
		var f *os.File
		f, err = os.Create(location)
		if err != nil {
			return err
		}
		defer f.Close()

		for d := range inMap {
			if d == "" {
				continue
			}
			_, err = f.WriteString(d + "\n")
			if err != nil {
				return err
			}
		}
		fmt.Println("...done!")
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

func AddList(category, source string, config *DNSConfig) error {
	var newList Source

	newList.Updated = time.Now()
	newList.Category = category
	newList.Source = source

	config.Sources = append(config.Sources, newList)

	err := updateConfigFile(config)
	if err != nil {
		return err
	}

	return nil
}

func RemoveList(searchStr string, config *DNSConfig) error {
	fmt.Println("Searching... ")
	for i, s := range config.Sources {
		if s.Source == searchStr {
			fmt.Println("FOUND!!!")
			config.Sources = append(
				config.Sources[:i], config.Sources[i+1:]...)
			err := updateConfigFile(config)
			if err != nil {
				return err
			}

			fmt.Println(s.Source, " removed!")

			return nil
		}
	}

	return errors.New("the Source doesn't exists")
}

func RemoveCategory(category string, config *DNSConfig) error {
	fmt.Println("Searching... ")
	for i, s := range config.Sources {
		if s.Category == category {
			config.Sources = append(config.Sources[:i], config.Sources[i+1:]...)
		}
	}
	err := updateConfigFile(config)
	if err != nil {
		return err
	}

	fmt.Println(category, " removed!")

	return nil
}

func updateConfigFile(config *DNSConfig) error {
	f, err := os.OpenFile("./blocking.json", os.O_RDWR|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	jason, err := json.MarshalIndent(config, "", "	")
	if err != nil {
		return err
	}

	f.Write(jason)

	return nil
}
