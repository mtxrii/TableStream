package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strings"
)

func main() {
	println("TableStream - See inside your Postgres tables\n+" +
		"-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")

	loggedIn := false
	useEnvCredentials := false
	var DB *sql.DB
	schema := ""

	for {
		// Log in
		var user, password, server, dbname, sslmode string
		for !useEnvCredentials || !loggedIn {
			var useEnvs string
			print("Would you like to log in using credentials in your auth.env? (Y/N): ")
			fmt.Scanln(&useEnvs)
			useEnvs = strings.ToUpper(useEnvs)
			if useEnvs == "Y" {
				useEnvCredentials = true
				err := godotenv.Load("./auth.env")
				if err != nil {
					useEnvCredentials = false
					println("Error loading auth.env file")
				} else {
					user = os.Getenv("DB_USERNAME")
					password = os.Getenv("DB_PASSWORD")
					server = os.Getenv("DB_SERVER")
					dbname = os.Getenv("DB_NAME")
					sslmode = os.Getenv("SSLMODE")
				}
				break
			}
			if useEnvs == "N" {
				useEnvCredentials = false
				break
			}
		}
		for !loggedIn {
			if !useEnvCredentials {
				print("Log in manually...\n" +
					"Server URL: ")
				fmt.Scanln(&server)
				print("Username: ")
				fmt.Scanln(&user)
				print("Password: ")
				fmt.Scanln(&password)
				print("DB Name: ")
				fmt.Scanln(&dbname)
				for {
					print("SSL Mode (T/F): ")
					fmt.Scanln(&sslmode)
					sslmode = strings.ToUpper(sslmode)
					if sslmode == "T" {
						sslmode = "require"
						break
					}
					if sslmode == "F" {
						sslmode = "disable"
						break
					}
				}
			}
			println()

			connectionString := "postgres://" + user + ":" + password + "@" + server + "/" + dbname + "?sslmode=" + sslmode
			var err error
			DB, err = sql.Open("postgres", connectionString)

			if err != nil {
				println(err)
				useEnvCredentials = false
			} else {
				loggedIn = true
			}
		}

		print("Schema: ")
		fmt.Scanln(&schema)

		tables, err := DB.Query("SELECT * FROM information_schema.tables WHERE table_schema = '" + schema + "' AND table_type = 'BASE TABLE'")
		if err != nil {
			log.Fatalln(err)
		}
		var owner, schemaname, tablename, isbase, t5, t6, t7, t8, t9, t10, t11, t12 sql.NullString
		found := false
		println("Tables in schema '" + schema + "':")
		for tables.Next() {
			err := tables.Scan(&owner, &schemaname, &tablename, &isbase, &t5, &t6, &t7, &t8, &t9, &t10, &t11, &t12)
			if err != nil {
				log.Fatal(err)
			}
			println("- " + tablename.String)
			if !found {
				found = true
			}
		}
		println()
		if !found {
			println("- No tables found.")
			continue
		}

		// Main program loop
	out:
		for {
			cmd := ""
			println("What would you like to do now?\n" +
				" peek <table> - see contents of a table\n" +
				" newschema    - exit schema and enter a new one")
			print("-> ")
			fmt.Scanln(&cmd)
			parts := strings.Split(cmd, " ")

			switch parts[0] {
			case "newschema":
				println()
				break out
			case "peek":
				if len(parts) < 2 {
					println("\nUsage: peek <table>\n")
					break
				} else {
					print("Tables in: " + parts[1])
				}
			default:
				println("\nUnknown command.\n")
			}
		}
	}
}
