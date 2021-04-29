package users

import (
	"database/sql"

	"github.com/Densuke-fitness/GoDojoTechTrain/dbConnection"
)

func Insert(name string) (int, error) {
	//Insert a name into a user table using sql
	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return -1, err
	}

	const sql = "INSERT INTO users(name) VALUES (?)"
	//Save the name data (id is automatically generated)
	r, err := tx.Exec(sql, name)
	if err != nil {
		tx.Rollback() //nolint
		return -1, err
	}

	id, err := r.LastInsertId()
	if err != nil {
		tx.Rollback() //nolint
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}

func SelectNameById(id int) (string, error) {
	//search name by using id
	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return "", err
	}

	const sql = "SELECT name FROM users WHERE id = ?"
	row := tx.QueryRow(sql, id)

	var name string
	if err := row.Scan(&name); err != nil {
		tx.Rollback() //nolint
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return name, nil
}

func UpdateNameById(name string, id int) (sql.Result, error) {
	//Insert a name into a user table using sql
	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	const sql = "UPDATE users SET name = ? WHERE id = ?"
	//Since the number of updates was originally returned, the result was adopted as the return value.
	result, err := tx.Exec(sql, name, id)
	if err != nil {
		tx.Rollback() //nolint
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}
