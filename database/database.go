package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time-tracker/logger"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	logger.Log.Debug("(ConnectDB)")

	psqlInfo := "host=localhost port=5432 user=postgres password=root dbname=tz_iul sslmode=disable"
	logger.Log.Debugf("(psqlInfo) %s", psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Log.Fatalf("(Open): %v", err)
	}

	err = db.Ping()
	if err != nil {
		logger.Log.Fatalf("(Ping): %v", err)
	}

	logger.Log.Info("Database connected")

	return db
}

func CreateSchema(db *sql.DB) {
	logger.Log.Debug("(CreateSchema)")

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
		logger.Log.Fatalf("(Exec): %v", err)
	}

	logger.Log.Info("Table created successfully")
}

func SelectUsers(db *sql.DB, filter User, limit int, page int) ([]User, error) {
	logger.Log.Debug("(SelectUsers)")
	var Users []User
	query, err := db.Query(`
		SELECT id, surname, name, patronymic, address, passport_serie, passport_number
		FROM users
		WHERE (passport_serie = $1 OR $1 = '') AND
			(passport_number = $2 OR $2 = '') AND
			(surname = $3 OR $3 = '') AND
			(name = $4 OR $4 = '') AND
			(patronymic = $5 OR $5 = '') AND
			(address = $6 OR $6 = '')
		ORDER BY id ASC
		LIMIT $7 OFFSET $8;
	`,
		filter.PassportSerie,
		filter.PassportNumber,
		filter.LastName,
		filter.FirstName,
		filter.Patronymic,
		filter.Address,
		limit,
		(page-1)*limit)

	if err != nil {
		return Users, errors.New("FailedtoGetUsers")
	}
	defer query.Close()

	var isFound bool
	for query.Next() {
		var user User
		err := query.Scan(&user.ID, &user.LastName, &user.FirstName, &user.Patronymic, &user.Address, &user.PassportSerie, &user.PassportNumber)
		if err != nil {
			return Users, errors.New("FailedtoGetUsers")
		}
		Users = append(Users, user)
		isFound = true
	}

	if !isFound {
		return Users, errors.New("UsersNotFound")
	}

	return Users, nil
}

func DeletUserFromID(db *sql.DB, id string) error {
	logger.Log.Debug("(DeletUserFromID)")

	res, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return errors.New("FailedToDeleteUser")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.New("FailedToDeleteUser")
	}

	if rowsAffected == 0 {
		return errors.New("UserNotFound")
	}

	return nil
}

func UpdateUser(db *sql.DB, id string, set string, paramsToUpdate []interface{}) error {
	logger.Log.Debug("(UpdateUser)")

	sqlStatement := fmt.Sprintf("UPDATE users SET %s WHERE id=%s;", set, id)
	result, err := db.Exec(sqlStatement, paramsToUpdate...)
	if err != nil {
		return errors.New("FailedToUpdateUser")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("FailedToUpdateUser")
	}

	if rowsAffected == 0 {
		return errors.New("UserNotFound")
	}

	return nil
}

func IsUserExists(db *sql.DB, user User) bool {
	logger.Log.Debug("(IsUserExists)")

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

func InsertUser(db *sql.DB, newUser User) error {
	logger.Log.Debug("(InsertUser)")

	query := `
        INSERT INTO users (surname, name, patronymic, address, passport_serie, passport_number)
        VALUES ($1, $2, $3, $4, $5, $6);
    `

	_, err := db.Exec(query, newUser.LastName, newUser.FirstName, newUser.Patronymic, newUser.Address, newUser.PassportSerie, newUser.PassportNumber)
	if err != nil {
		return err
	}

	fmt.Println("User added successfully")
	return nil
}
