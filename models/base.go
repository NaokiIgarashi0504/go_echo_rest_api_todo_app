package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"go_echo_rest_api/config"
	"log"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

// テーブル名を定義
const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)

func init() {
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
}

// userのUUIDを作成する関数
func CreateUUID() (uuidobj uuid.UUID) {
	// UUIDを作成
	uuidobj, _ = uuid.NewUUID()

	// UUIDを返す
	return uuidobj
}

// passwordをハッシュ化する関数
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
