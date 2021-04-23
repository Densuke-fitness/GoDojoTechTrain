package gacha

import (
	"github.com/Densuke-fitness/GoDojoTechTrain/dbConnection"
)

func Select(userLotteryNum float64) (string, int, error) {

	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	const sql = (`
	SELECT name, character_id, lottery_rate FROM characters_master JOIN characters_lottery_rate ON
    characters_master.id=characters_lottery_rate.character_id
	`)

	rows, err := db.Query(sql)
	if err != nil {
		return "", -1, err
	}

	var characterName string
	var characterId int
	var CharacterRate float64
	var rate float64
	for rows.Next() {

		if err := rows.Scan(&characterName, &characterId, &rate); err != nil {
			return "", -1, err
		}

		CharacterRate += rate
		if userLotteryNum <= CharacterRate {
			return characterName, characterId, nil
		}
	}

	if err := rows.Err(); err != nil {
		return "", -1, err
	}

	return characterName, characterId, nil
}
