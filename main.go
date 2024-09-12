package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const usgMsg = "Commands:\n" +
	"\t Running without arguments Updates and Merges the blocklists.\n" +
	"\t update Updates and Merges the blocklists.\n" +
	"\t add -c <category> <blocklists> Adds the following blocklists to the config.\n" +
	"\t remove -c <category> <blocklists> Removes the blocklists.\n" +
	"The blocklists must be given with their full Path or URL!\n"

var WEBHOOK = ""

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	WEBHOOK = os.Getenv("WEBHOOK")

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
	remCategory := removeCmd.String("c", "", "Choose a category: ads, adult, etc...")

	flag.Usage = func() {
		fmt.Printf("Usage: %s [command] [args]\n", os.Args[0])
		fmt.Print(usgMsg)
	}

	flag.Parse()

	// just running the binary updates and merges the blocklists no questions asked
	if len(os.Args) < 2 {
		fmt.Printf("No arguments, default behaviour is to update all blocklists!\n\n")
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
		fmt.Println("NOT IMPLEMENTED YET!")
		// add function
	case "remove":
		removeCmd.Parse(os.Args[2:])
		fmt.Printf("Removing blocklist from category %s\n", *remCategory)
		fmt.Println(removeCmd.Args())
		fmt.Println("NOT IMPLEMENTED YET!")
		// remove function
	default:
		flag.Usage()
		os.Exit(1)
	}
}
