package repositories

import (
	"go_echo_rest_api/models"
	"log"
	"time"
)

// ユーザー登録の処理
func CreateUser(user *models.User) (err error) {
	// ユーザーを作成する、insert文を定義
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	// insert文を実行
	_, err = models.Db.Exec(cmd,
		models.CreateUUID(),
		user.Name,
		user.Email,
		models.Encrypt(user.PassWord),
		time.Now())

	// エラーの場合はログを吐く
	if err != nil {
		log.Fatalln(err)
	}

	// エラーを返す
	return err
}

// 入力されたemailからユーザー情報を取得する関数
func GetUserByEmail(email string) (user models.User, err error) {
	// ユーザー情報を取得するコマンドを定義
	cmd := `SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?`

	// ユーザー情報を取得するコマンドを実行
	err = models.Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt)

	return user, err
}

// セッションを作成する関数
func CreateSession(user models.User) (session models.Session, err error) {
	// セッションを定義
	session = models.Session{}

	// セッションを登録する、insert文を定義
	cmd1 := `INSERT INTO sessions(
		uuid,
		email,
		user_id,
		created_at) VALUES (?, ?, ?, ?)`

	// insert文を実行
	_, err = models.Db.Exec(cmd1, models.CreateUUID(), user.Email, user.ID, time.Now())

	// エラーの場合はログを吐く
	if err != nil {
		log.Fatalln(err)
	}

	// 作成したセッションを取得する、select文を定義
	cmd2 := `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = ? and email = ?`

	// select文を実行して、変数sessionに代入
	err = models.Db.QueryRow(cmd2, user.ID, user.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)

	// 変数sessionとエラーを返す
	return session, err
}

// ログアウトの関数
func DeleteSessionByUUID(sess models.Session) (err error) {
	// セッションを削除する、delete文を定義
	cmd := `DELETE FROM sessions WHERE uuid = ?`

	// コマンドを実行
	_, err = models.Db.Exec(cmd, sess.UUID)

	// エラーの場合はログを吐く
	if err != nil {
		log.Fatalln(err)
	}

	// エラーを返す
	return err
}

func GetUserBySession(sess models.Session) (user models.User, err error) {
	// ユーザーを定義
	user = models.User{}

	// ユーザー情報を取得する、select文を定義
	cmd := `SELECT id, uuid, name, email, created_at FROM users WHERE id = ?`

	// select文を実行して、変数userに代入
	err = models.Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt)

	// ユーザーの情報とエラーを返す
	return user, err
}

// セッションに存在するかチェックする関数
func CheckSession(sess *models.Session) (valid bool, err error) {
	// セッションの情報を取得するコマンドを定義
	cmd := `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = ?`

	// セッションの情報を取得するコマンドの実行
	err = models.Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt)

	// エラーの場合はセッションが存在しない
	if err != nil {
		valid = false
		return
	}

	// セッションのUUIDが初期値でない場合は、セッションが存在する
	if sess.ID != 0 {
		valid = true
	}

	return valid, err
}
