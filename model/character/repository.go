package character

import "github.com/Densuke-fitness/GoDojoTechTrain/dbConnection"

type Character struct {
	/* About: CharacterSeq
	In the api specification, the response name is described as userCharacterID,
	but the id is a sequence number(CharacterSeq) and is difficult to understand,
	so the field name is left as a sequence.
	*/
	CharacterSeq string `json:"userCharacterID"`
	CharacterId  string `json:"characterId"`
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

	var characters []Character

	for rows.Next() {
		var c Character
		if err := rows.Scan(&c.Name, &c.CharacterId, &c.CharacterSeq); err != nil {
			tx.Rollback() //nolint
			return nil, err
		}
		characters = append(characters, c)
	}

	if err := rows.Err(); err != nil {
		tx.Rollback() //nolint
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return characters, nil

}
