package repository

import "database/sql"

func CommitOrRollBack(tx *sql.Tx, err error) {
	if err != nil {
		tx.Rollback() //nolint
		return
	}

	err = tx.Commit()
	if err != nil {
		return
	}
}
