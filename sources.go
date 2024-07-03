package main

// add more []string variables and expand the urls_map with those for more categories and/or
// expand the already existing []string variables for more sources
// probably not the most convenient way but seems to work fine...
// keeping it all in the Pi-hole format

// some pi-hole style blocklists like from https://github.com/hagezi/dns-blocklists have sliglty different format they || at the start

// var (
// 	adultcontentblocklists = []string{
// 		"https://nsfw.oisd.nl/",
// 		"https://raw.githubusercontent.com/StevenBlack/hosts/master/alternates/porn-only/hosts",
// 	}
//
// 	cryptoblocklists = []string{}
//
// 	socialmediablocklists = []string{
// 		"https://raw.githubusercontent.com/StevenBlack/hosts/master/alternates/social-only/hosts",
// 	}
//
// 	surveillancebllocklists = []string{}
//
// 	adsblocklists = []string{
// 		"https://o0.pages.dev/Pro/domains.txt",
// 		"https://cdn.jsdelivr.net/gh/hagezi/dns-blocklists@latest/adblock/light.txt",
// 		"https://adaway.org/hosts.txt",
// 		"https://v.firebog.net/hosts/AdguardDNS.txt",
// 		"https://v.firebog.net/hosts/Admiral.txt",
// 		"https://raw.githubusercontent.com/anudeepND/blacklist/master/adservers.txt",
// 		"https://v.firebog.net/hosts/Easylist.txt",
// 		"https://raw.githubusercontent.com/FadeMind/hosts.extras/master/UncheckyAds/hosts",
// 		"https://raw.githubusercontent.com/bigdargon/hostsVN/master/hosts",
// 		"https://v.firebog.net/hosts/Easyprivacy.txt",
// 		"https://v.firebog.net/hosts/Prigent-Ads.txt",
// 		"https://raw.githubusercontent.com/FadeMind/hosts.extras/master/add.2o7Net/hosts",
// 		"https://raw.githubusercontent.com/crazy-max/WindowsSpyBlocker/master/data/hosts/spy.txt",
// 		"https://hostfiles.frogeye.fr/firstparty-trackers-hosts.txt",
// 	}
//
// 	drugsblocklists = []string{}
//
// 	fakenewsblocklists = []string{
// 		"https://raw.githubusercontent.com/StevenBlack/hosts/master/alternates/fakenews-only/hosts",
// 	}
//
// 	fraudblocklists = []string{
// 		"https://cdn.jsdelivr.net/gh/hagezi/dns-blocklists@latest/adblock/fake.txt",
// 	}
//
// 	gamblingblocklists = []string{
// 		"https://cdn.jsdelivr.net/gh/hagezi/dns-blocklists@latest/adblock/gambling.txt",
// 	}
//
// 	malwareblocklists = []string{
// 		"https://cdn.jsdelivr.net/gh/hagezi/dns-blocklists@latest/adblock/hoster.txt",
// 		"https://cdn.jsdelivr.net/gh/hagezi/dns-blocklists@latest/adblock/dyndns.txt",
// 		"https://cdn.jsdelivr.net/gh/hagezi/dns-blocklists@latest/adblock/tif.txt",
// 		"https://raw.githubusercontent.com/DandelionSprout/adfilt/master/Alternate%20versions%20Anti-Malware%20List/AntiMalwareHosts.txt",
// 		"https://osint.digitalside.it/Threat-Intel/lists/latestdomains.txt",
// 		"https://v.firebog.net/hosts/Prigent-Crypto.txt",
// 		"https://v.firebog.net/hosts/Prigent-Crypto.txt",
// 		"https://raw.githubusercontent.com/FadeMind/hosts.extras/master/add.Risk/hosts",
// 		"https://bitbucket.org/ethanr/dns-blacklists/raw/8575c9f96e5b4a1308f2f12394abd86d0927a4a0/bad_lists/Mandiant_APT1_Report_Appendix_D.txt",
// 		"https://phishing.army/download/phishing_army_blocklist_extended.txt",
// 		"https://gitlab.com/quidsup/notrack-blocklists/raw/master/notrack-malware.txt",
// 		"https://v.firebog.net/hosts/RPiList-Malware.txt",
// 		"https://v.firebog.net/hosts/RPiList-Phishing.txt",
// 		"https://raw.githubusercontent.com/Spam404/lists/master/main-blacklist.txt",
// 		"https://raw.githubusercontent.com/AssoEchap/stalkerware-indicators/master/generated/hosts",
// 		"https://urlhaus.abuse.ch/downloads/hostfile/",
// 	}
// )
//
// var urls_map = map[string][]string{
// 	"adultcontent": adultcontentblocklists,
// 	"crypto":       cryptoblocklists,
// 	"socialmedia":  socialmediablocklists,
// 	"surveillance": surveillancebllocklists,
// 	"adsblock":     adsblocklists,
// 	"drugs":        drugsblocklists,
// 	"fakenews":     fakenewsblocklists,
// 	"fraud":        fraudblocklists,
// 	"gambling":     gamblingblocklists,
// 	"malware":      malwareblocklists,
// }
