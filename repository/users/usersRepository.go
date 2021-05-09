package users

import (
	"database/sql"

	"github.com/Densuke-fitness/GoDojoTechTrain/dbConnection"
	"github.com/Densuke-fitness/GoDojoTechTrain/model"
)

func Insert(user model.User) (*model.User, error) {
	//Insert a name into a user table using sql
	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback() //nolint

	const sql = "INSERT INTO users(name) VALUES (?)"
	//Save the name data (id is automatically generated)
	r, err := tx.Exec(sql, user.Name)
	if err != nil {
		return nil, err
	}

	id, err := r.LastInsertId()
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	clonedUser := user.Clone(int(id))

	return clonedUser, nil
}

func SelectNameById(user model.User) (*model.User, error) {
	//search name by using id
	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback() //nolint

	const sql = "SELECT name FROM users WHERE id = ?"
	row := tx.QueryRow(sql, user.Id)

	if err := row.Scan(&user.Name); err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	clonedUser := user.Clone(user.Id)

	return clonedUser, nil
}

func UpdateNameById(user model.User) (sql.Result, error) {
	//Insert a name into a user table using sql
	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback() //nolint

	const sql = "UPDATE users SET name = ? WHERE id = ?"
	//Since the number of updates was originally returned, the result was adopted as the return value.
	result, err := tx.Exec(sql, user.Name, user.Id)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}
