#!/bin/bash

mkdir -p "./dl_blocklists"

categories=("adblock" "domain" "host")

download_lists() {
	echo "Downloading blocklists..."
	sed '/^[ \t]*$/d' "./urls.txt" | while read -r url; do
		extracted_url_name=$(echo $url | sed 's|^\([^/]*/*\)\{5\}||; s|/|-|g')
		wget "$url" -O "./dl_blocklists/$extracted_url_name"
	done
}

merge_lists() {
	echo "Merging adblock lists..."
	
	adblock_lists=()

	while IFS= read -r -d '' file; do
		adblock_lists+=("$file")
	done < <(find "./dl_blocklists" -type f -name "*adblock*" -print0)

	
	for f in "${adblock_lists[@]}"; do
		echo "$f"
	done

	cat "${adblock_lists[@]}" | sort | uniq > "./adblock_merged.txt"
}

download_lists

merge_lists
