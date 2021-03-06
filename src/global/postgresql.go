package global

import (
	"database/sql"
	"fmt"
	"log"
	"sactools/src/conf"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	dbConf := conf.Conf.DataBase
	var err error

	DB, err = sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConf.Host, dbConf.Port, dbConf.Username, dbConf.Password, dbConf.DbName))
	if err != nil {
		log.Fatalln(err)
	}
}
