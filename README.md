# BlockList Updater And Merger
Downloads blocklists from a variety of sources and merges them in a single file for each category.

# Usage
Add the source and the category it belongs to in the blocking.json and then run `bluam` in the same directory.
It will update merge and then push, assuming you have git configured correctly, in the same repo the merged lists.

# Build
`go build .`
