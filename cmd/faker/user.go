package main

import (
	"database/sql"
	"github.com/brianvoe/gofakeit/v5"
	"math/rand"
)

func insertUsers(tx *sql.Tx) []uint64 {
	const usersToInsert = 25
	userIDs := make([]uint64, 0, usersToInsert)
	for i := 0; i <= usersToInsert; i++ {
		userIDs = append(userIDs, insertUser(tx))
	}
	return userIDs
}

func insertUser(tx *sql.Tx) uint64 {
	res, err := tx.Exec(
		`INSERT INTO users (name, phone, password, bio, photo_url) VALUES (?, ?, ?, ?, ?) 
				ON DUPLICATE KEY UPDATE id = id`,
		gofakeit.Name(),
		gofakeit.Number(7e11, 8e11-1),
		gofakeit.Password(true, true, true, false, false, 10),
		gofakeit.LoremIpsumSentence(rand.Intn(14)+4),
		gofakeit.URL(),
	)
	if err != nil {
		panic(err)
	}

	userID, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	return uint64(userID)
}
