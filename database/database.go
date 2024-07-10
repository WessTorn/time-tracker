package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	psqlInfo := "host=localhost port=5432 user=postgres password=root dbname=tz_iul sslmode=disable"

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func CreateSchema(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			surname VARCHAR(255) NOT NULL,
			name VARCHAR(255) NOT NULL,
			patronymic VARCHAR(255) NOT NULL,
			address VARCHAR(255) NOT NULL,
			passport_serie VARCHAR(255) NOT NULL,
			passport_number VARCHAR(255) NOT NULL
		);
	`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table created successfully")
}

func IsUserExists(db *sql.DB, user User) bool {
	var exists bool
	query := `
        SELECT EXISTS (
            SELECT 1 FROM users
            WHERE passport_serie = $1 AND passport_number = $2
        );
    `
	err := db.QueryRow(query, user.PassportSerie, user.PassportNumber).Scan(&exists)
	if err != nil {
		return false
	}

	return exists
}

func AddNewUser(db *sql.DB, user User) error {
	query := `
        INSERT INTO users (surname, name, patronymic, address, passport_serie, passport_number)
        VALUES ($1, $2, $3, $4, $5, $6);
    `

	_, err := db.Exec(query, user.LastName, user.FirstName, user.Patronymic, user.Address, user.PassportSerie, user.PassportNumber)
	if err != nil {
		return err
	}

	fmt.Println("User added successfully")
	return nil
}
