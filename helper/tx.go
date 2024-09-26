package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	defer func() {
		err := recover()
		if err != nil {
			errorRollback := tx.Rollback()
			PanicIfError(errorRollback)
			panic(err)
		} else {
			errCommit := tx.Commit()
			PanicIfError(errCommit)
		}
	}()
}
