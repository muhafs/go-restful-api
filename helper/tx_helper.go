package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	if err := recover(); err != nil {
		errRb := tx.Rollback()
		PanicIfError(errRb)

		panic(err)
	} else {
		errCm := tx.Commit()
		PanicIfError(errCm)
	}
}
