package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const usgMsg = "Running without arguments Updates, Merges, Commits and Pushes the blocklists.\n\n" +
	"Commands:\n\n" +
	"\t update Updates, Merges, Commits and Pushes the blocklists.\n" +
	"\t\t -noPush Stops it from Commiting and Pushing.\n" +
	"\n" +
	"\t add -c <category> <blocklists> Adds the following blocklists to the config.\n" +
	"\t remove <blocklist> Removes this list.\n" +
	"\t remove -c <category> Removes all lists of the category.\n" +
	"\n" +
	"The blocklists must be given with their full Path or URL!\n"

var WEBHOOK = ""

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Cannot load .env: ", err)
	} else {
		WEBHOOK = os.Getenv("WEBHOOK")
	}

	// load the config first thing!
	config, err := LoadConfig("./blocking.json")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	noPush := updateCmd.Bool("noPush", false, "Don't push the blocklists.")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addCategory := addCmd.String("c", "", "Choose a category: ads, adult, etc...")

	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	remCategory := removeCmd.Bool("c", false, "Choose a category: ads, adult, etc...")

	flag.Usage = func() {
		fmt.Printf("Usage: %s [command] [args]\n", os.Args[0])
		fmt.Print(usgMsg)
	}

	flag.Parse()

	// just running the binary updates and merges the blocklists no questions asked
	if len(os.Args) < 2 {
		fmt.Printf("No arguments, default behaviour is to update all blocklists and not git push!\n\n")
		err = UpdateListsAndMergeTags(config, "./dns")
		if err != nil {
			log.Println(err)
		}
		os.Exit(0)
	}

	// TODO: Need to also check for valid category and for empty arguments!
	switch os.Args[1] {
	case "update":
		updateCmd.Parse(os.Args[2:])
		fmt.Println("Updating the blocklists...")
		err = UpdateListsAndMergeTags(config, "./dns")
		if err != nil {
			log.Println(err)
		}
		if !*noPush {
			err := gitAddCommitPushLists()
			if err != nil {
				log.Println(err)
			}
		} else {
			fmt.Println("Finished downloading, no commit and push...")
		}
	case "add":
		addCmd.Parse(os.Args[2:])
		fmt.Printf("Adding new blocklist in category %s\n", *addCategory)
		fmt.Println(addCmd.Args())

		err := AddList(*addCategory, addCmd.Arg(0), config)
		if err != nil {
			log.Println(err)
		}
	case "remove":
		removeCmd.Parse(os.Args[2:])

		if *remCategory {
			fmt.Println("Removing all lists from category ", *remCategory)
			err := RemoveCategory(removeCmd.Arg(0), config)
			if err != nil {
				log.Println(err)
			}
		} else {
			fmt.Println("Removing ", removeCmd.Arg(0))
			err := RemoveList(removeCmd.Arg(0), config)
			if err != nil {
				log.Println(err)
			}
		}
	default:
		flag.Usage()
		os.Exit(1)
	}
}
