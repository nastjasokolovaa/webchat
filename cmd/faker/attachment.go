package main

import (
	"database/sql"
	"github.com/brianvoe/gofakeit/v5"
)

func insertAttachment(tx *sql.Tx, messageID uint64) {
	_, err := tx.Exec(`INSERT INTO attachments (message_id, attachment_type_id, content_url) VALUES (
            	?, (SELECT id FROM attachment_types WHERE alias = ?), ?
			)`, messageID, orStrings("photo", "link"), gofakeit.URL())
	if err != nil {
		panic(err)
	}
}
