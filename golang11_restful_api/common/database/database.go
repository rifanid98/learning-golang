package database

import (
	"database/sql"
	error2 "golang11_restful_api/common/error"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		txErr := tx.Rollback()
		error2.PanicIfError(txErr)
	} else {
		txErr := tx.Commit()
		error2.PanicIfError(txErr)
	}
}
