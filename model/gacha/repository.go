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
		//Initial value is 0, but since the value of rows can be 0 (normal), we set it to -1
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

// func Insert(userId int, characterId int) (sql.Result, error) {
// 	//search character_name,character_id  by using userId
// 	dbConn := dbConnection.GetInstance()

// 	db := dbConn.GetConnection()

// 	//Execute select to get the sequence
// 	const sql1 =  "INSERT INTO possession_character(user_id, character_id, ) VALUES (?)"

// 	const sql2 = "INSERT INTO possession_character(user_id, character_id, ) VALUES (?)"
// 	result, err := db.Exec(sql, e, id)
// 	return result, err
// }
