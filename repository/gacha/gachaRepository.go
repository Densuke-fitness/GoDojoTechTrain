package gacha

import (
	"fmt"

	"github.com/Densuke-fitness/GoDojoTechTrain/dbConnection"
	"github.com/Densuke-fitness/GoDojoTechTrain/model"
	"github.com/Densuke-fitness/GoDojoTechTrain/repository"
)

func SelectLotteryRateList() (LotteryRateList []float64, err error) {

	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer repository.CommitOrRollBack(tx, err)

	const sql = (`
		SELECT lottery_rate 
		FROM characters_lottery_rate 
		WHERE event_id = ? 
		GROUP BY lottery_rate
	`)

	//event_id = 1
	const eventId = 1
	rows, err := db.Query(sql, eventId)

	if err != nil {
		return
	}

	var rate float64

	for rows.Next() {

		if err = rows.Scan(&rate); err != nil {
			return
		}

		LotteryRateList = append(LotteryRateList, rate)
	}

	fmt.Println(LotteryRateList)

	if err = rows.Err(); err != nil {
		return
	}

	return
}

func RandSelectCharacterByRate(rate float64) (character model.Character, err error) {
	//search character_name,character_id  by using userId
	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer repository.CommitOrRollBack(tx, err)

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

	if err = row.Scan(&character.Name, &character.Id); err != nil {
		return
	}

	return
}
