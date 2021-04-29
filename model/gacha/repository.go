package gacha

import (
	"github.com/Densuke-fitness/GoDojoTechTrain/dbConnection"
	logger "github.com/sirupsen/logrus"
)

func SelectLotteryRateAndCount() (map[float64]int, error) {

	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	const sql = "SELECT lottery_rate, COUNT(lottery_rate) FROM characters_lottery_rate GROUP BY lottery_rate;"

	rows, err := db.Query(sql)

	if err != nil {
		tx.Rollback() //nolint
		return nil, err
	}

	var rate float64
	var count int

	LotteryRateMap := map[float64]int{}
	for rows.Next() {

		if err := rows.Scan(&rate, &count); err != nil {
			tx.Rollback()
			return nil, err
		}

		LotteryRateMap[rate] = count
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return LotteryRateMap, nil

}

func RandSelectCharacterByRate(rate float64) (string, int, error) {
	//search character_name,character_id  by using userId
	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		// In model/user/repository.go,
		// the return value in case of an error is set to -1 instead of 0,
		// so it is described in the same way.
		return "", -1, err
	}

	//Execute select to get the sequence
	const sql = (`
		SELECT T1.name, T2.character_id 
		FROM characters_master AS T1
		JOIN characters_lottery_rate AS T2
		ON T1.id=T2.character_id 
		WHERE CAST(T2.lottery_rate AS CHAR) = ?
		ORDER BY RAND() LIMIT 1
	`)
	row := tx.QueryRow(sql, rate)

	var name string
	var characterId int
	if err := row.Scan(&name, &characterId); err != nil {
		tx.Rollback() //nolint
		return "", -1, err
	}

	err = tx.Commit()
	if err != nil {
		return "", -1, err
	}

	return name, characterId, nil
}

func Insert(userId int, characterId int) error {

	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	const sql = (`
		INSERT INTO possession_characters(user_id, character_id, character_seq) VALUES (?, ?, ?);
	`)

	maxSeq := SelectMaxSeqNum(userId, characterId)
	maxSeq += 1

	_, err = tx.Exec(sql, userId, characterId, maxSeq)
	if err != nil {
		tx.Rollback() //nolint
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func SelectMaxSeqNum(userId int, characterId int) int {
	dbConn := dbConnection.GetInstance()

	db := dbConn.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		logger.Errorf("Error b.Begin: %s", err)
		return -1
	}

	const sql = (`
		SELECT  MAX(character_seq)
		FROM possession_characters
		WHERE user_id = ?
		AND character_id = ?
	`)

	row := tx.QueryRow(sql, userId, characterId)

	var maxSeq int
	if err := row.Scan(&maxSeq); err != nil {
		logger.Errorf("Error row.Scan: %s", err)
		tx.Rollback() //nolint
		return 0
	}

	err = tx.Commit()
	if err != nil {
		return 0
	}

	return maxSeq

}
