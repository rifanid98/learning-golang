package common

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		txErr := tx.Rollback()
		PanicIfError(txErr)
	}
	txErr := tx.Commit()
	PanicIfError(txErr)
}