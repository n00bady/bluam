package main

var config *DNSConfig

func main() {
	config = LoadConfig("./blocking.json")
	UpdateListsAndMergeTags(config)
}
