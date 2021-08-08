package dbutils

import (
	"database/sql"
	"log"
)

func Initialize(dbDriver *sql.DB) {
	// creating train table
	statement, driverError := dbDriver.Prepare(train)
	if driverError != nil {
		log.Println(driverError)
	}
	// Create train table
	_, statementError := statement.Exec()
	if statementError != nil {
		log.Println("Table already exists!")
	}
	// creating station table
	statement, _ = dbDriver.Prepare(station)
	_, err := statement.Exec()
	if err != nil {
		return
	}

	// creating schedule table
	statement, _ = dbDriver.Prepare(schedule)
	_, err = statement.Exec()
	if err != nil {
		return
	}
	log.Println("All tables created/initialized successfully!")
}