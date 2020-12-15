package main

import (
	"database/sql"
	"errors"
)

func insertDialog(tx *sql.Tx, userA, userB uint64) (dialogID uint64) {
	row := tx.QueryRow(
		`
WITH
    us AS (
        SELECT
            dialog_id,
            group_concat(user_id) AS user_ids
        FROM dialog_users
        GROUP BY dialog_id
        HAVING count(*) = 2
    )
SELECT
    dialog_id
FROM us
WHERE user_ids = CONCAT(?, ',', ?)
   OR user_ids = CONCAT(?, ',', ?)`,
		userA, userB, userB, userA,
	)
	if row.Err() != nil {
		panic(row.Err())
	}

	err := row.Scan(&dialogID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			res, err := tx.Exec(`INSERT INTO dialogs(owner_id) VALUES (?)`, userA)
			if err != nil {
				panic(err)
			}
			lastInsertedDialog, err := res.LastInsertId()
			if err != nil {
				panic(err)
			}
			dialogID = uint64(lastInsertedDialog)

			for _, userID := range []uint64{userA, userB} {
				_, err := tx.Exec(`INSERT INTO dialog_users (dialog_id, user_id) VALUES (?, ?)`,
					dialogID, userID)
				if err != nil {
					panic(err)
				}
			}

			return dialogID
		}

		panic(err)
	}

	return dialogID
}
