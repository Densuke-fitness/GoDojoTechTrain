package users

import (
	"database/sql"

	"github.com/Densuke-fitness/GoDojoTechTrain/dbConnection"
	"github.com/Densuke-fitness/GoDojoTechTrain/model"
	"github.com/Densuke-fitness/GoDojoTechTrain/repository"
)

func Insert(user model.User) (ret *model.User, err error) {
	//Insert a name into a user table using sql
	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer repository.CommitOrRollBack(tx, err)

	const sql = "INSERT INTO users(name) VALUES (?)"
	//Save the name data (id is automatically generated)
	r, err := tx.Exec(sql, user.Name)
	if err != nil {
		return
	}

	id, err := r.LastInsertId()
	if err != nil {
		return
	}

	ret = user.Clone(int(id))

	return
}

func SelectNameById(user model.User) (ret *model.User, err error) {
	//search name by using id
	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	const sql = "SELECT name FROM users WHERE id = ?"
	row := db.QueryRow(sql, user.Id)

	if err = row.Scan(&user.Name); err != nil {
		return
	}

	ret = user.Clone(user.Id)

	return
}

func UpdateNameById(user model.User) (result sql.Result, err error) {
	//Insert a name into a user table using sql
	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer repository.CommitOrRollBack(tx, err)

	const sql = "UPDATE users SET name = ? WHERE id = ?"
	//Since the number of updates was originally returned, the result was adopted as the return value.
	result, err = tx.Exec(sql, user.Name, user.Id)
	if err != nil {
		return
	}

	return
}
