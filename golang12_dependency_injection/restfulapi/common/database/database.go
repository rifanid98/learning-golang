package database

import (
	"database/sql"
	error2 "golang12_dependency_injection/restfulapi/common/error"
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
