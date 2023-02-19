package repositories

import (
	"database/sql"
	"go_echo_rest_api/models"
	"log"
	"time"
)

type TodoRepository interface {
	GetTodosByUser(sess models.Session, userId int) (todos []models.Todo, err error)
	CreateTodo(content string, userId int) (err error)
	GetTodo(id int) (todo models.Todo, err error)
	UpdateTodo(newTodoData *models.Todo) error
	DeleteTodo(deleteTodo models.Todo) (err error)
}

type todoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &todoRepository{db}
}

// 特定のuserのtodoを取得する関数
func (tr *todoRepository) GetTodosByUser(sess models.Session, userId int) (todos []models.Todo, err error) {
	// ユーザーに紐づくデータを取得する、select文を定義
	cmd := `SELECT id, content, user_id, created_at FROM todos WHERE user_id = ?`

	// select文の実行
	rows, err := tr.db.Query(cmd, userId)

	// エラーの場合はログを吐く
	if err != nil {
		log.Fatalln(err)
	}

	// 取得したtodoを回す
	for rows.Next() {
		// 変数todoを定義
		var todo models.Todo

		// todoに代入
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)

		// エラーの場合はログを吐く
		if err != nil {
			log.Fatalln(err)
		}

		// todosに追加
		todos = append(todos, todo)
	}

	// DBをクローズ
	rows.Close()

	// todosとエラーを返す
	return todos, err
}

// todoを作成する関数
func (tr *todoRepository) CreateTodo(content string, userId int) (err error) {
	// todoを作成するinsert文を定義
	cmd := `INSERT INTO todos (
		content,
		user_id,
		created_at) VALUES (?, ?, ?)`

	// insert文の実行
	_, err = tr.db.Exec(cmd, content, userId, time.Now())

	// エラーの場合はログを吐く
	if err != nil {
		log.Fatalln(err)
	}

	// エラーを返す
	return err
}

// IDに紐づくtodoを取得する
func (tr *todoRepository) GetTodo(id int) (todo models.Todo, err error) {
	// todoを定義
	todo = models.Todo{}

	// idを指定して、todoを取得するselect文を定義
	cmd := `SELECT id, content, user_id, created_at FROM todos WHERE id = ?`

	// select文を実行して、todoに代入
	err = tr.db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt)

	// todoとエラーを返す
	return todo, err
}

// todoを更新する関数
func (tr *todoRepository) UpdateTodo(newTodoData *models.Todo) error {
	// todoを更新する、update文を定義
	cmd := `UPDATE todos SET content = ?, user_id = ? WHERE id = ?`

	// update分の実行
	_, err := tr.db.Exec(cmd, newTodoData.Content, newTodoData.UserID, newTodoData.ID)

	// エラーの場合はログを吐く
	if err != nil {
		log.Fatalln(err)
	}

	// エラーを返す
	return err
}

// todoを削除する関数
func (tr *todoRepository) DeleteTodo(deleteTodo models.Todo) (err error) {
	// todoを削除する、delete文を定義
	cmd := `DELETE FROM todos WHERE id = ?`

	// delete文の実行
	_, err = tr.db.Exec(cmd, deleteTodo.ID)

	// エラーの場合はログを吐く
	if err != nil {
		log.Fatalln(err)
	}

	// エラーを返す
	return err
}
