package services

import (
	"go_echo_rest_api/models"
	"go_echo_rest_api/repositories"
	"log"
	"net/http"
	"strconv"
)

// todo一覧の処理
func Index(w http.ResponseWriter, r *http.Request, sess models.Session) (userData models.User) {
	// セッションの情報からユーザーの情報を取得
	user, err := repositories.GetUserBySession(sess)

	// エラーの場合はログを吐く
	if err != nil {
		log.Panicln(err)
	}

	// ユーザーに紐づくtodoを取得
	todos, _ := repositories.GetTodosByUser(sess, user.ID)

	// 取得したtodosを、ユーザーのTodosに代入
	user.Todos = todos

	// ユーザーのデータを返す
	return user
}

// todoを作成する処理
func TodoSave(w http.ResponseWriter, r *http.Request, sess models.Session) (err error) {
	// フォームの情報を取得
	err = r.ParseForm()

	// エラーの場合はログを吐く
	if err != nil {
		log.Panicln(err)
	}

	// セッションの情報からユーザーの情報を取得
	user, err := repositories.GetUserBySession(sess)

	// エラーの場合はログを吐く
	if err != nil {
		log.Panicln(err)
	}

	// POSTされたtodoの内容を変数contentに代入
	content := r.PostFormValue("content")

	// 新たなtodoを保存
	if err := repositories.CreateTodo(content, user.ID); err != nil {
		log.Panicln(err)
	}

	// エラーを返す
	return err
}

// todoの編集画面を表示する処理
func TodoEdit(w http.ResponseWriter, r *http.Request, sess models.Session, id string) (todo models.Todo) {
	// セッションの情報からユーザーの情報を取得
	_, err := repositories.GetUserBySession(sess)

	// エラーの場合はログを吐く
	if err != nil {
		log.Panicln(err)
	}

	// 渡されたIDを数値型にする
	todoId, _ := strconv.Atoi(id)

	// todoを取得する
	editTodo, err := repositories.GetTodo(todoId)

	// エラーの場合はログを吐く
	if err != nil {
		log.Panicln(err)
	}

	return editTodo
}

// todoの編集の処理
func TodoUpdate(w http.ResponseWriter, r *http.Request, sess models.Session, id string) {
	// セッションの情報からユーザーの情報を取得
	user, err := repositories.GetUserBySession(sess)

	// エラーの場合はログを吐く
	if err != nil {
		log.Panicln(err)
	}

	// POSTされたtodoの内容を変数contentに代入
	content := r.PostFormValue("content")

	// 渡されたIDを数値型にする
	todoId, _ := strconv.Atoi(id)

	// 編集後のtodoの情報を変数newTodoDataに代入
	newTodoData := &models.Todo{ID: todoId, Content: content, UserID: user.ID}

	// todoのupdateの実行
	if err := repositories.UpdateTodo(newTodoData); err != nil {
		log.Panicln(err)
	}
}

// todoの削除の処理
func TodoDelete(w http.ResponseWriter, r *http.Request, sess models.Session, id string) {
	// セッションの情報からユーザーの情報を取得
	_, err := repositories.GetUserBySession(sess)

	// エラーの場合はログを吐く
	if err != nil {
		log.Panicln(err)
	}

	// 渡されたIDを数値型にする
	todoId, _ := strconv.Atoi(id)

	// todoを取得する
	deleteTodo, err := repositories.GetTodo(todoId)

	// todoのdeleteの実行
	if err := repositories.DeleteTodo(deleteTodo); err != nil {
		log.Panicln(err)
	}
}
