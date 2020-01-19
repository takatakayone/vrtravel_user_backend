package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/takatakayone/vrtravel_user_backend/src/config"
	"log"
)

var (
	UsersDb *sql.DB
)

func init() {
	dbConfig := config.GetDbConfig()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Schema)

	var err error
	UsersDb, err = sql.Open("mysql", dataSourceName)
	if err != nil{
		panic(err)
	}

	log.Println("database successfully configured")
}
