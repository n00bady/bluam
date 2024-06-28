package main

// add more []string variables and expand the urls_map with those for more categories and/or
// expand the already existing []string variables for more sources
// probably not the most convenient way but seems to work fine...

var adblocklists = []string{
	"https://github.com/badmojr/1Hosts/releases/download/latest/1hosts-Pro_adblock.txt",
	"https://cdn.jsdelivr.net/gh/hagezi/dns-blocklists@latest/adblock/light.txt",
}

var domainsblocklists = []string{
	"https://github.com/badmojr/1Hosts/releases/download/latest/1hosts-Pro_domains.txt",
	"https://cdn.jsdelivr.net/gh/hagezi/dns-blocklists@latest/domains/light.txt",
}

var hostsblocklists = []string{
	"https://github.com/badmojr/1Hosts/releases/download/latest/1hosts-Pro_hosts.txt",
	"https://cdn.jsdelivr.net/gh/hagezi/dns-blocklists@latest/hosts/light.txt",
}

var urls_map = map[string][]string{
	"adblock": adblocklists,
	"domains": domainsblocklists,
	"hosts":   hostsblocklists,
}
