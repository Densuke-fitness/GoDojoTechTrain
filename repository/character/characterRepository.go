package character

import (
	"github.com/Densuke-fitness/GoDojoTechTrain/dbConnection"
	"github.com/Densuke-fitness/GoDojoTechTrain/model"
	"github.com/Densuke-fitness/GoDojoTechTrain/repository"
	logger "github.com/sirupsen/logrus"
)

func Insert(user model.User, character model.Character) (err error) {

	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer repository.CommitOrRollBack(tx, err)

	const sql = (`
		INSERT INTO
		possession_characters(user_id, character_id, character_seq)
		VALUES (?, ?, ?);
    `)

	_, err = tx.Exec(sql, user.Id, character.Id, character.CharacterSeq)
	if err != nil {
		return
	}
	return
}

func SelectMaxSeqNum(user model.User, character model.Character) (maxSeq int, err error) {
	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		logger.Errorf("Error b.Begin: %s", err)
		return
	}

	defer repository.CommitOrRollBack(tx, err)

	const sql = (`
		SELECT COALESCE(MAX(character_seq), 0)
		FROM possession_characters
		WHERE user_id = ?
		AND character_id = ?
	`)

	row := tx.QueryRow(sql, user.Id, character.Id)

	if err = row.Scan(&maxSeq); err != nil {
		logger.Errorf("Error row.Scan: %s", err)
		return
	}

	return
}

func SelectCharactersByUserId(user model.User) (characters []model.Character, err error) {
	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	defer repository.CommitOrRollBack(tx, err)

	const sql = (`
        SELECT T1.name, T2.character_id, T2.character_seq
		FROM characters_master AS T1
		JOIN possession_characters AS T2
		ON T1.id = T2.character_id
		WHERE user_id = ?
	`)

	rows, err := tx.Query(sql, user.Id)
	if err != nil {
		return
	}

	for rows.Next() {
		var c model.Character
		if err = rows.Scan(&c.Name, &c.Id, &c.CharacterSeq); err != nil {
			return
		}
		characters = append(characters, c)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}
