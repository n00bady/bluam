#!/bin/bash

# where the lists will be downloaded 
mkdir -p "./dl_blocklists"
mkdir -p "./merged_lists"

# for merging the lists on the same categories
categories=("adblock" "domain" "host")

download_lists() {
	echo "Downloading blocklists..."
	sed '/^[ \t]*$/d' "./urls.txt" | while read -r url; do
		extracted_url_name=$(echo $url | sed 's|^\([^/]*/*\)\{5\}||; s|/|-|g')
		wget "$url" -O "./dl_blocklists/$extracted_url_name"
	done
}

# will make it loop through the categories later...
merge_lists() {
	echo "Merging adblock lists..."
	
	adblock_lists=()

	# this might be a little silly but it seems like it works
	while IFS= read -r -d '' file; do
		adblock_lists+=("$file")
	done < <(find "./dl_blocklists" -type f -name "*adblock*" -print0)

	
	# for f in "${adblock_lists[@]}"; do
	# 	echo "$f"
	# done

	# concatenate the lists
	cat "${adblock_lists[@]}" | sort | uniq > "./merged_lists/adblock_merged.txt"
}

download_lists

merge_lists
