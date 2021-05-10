package character

import (
	"github.com/Densuke-fitness/GoDojoTechTrain/dbConnection"
	"github.com/Densuke-fitness/GoDojoTechTrain/model"
	logger "github.com/sirupsen/logrus"
)

func Insert(user model.User, character model.Character) (err error) {

	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback() //nolint
		} else {
			err = tx.Commit()
			if err != nil {
				return
			} else {
				return
			}
		}
	}()

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

func SelectMaxSeqNum(user model.User, character model.Character) (int, error) {
	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		logger.Errorf("Error b.Begin: %s", err)
		return -1, err
	}

	defer tx.Rollback() //nolint

	const sql = (`
		SELECT COALESCE(MAX(character_seq), 0)
		FROM possession_characters
		WHERE user_id = ?
		AND character_id = ?
	`)

	row := tx.QueryRow(sql, user.Id, character.Id)

	var maxSeq int
	if err := row.Scan(&maxSeq); err != nil {
		logger.Errorf("Error row.Scan: %s", err)
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return maxSeq, nil
}

func SelectCharactersByUserId(user model.User) ([]model.Character, error) {
	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback() //nolint

	const sql = (`
        SELECT T1.name, T2.character_id, T2.character_seq
		FROM characters_master AS T1
		JOIN possession_characters AS T2
		ON T1.id = T2.character_id
		WHERE user_id = ?
	`)

	rows, err := tx.Query(sql, user.Id)
	if err != nil {
		return nil, err
	}

	var characters []model.Character

	for rows.Next() {
		var c model.Character
		if err := rows.Scan(&c.Name, &c.Id, &c.CharacterSeq); err != nil {
			return nil, err
		}
		characters = append(characters, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return characters, nil

}
