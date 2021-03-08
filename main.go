package main

import (
	"flag"
	"fmt"
	"os"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type configStruct struct {
	AdminUser       string
	GrafanaDatabase string
}

func main() {
	config := parseArgs()

	db, err := sql.Open("sqlite3", config.GrafanaDatabase)

	if err != nil {
		fmt.Print(err)
		fmt.Println("")
		os.Exit(-1)
	}

	db.SetMaxOpenConns(1)
	defer db.Close()

	users, err := db.Query("SELECT login FROM user WHERE login=?;", config.AdminUser)
	if err != nil {
		fmt.Print(err)
		fmt.Println("")

		os.Exit(-1)
	}

	var foundLogin string
	if !users.Next() {
		fmt.Printf("User %s not found. Exiting...\n", config.AdminUser)
		os.Exit(0)
	}

	err = users.Scan(&foundLogin)
	if err != nil {
		fmt.Print(err)
		fmt.Println("")

		os.Exit(-1)
	}

	fmt.Printf("User %s found. Promoting to admin...\n", foundLogin)
	users.Close()

	_, err = db.Exec("UPDATE user SET is_admin=1 WHERE login=?;", config.AdminUser)
	if err != nil {
		fmt.Print(err)
		fmt.Println("")

		os.Exit(-1)
	}

	newAdminUsers, err := db.Query("SELECT login FROM user WHERE is_admin=1;")
	if err != nil {
		fmt.Print(err)
		fmt.Println("")

		os.Exit(-1)
	}

	defer newAdminUsers.Close()

	fmt.Println("Your new admin users are:")
	for newAdminUsers.Next() {
		var user string
		newAdminUsers.Scan(&user)

		fmt.Printf("- %s\n", user)
	}

	fmt.Println("")

}

func parseArgs() configStruct {

	var conf = configStruct{}
	var help = flag.Bool("help", false, "Prints this message")
	flag.StringVar(&conf.AdminUser, "user", "admin", "The `login` of the user you wish to promote")
	flag.StringVar(&conf.GrafanaDatabase, "database", "/var/lib/grafana/grafana.db", "The grafana `database` to update")

	flag.Parse()

	if *help == true {
		displayHelp()
		os.Exit(0)
	}

	return conf
}

func displayHelp() {
	fmt.Println("Grafana Adminizer")
	fmt.Println("For when you have lost your server admin account")
	fmt.Println("")

	flag.PrintDefaults()
	fmt.Println("")
}
