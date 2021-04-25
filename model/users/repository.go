package users

import (
	"database/sql"

	"github.com/Densuke-fitness/GoDojoTechTrain/dbConnection"
)

func Insert(name string) (int, error) {
	//Insert a name into a user table using sql
	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	const sql = "INSERT INTO users(name) VALUES (?)"
	//Save the name data (id is automatically generated)
	r, err := db.Exec(sql, name)
	if err != nil {
		return -1, err
	}

	id, err := r.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil

}

func SelectNameById(id int) (string, error) {
	//search name by using id
	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	var name string

	const sql = "SELECT name FROM users WHERE id = ?"
	row := db.QueryRow(sql, id)
	if err := row.Scan(&name); err != nil {
		return "", err
	}
	return name, nil
}

func UpdateNameById(name string, id int) (sql.Result, error) {
	//Insert a name into a user table using sql
	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	const sql = "UPDATE users SET name = ? WHERE id = ?"
	//Since the number of updates was originally returned, the result was adopted as the return value.
	result, err := db.Exec(sql, name, id)
	return result, err
}
