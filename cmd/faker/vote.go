package main

import (
	"database/sql"
	"github.com/brianvoe/gofakeit/v5"
	"math/rand"
)

func insertVote(tx *sql.Tx, messageID uint64) uint64 {
	res, err := tx.Exec(`INSERT INTO votes (message_id, description) VALUES (
                ?, ?
            )`, messageID, gofakeit.Question())
	if err != nil {
		panic(err)
	}
	voteID, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return uint64(voteID)
}

func insertVoteOptions(tx *sql.Tx, voteID uint64, voteOptionsCount int) []uint64 {
	optionIDs := make([]uint64, 0, voteOptionsCount)
	for i := 0; i <= voteOptionsCount; i++ {
		res, err := tx.Exec(`INSERT INTO vote_options (vote_id, label) VALUES (
                ?, ?
            )`, voteID, gofakeit.Sentence(rand.Intn(2)+3))
		if err != nil {
			panic(err)
		}
		optionID, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}
		optionIDs = append(optionIDs, uint64(optionID))
	}

	return optionIDs
}
