package main

import (
	"fmt"
	"os"

	"cellihub-cli/internal/cloud/aws"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("cellihub cloud aws devcontainer")
		os.Exit(1)
	}

	cmd1 := os.Args[1]
	cmd2 := os.Args[2]
	cmd3 := os.Args[3]

	switch cmd1 {
	case "cloud":
		if cmd2 == "aws" && cmd3 == "devcontainer" {
			aws.GenerateDevContainer()
		} else {
			fmt.Println("command not found. Try: cellihub cloud aws devcontainer")
		}
	default:
		fmt.Println("command not found. Try: cellihub cloud aws devcontainer")
	}
}
