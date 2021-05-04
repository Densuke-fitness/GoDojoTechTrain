package gacha

import (
	"github.com/Densuke-fitness/GoDojoTechTrain/dbConnection"
	logger "github.com/sirupsen/logrus"
)

func SelectLotteryRateAndCount() ([]float64, error) {

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

	var name string
	var characterId int
	if err := row.Scan(&name, &characterId); err != nil {
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

	defer tx.Rollback() //nolint

	const sql = (`
		INSERT INTO possession_characters(user_id, character_id, character_seq) VALUES (?, ?, ?);
	`)

	maxSeq, err := SelectMaxSeqNum(userId, characterId)
	if err != nil {
		return err
	}

	maxSeq += 1

	_, err = tx.Exec(sql, userId, characterId, maxSeq)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func SelectMaxSeqNum(userId int, characterId int) (int, error) {
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

	row := tx.QueryRow(sql, userId, characterId)

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
