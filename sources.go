package main

// add more []string variables and expand the urls_map with those for more categories and/or
// expand the already existing []string variables for more sources
// probably not the most convenient way but seems to work fine...

var (
	adultcontentblocklists = []string{
		
	}

	cryptoblocklists = []string{

	}

	socialmediablocklists = []string{

	}

	surveillancebllocklists = []string{

	}

	adsblocklists = []string{
		"https://github.com/badmojr/1Hosts/releases/download/latest/1hosts-Pro_adblock.txt",
		"https://cdn.jsdelivr.net/gh/hagezi/dns-blocklists@latest/adblock/light.txt",
	}

	drugsblocklists = []string{

	}

	fakenewsblocklists = []string{

	}

	fraudblocklists = []string{

	}

	gamblingblocklists = []string{

	}

	malwareblocklists = []string{

	}
)

var urls_map = map[string][]string{
	"adultcontent": adultcontentblocklists,
	"crypto": cryptoblocklists,
	"socialmedia": socialmediablocklists,
	"surveillance": surveillancebllocklists,
	"adsblock": adsblocklists,
	"drugs": drugsblocklists,
	"fakenews": fakenewsblocklists,
	"fraud": fraudblocklists,
	"gambling": gamblingblocklists,
	"malware": malwareblocklists,
}
