package main

import (
	"database/sql"
	"github.com/brianvoe/gofakeit/v5"
	"math/rand"
)

func insertMessages(tx *sql.Tx, userIDs []uint64, messagesCount int) {
	rand.Shuffle(len(userIDs), func(i, j int) {
		userIDs[i], userIDs[j] = userIDs[j], userIDs[i]
	})
	userA, userB := userIDs[0], userIDs[1]
	dialogID := insertDialog(tx, userA, userB)

	for i := 0; i < messagesCount; i++ {
		res, err := tx.Exec(
			`INSERT INTO messages (dialog_id, sender_id, message) VALUES (?, ?, ?)`,
			dialogID,
			orUInts64(userA, userB),
			gofakeit.Sentence(rand.Intn(2)+20),
		)
		if err != nil {
			panic(err)
		}
		msgIntID, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}

		messageID := uint64(msgIntID)
		if rand.Intn(100) < 10 {
			insertAttachment(tx, messageID)
		}

		for rand.Intn(100) < 2 {
			voteID := insertVote(tx, messageID)

			voteOptionsCount := rand.Intn(4) + 3
			voteOptionIDs := insertVoteOptions(tx, voteID, voteOptionsCount)
			rand.Shuffle(len(voteOptionIDs), func(i, j int) {
				voteOptionIDs[i], voteOptionIDs[j] = voteOptionIDs[j], voteOptionIDs[i]
			})

			for i := 0; i < 2; i++ {
				_, err = tx.Exec(`INSERT INTO vote_logs (vote_option_id, user_id) VALUES (?, ?)`,
					voteOptionIDs[i], userIDs[i])
				if err != nil {
					panic(err)
				}
			}

		}
	}
}
