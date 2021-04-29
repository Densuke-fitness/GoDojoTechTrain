package character

import "github.com/Densuke-fitness/GoDojoTechTrain/dbConnection"

type Character struct {
	CharacterSeq int    `json:"characterSeq"`
	CharacterId  int    `json:"characterId"`
	Name         string `json:"name"`
}

func SelectCharactersById(userId int) ([]Character, error) {
	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	const sql = (`
        SELECT T1.name, T2.character_id, T2.character_seq
		FROM characters_master AS T1
		JOIN possession_characters AS T2
		ON T1.id = T2.character_id
		WHERE user_id = ?
	`)

	rows, err := tx.Query(sql, userId)
	if err != nil {
		tx.Rollback() //nolint
		return nil, err
	}

	var Characters []Character

	for rows.Next() {
		var c Character
		if err := rows.Scan(&c.Name, &c.CharacterId, &c.CharacterSeq); err != nil {
			tx.Rollback() //nolint
			return nil, err
		}
		Characters = append(Characters, c)
	}

	if err := rows.Err(); err != nil {
		tx.Rollback() //nolint
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return Characters, nil

}
