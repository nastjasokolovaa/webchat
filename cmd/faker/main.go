package main

import (
	"github.com/brianvoe/gofakeit/v5"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	gofakeit.Seed(time.Now().UnixNano())
}

func main() {
	db, err := dbConnect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	userIDs := insertUsers(tx)
	for i := 0; i < 10; i++ {
		insertMessages(tx, userIDs, rand.Intn(30)+15)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}

func orUInts64(ids ...uint64) uint64 {
	return ids[rand.Intn(len(ids))]
}

func orStrings(ids ...string) string {
	return ids[rand.Intn(len(ids))]
}
