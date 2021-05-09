package gacha

import (
	"github.com/Densuke-fitness/GoDojoTechTrain/dbConnection"
	"github.com/Densuke-fitness/GoDojoTechTrain/model"
)

func SelectLotteryRateList() ([]float64, error) {

	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback() //nolint

	const sql = (`
		SELECT lottery_rate 
		FROM characters_lottery_rate 
		WHERE event_id = ? 
		GROUP BY lottery_rate
	`)

	//event_i = 1
	const eventId = 1
	rows, err := db.Query(sql, eventId)

	if err != nil {
		return nil, err
	}

	var LotteryRateList []float64
	var rate float64

	for rows.Next() {

		if err := rows.Scan(&rate); err != nil {
			return nil, err
		}

		LotteryRateList = append(LotteryRateList, rate)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return LotteryRateList, nil

}

func RandSelectCharacterByRate(rate float64) (*model.Character, error) {
	//search character_name,character_id  by using userId
	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback() //nolint

	//Execute select to get the sequence
	const sql = (`
		SELECT T1.name, T2.character_id 
		FROM characters_master AS T1
		JOIN characters_lottery_rate AS T2
		ON T1.id = T2.character_id 
		WHERE CAST(T2.lottery_rate AS CHAR) = ?
		ORDER BY RAND() LIMIT 1
	`)
	row := tx.QueryRow(sql, rate)

	character := model.Character{}

	if err := row.Scan(&character.Name, &character.Id); err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &character, nil
}
