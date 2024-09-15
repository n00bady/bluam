# BlockList Updater And Merger
Downloads blocklists from a variety of sources and merges them in a single file for each category.

# Usage
Add the source and the category it belongs to in the blocking.json and then run `bluam` in the same directory.
It will update merge and then push, assuming you have git configured correctly, in the same repo the merged lists.

# Build
`go build .`

# Sources 
- https://oisd.nl/ 
- https://github.com/StevenBlack/hosts
- https://github.com/badmojr/1Hosts/
- https://github.com/hagezi/dns-blocklists
- https://adaway.org/
- https://firebog.net/
- https://github.com/anudeepND/blacklist
- https://github.com/FadeMind/hosts.extras
- https://github.com/bigdargon/hostsVN
- https://github.com/crazy-max/WindowsSpyBlocker
- https://hostfiles.frogeye.fr/
- https://github.com/DandelionSprout/adfilt
- https://osint.digitalside.it/
- https://phishing.army/
- https://gitlab.com/quidsup/notrack-blocklists
- https://github.com/Spam404/lists
- https://github.com/AssoEchap/stalkerware-indicators
- https://urlhaus.abuse.ch
- And others... maybe
