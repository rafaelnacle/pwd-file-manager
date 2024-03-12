package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Command-line flags
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	editCmd := flag.NewFlagSet("edit", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	// Flag variables
	addWebsite := addCmd.String("w", "", "Website")
	addUsername := addCmd.String("u", "", "Username")
	addPassword := addCmd.String("p", "", "Password")

	editWebsite := addCmd.String("w", "", "Website")
	editUsername := addCmd.String("u", "", "New username")
	editPassword := addCmd.String("p", "", "New password")

	deleteWebsite := deleteCmd.String("w", "", "Website")

	// Check subcommands
	if len(os.Args) < 2 {
		fmt.Println("Usage: pwd-manager <command>")
		os.Exit(1)
	}
}
