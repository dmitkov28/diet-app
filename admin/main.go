// cmd/admin/main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dmitkov28/dietapp/data"
)

func main() {

	db, err := data.NewDB()

	if err != nil {
		fmt.Println("couldn't connect to db")
		os.Exit(1)
	}

	usersRepo := data.NewUsersRepository(db)

	// Command line flags
	createUserCmd := flag.NewFlagSet("create-user", flag.ExitOnError)
	// listUsersCmd := flag.NewFlagSet("list-users", flag.ExitOnError)
	// deleteUserCmd := flag.NewFlagSet("delete-user", flag.ExitOnError)

	// create-user flags
	email := createUserCmd.String("email", "", "User's email")
	password := createUserCmd.String("password", "", "User's password")

	// delete-user flags
	// userEmail := deleteUserCmd.String("email", "", "Email of user to delete")

	if len(os.Args) < 2 {
		fmt.Println("expected 'create-user', 'list-users', or 'delete-user' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create-user":
		createUserCmd.Parse(os.Args[2:])
		if *email == "" || *password == "" {
			createUserCmd.PrintDefaults()
			os.Exit(1)
		}

		user, err := usersRepo.CreateUser(*email, *password)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Created user: %s, id: %d \n", user.Email, user.ID)

	// case "list-users":

	case "delete-user":
		fmt.Println("delete")

	default:
		fmt.Println("expected 'create-user', 'list-users', or 'delete-user' subcommands")
		os.Exit(1)
	}
}
