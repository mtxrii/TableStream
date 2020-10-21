package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"strings"
)

func main() {
	println("TableStream - See inside your Postgres tables")
	println("---------------------------------------------")

	loggedIn := false
	var DB *sql.DB

	for {
		// Log in
		for !loggedIn {
			var user, password, server, dbname string
			print("Log in...\n" +
				"Server URL: ")
			fmt.Scanf("%s\n", &server)
			print("Username: ")
			fmt.Scanf("%s\n", &user)
			print("Password: ")
			fmt.Scanf("%s\n", &password)
			print("DB Name: ")
			fmt.Scanf("%s\n", &dbname)
			sslmode := "none"
			for {
				print("SSL Mode (T/F): ")
				fmt.Scanf("%s\n", &sslmode)
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


			connectionString := "postgres://" + user + ":" + password + "@" + server + "/" + dbname + "?sslmode=" + sslmode
			var err error
			DB, err = sql.Open("postgres", connectionString)

			if err != nil {
				log.Fatal(err)
			} else {
				loggedIn = true
			}
		}

		tables, err := DB.Query("SELECT * FROM information_schema.tables WHERE table_schema = 'Lab1'")
		if err != nil {
			log.Fatalln(err)
		}
		for tables.Next() {
			var t1, t2 string
			tables.Scan(&t1, &t2)
			println(t1 + ", " + t2)
		}




	}

}
