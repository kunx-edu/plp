package services

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"plp/lib"
)

var DB_CONN *sql.DB

func InitDBConn() {

	dsn_tpl := "%s:%s@%s(%s:%s)/plp?charset=%s"
	db_user := lib.Read("db", "user")
	db_password := lib.Read("db", "password")
	db_protocol := lib.Read("db", "protocol")
	db_server := lib.Read("db", "server")
	db_port := lib.Read("db", "port")
	db_charset := lib.Read("db", "charset")
	dsn := fmt.Sprintf(dsn_tpl, db_user, db_password, db_protocol, db_server, db_port, db_charset)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败 "+dsn, err)
	}
	DB_CONN = db
}
