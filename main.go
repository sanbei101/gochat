package main

import "go-chat/database"

func main() {
	if err := database.InitDB(); err != nil {
		panic(err)
	}
	if err := database.InsertFakeData(); err != nil {
		panic(err)
	}
}
