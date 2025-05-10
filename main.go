package main

import (
	"tugas15/configs"
	"tugas15/databases/connection"
	"tugas15/databases/migration"
	"tugas15/routers"

	_ "github.com/lib/pq"
)

func main() {
	configs.Initiator()
	connection.Initiator()
	defer connection.DBConnection.Close()
	migration.Initiator(connection.DBConnection)
	routers.StartServer()
}
