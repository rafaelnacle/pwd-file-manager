package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rafaelnacle/password-manager/src/passwords"
)

func main() {
	// Command-line flags
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	editCmd := flag.NewFlagSet("edit", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	// Flag variables
	addWebsite := addCmd.String("w", "", "Website")
	addUsername := addCmd.String("u", "", "Username")
	addPassword := addCmd.String("p", "", "Password")

	editWebsite := editCmd.String("w", "", "Website")
	editUsername := editCmd.String("u", "", "New username")
	editPassword := editCmd.String("p", "", "New password")

	deleteWebsite := deleteCmd.String("w", "", "Website")

	// Check subcommands
	if len(os.Args) < 2 {
		fmt.Println("Usage: pwd-manager <command>")
		os.Exit(1)
	}

	// Parse commands
	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if *addWebsite == "" || *addUsername == "" || *addPassword == "" {
			fmt.Println("Usage: pwd-manager add -w <website> -u <username> -p <password>")
			os.Exit(1)
		}
		err := passwords.AddPassword(*addWebsite, *addUsername, *addPassword)
		if err != nil {
			fmt.Println("Error adding password:", err)
			os.Exit(1)
		}
		fmt.Println("Password added succesfully")
	case "edit":
		editCmd.Parse(os.Args[2:])
		if *editWebsite == "" || (*editUsername == "" && *editPassword == "") {
			fmt.Println("Usage: password-manager edit -w <website> [-u <new_username>] [-p <new_password>]")
			os.Exit(1)
		}
		err := passwords.EditPassword(*editWebsite, *editUsername, *editPassword)
		if err != nil {
			fmt.Println("Error editing password:", err)
			os.Exit(1)
		}
		fmt.Println("Password edited successfully.")
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if *deleteWebsite == "" {
			fmt.Println("Usage: password-manager delete -w <website>")
			os.Exit(1)
		}
		err := passwords.DeletePassword(*deleteWebsite)
		if err != nil {
			fmt.Println("Error deleting password:", err)
			os.Exit(1)
		}
		fmt.Println("Password deleted successfully.")
	case "list":
		passwordsList, err := passwords.ListPasswords()
		if err != nil {
			fmt.Println("Error listing passwords:", err)
			os.Exit(1)
		}
		fmt.Println("Passwords:")
		for _, pw := range passwordsList {
			fmt.Printf("Website: %s, Username: %s, Password: %s\n", pw.Website, pw.Username, pw.Password)
		}
	default:
		fmt.Println("Unknown command:", os.Args[1])
		fmt.Println("Usage: pwd-manager <command>")
		os.Exit(1)
	}
}
