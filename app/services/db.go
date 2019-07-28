package services

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"plp/lib"
)

type DB_Connect struct {
	Conn *sql.DB
}

var DB_CONN DB_Connect

func (db_conn *DB_Connect) GetDb() *sql.DB {
	return db_conn.Conn
}

func init() {
	config := new(lib.Config)
	config.InitConfig("./app/conf/config.ini")

	dsn_tpl := "%s:%s@%s(%s:%s)/plp?charset=%s"
	db_user := config.Read("db", "user")
	db_password := config.Read("db", "password")
	db_protocol := config.Read("db", "protocol")
	db_server := config.Read("db", "server")
	db_port := config.Read("db", "port")
	db_charset := config.Read("db", "charset")
	dsn := fmt.Sprintf(dsn_tpl, db_user, db_password, db_protocol, db_server, db_port, db_charset)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	DB_CONN.Conn = db
}
