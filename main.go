package main

import (
	"cellihub-cli/internal/builder"
	"fmt"
	"log"
	"os"
	"strings"
)

// Add new supported categories here
var supportedCategories = []string{"cloud", "light"}

func usage() {
	fmt.Println("Usage: cellihub-cli <command> <category>")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  devcontainer <category>   Create a devcontainer with the given category")
	fmt.Println()
	fmt.Println("Supported categories:")
	for _, c := range supportedCategories {
		fmt.Printf("  - %s\n", c)
	}
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  cellihub-cli devcontainer cloud")
	fmt.Println("  cellihub-cli devcontainer golang")
}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) < 2 {
		log.Println("No command provided.")
		usage()
		os.Exit(1)
	}

	cmd := strings.ToLower(os.Args[1])

	switch cmd {
	case "devcontainer":
		if len(os.Args) < 3 {
			log.Println("No category provided for devcontainer.")
			usage()
			os.Exit(1)
		}

		category := strings.ToLower(os.Args[2])
		if !contains(supportedCategories, category) {
			log.Printf("Unsupported category: %s\n", category)
			usage()
			os.Exit(1)
		}

		b := builder.NewBuilder()
		// folder name is fixed for now to keep compatibility with templates
		folderName := ".devcontainer"
		log.Printf("Generating devcontainer (category=%s) in %s...", category, folderName)
		b.CreateFolder(folderName)
		b.CheckFolder()

		b.CreateFiles(category)
		log.Println("Devcontainer generation finished.")

	case "help", "-h", "--help":
		usage()

	default:
		log.Printf("Unknown command: %s\n", cmd)
		usage()
		os.Exit(1)
	}
}
