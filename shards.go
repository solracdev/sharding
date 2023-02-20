package main

import (
	"database/sql"
	"fmt"
	"log"
)

const shards = 5

var (
	shardDatabases = make([]*sql.DB, shards)
)

func init() {
	var err error
	for i := 0; i < shards; i++ {
		dsn := fmt.Sprintf("root:root@tcp(localhost:3306)/db_%d", i)
		shardDatabases[i], err = sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}
	}
}

func GetShard(key string) *sql.DB {
	shard := Hash([]byte(key)) % uint32(shards)
	return shardDatabases[shard]
}

func insert(db *sql.DB, uuid string) {
	_, err := db.Exec("INSERT INTO model (id, reference) VALUES (?, ?)", uuid, "SKU-XXX")
	if err != nil {
		log.Fatal("Error inserting data:", err)
	}
	log.Println("insert OK!")
}

func fetch(db *sql.DB, uuid string) {
	rows, err := db.Query("SELECT reference FROM model WHERE id = ?", uuid)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var reference string
		err = rows.Scan(&reference)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		fmt.Println("reference:", reference)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error after iterating over result set:", err)
	}
}
