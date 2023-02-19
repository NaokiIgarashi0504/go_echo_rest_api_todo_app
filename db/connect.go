package db

import (
	"database/sql"
	"fmt"
	"go_echo_rest_api/config"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	Db *sql.DB

	err error
)

// テーブル名を定義
const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)

func Init() *sql.DB {
	// データベースと接続
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)

	// データベースとの接続で失敗した場合は、ログを吐く
	if err != nil {
		log.Fatalln(err)
	}

	// usersテーブルがない場合は、usersテーブルを作成するコマンドを定義
	usersTableCreateCmd := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	// usersテーブルのCREATE文を実行
	Db.Exec(usersTableCreateCmd)

	// todosテーブルがない場合は、todosテーブルを作成するコマンドを定義
	todosTableCreateCmd := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`, tableNameTodo)

	// todosテーブルのCREATE文を実行
	Db.Exec(todosTableCreateCmd)

	// sessionsテーブルがない場合は、sessionsテーブルを作成するコマンドを定義
	sessionsTableCreateCmd := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id INTEGER,
		created_at DATETIME)`, tableNameSession)

	// sessionsテーブルのCREATE文を実行
	Db.Exec(sessionsTableCreateCmd)

	return Db
}
