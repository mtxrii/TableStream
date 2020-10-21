package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"strings"
)

func main() {
	println("TableStream - See inside your Postgres tables\n+" +
		"-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")

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
			println()

			connectionString := "postgres://" + user + ":" + password + "@" + server + "/" + dbname + "?sslmode=" + sslmode
			var err error
			DB, err = sql.Open("postgres", connectionString)

			if err != nil {
				println(err)
			} else {
				loggedIn = true
			}
		}

		schema := ""
		print("Schema: ")
		fmt.Scanf("%s\n", &schema)

		tables, err := DB.Query("SELECT * FROM information_schema.tables WHERE table_schema = '" + schema + "' AND table_type = 'BASE TABLE'")
		if err != nil {
			log.Fatalln(err)
		}
		for tables.Next() {
			var t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12 sql.NullString
			err := tables.Scan(&t1, &t2, &t3, &t4, &t5, &t6, &t7, &t8, &t9, &t10, &t11, &t12)
			println(t1.String + ", " + t2.String + ", " + t3.String + ", " + t4.String + ", " + t5.String + ", " + t6.String + ", " + t7.String + ", " + t8.String + ", " + t9.String + ", " + t10.String + ", " + t11.String + ", " + t12.String)
			if err != nil {
				log.Fatal(err)
			}
		}
		break

	}

}
