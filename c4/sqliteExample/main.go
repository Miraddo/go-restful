package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)
// Book is a placeholder for book
type Book struct {
	id int
	name string
	author string
}

func main()  {
	db, err := sql.Open("sqlite3", "./c4/sqliteExample/books.db")
	if err != nil{
		log.Println(err)
	}

	// Create table
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL)")
	if err != nil {
		log.Println("Error in creating table")
	} else {
		log.Println("Successfully created table books!")
	}
	_, err = statement.Exec()
	if err != nil {
		log.Println(err)
	}
	dbOperations(db)
}

func dbOperations(db *sql.DB) {
	// Create
	statement, _ := db.Prepare("INSERT INTO books (name, author, isbn) VALUES (?, ?, ?)")
	_, err := statement.Exec("A Tale of Two Cities", "Charles Dickens",
		140430547)
	if err != nil {
		log.Println(err)
	}
	log.Println("Inserted the book into database!")
	// Read
	rows, _ := db.Query("SELECT id, name, author FROM books")
	var tempBook Book
	for rows.Next() {
		err = rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		if err != nil {
			log.Println(err)
		}
		log.Printf("ID:%d, Book:%s, Author:%s\n", tempBook.id,
			tempBook.name, tempBook.author)
	}
	// Update
	statement, _ = db.Prepare("update books set name=? where id=?")
	_, err = statement.Exec("The Tale of Two Cities", 1)
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully updated the book in database!")
	//Delete
	statement, _ = db.Prepare("delete from books where id=?")
	_, err = statement.Exec(1)
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully deleted the book in database!")
}