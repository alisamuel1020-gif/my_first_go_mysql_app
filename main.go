package main

import (
	"log"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Add a new user
	insertQuery := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err = db.Exec(insertQuery, "Little Timmy", "timmy@example.com")
	if err != nil {
		log.Fatal("Oops, couldn't add the user: ", err)
	}
	log.Println("🎉 We added a new user!")

	// Read all users
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal("Oops, couldn't read the users: ", err)
	}
	defer rows.Close()

	log.Println("📋 Here are all the users in the database:")
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			log.Fatal("Couldn't read a line: ", err)
		}
		log.Printf("- ID: %d | Name: %s | Email: %s\n", u.ID, u.Name, u.Email)
	}
}
