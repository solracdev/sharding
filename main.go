package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	uuid := "a6a16f70-8ac6-4102-a3eb-208e597f985f"
	db := GetShard(uuid)

	var dbName string
	err := db.QueryRow("SELECT DATABASE()").Scan(&dbName)
	if err != nil {
		fmt.Println("Error retrieving database name:", err)
		return
	}
	fmt.Println("Database:", dbName)

	// INSERT FLOW
	insert(db, uuid)

	// SELECT FLOW
	//fetch(db, uuid)
}
