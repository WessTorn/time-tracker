package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time-tracker/config"
	"time-tracker/logger"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	logger.Log.Debug("(ConnectDB)")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBAddress(), config.DBPort(), config.DBUser(), config.DBPass(), config.DBDatabase())

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

	query = `
		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			task_id INTEGER NOT NULL,
			user_id INTEGER REFERENCES users(id) NOT NULL,
            start_time TIMESTAMP NOT NULL,
            end_time TIMESTAMP,
			duration INTERVAL DEFAULT NULL
		);
	`
	_, err = db.Exec(query)
	if err != nil {
		logger.Log.Fatalf("(Exec): %v", err)
	}

	logger.Log.Info("Tables created successfully")
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

func SelectTasts(db *sql.DB, id int) ([]Task, error) {
	var Tasks []Task
	query, err := db.Query(`
		SELECT task_id, start_time, end_time, duration
		FROM tasks
		WHERE 
			user_id = $1 AND end_time IS NOT NULL
		ORDER BY 
			(end_time - start_time)  DESC;
	`, id)

	if err != nil {
		return Tasks, errors.New("FailedtoGetTasks")
	}
	defer query.Close()

	var isFound bool
	for query.Next() {
		var task Task
		err := query.Scan(&task.TaskID, &task.StartTime, &task.EndTime, &task.Duration)
		if err != nil {

			return Tasks, errors.New("FailedtoGetTasks")
		}
		Tasks = append(Tasks, task)
		isFound = true
	}

	if !isFound {
		return Tasks, errors.New("TasksNotFound")
	}

	return Tasks, nil
}

func IsTaskStarted(db *sql.DB, task Task) bool {
	logger.Log.Debug("(IsTaskStarted)")

	var exists bool
	query := `
        SELECT EXISTS (
            SELECT 1 FROM tasks
            WHERE task_id = $1 AND user_id = $2 AND end_time IS NULL
        );
    `
	err := db.QueryRow(query, task.TaskID, task.UserID).Scan(&exists)
	if err != nil {
		return false
	}

	return exists
}

func InsertTask(db *sql.DB, newTask Task) error {
	logger.Log.Debug("(InsertTask)")

	query := `
        INSERT INTO tasks (task_id, user_id, start_time)
        VALUES ($1, $2, $3);
    `
	_, err := db.Exec(query, newTask.TaskID, newTask.UserID, newTask.StartTime)
	if err != nil {
		return err
	}

	fmt.Println("Task added successfully")
	return nil
}

func UpdateTask(db *sql.DB, newTask Task) error {
	logger.Log.Debug("(UpdateTask)")

	query := `
        UPDATE tasks
		SET end_time = $1, duration = DATE_TRUNC('second', $1 - start_time)
        WHERE user_id = $2 AND task_id = $3 AND end_time IS NULL;
    `
	_, err := db.Exec(query, newTask.EndTime, newTask.UserID, newTask.TaskID)
	if err != nil {
		return err
	}

	fmt.Println("Task end successfully")
	return nil
}
