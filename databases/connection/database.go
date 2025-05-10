package connection

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
)

var DBConnection *sql.DB

func Initiator() {
	psqlInfo := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		viper.GetString("migration.db.postgres.db_user"),
		viper.GetString("migration.db.postgres.db_password"),
		viper.GetString("migration.db.postgres.db_host"),
		viper.GetInt("migration.db.postgres.db_port"),
		viper.GetString("migration.db.postgres.db_name"),
	)

	db, err := sql.Open(viper.GetString("migration.db.postgres.db_engine"), psqlInfo)
	if err != nil {
		panic(err)
	}

	DBConnection = db
	err = DBConnection.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
}
