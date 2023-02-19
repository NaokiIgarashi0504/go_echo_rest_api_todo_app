package services

import (
	"go_echo_rest_api/models"
	"go_echo_rest_api/repositories"
	"log"
	"net/http"
	"strconv"
)

type TodoService interface {
	Index(w http.ResponseWriter, r *http.Request, sess models.Session) (userData models.User)
	TodoSave(w http.ResponseWriter, r *http.Request, sess models.Session) (err error)
	TodoEdit(w http.ResponseWriter, r *http.Request, sess models.Session, id string) (todo models.Todo)
	TodoUpdate(w http.ResponseWriter, r *http.Request, sess models.Session, id string)
	TodoDelete(w http.ResponseWriter, r *http.Request, sess models.Session, id string)
}

type todoService struct {
	ar repositories.AuthRepository
	tr repositories.TodoRepository
}

func NewTodoService(ar repositories.AuthRepository, tr repositories.TodoRepository) TodoService {
	return &todoService{ar, tr}
}

// todo一覧の処理
func (ts *todoService) Index(w http.ResponseWriter, r *http.Request, sess models.Session) (userData models.User) {
	// セッションの情報からユーザーの情報を取得
	user, err := ts.ar.GetUserBySession(sess)

	// エラーの場合はログを吐く
	if err != nil {
		log.Fatalln(err)
	}

	// ユーザーに紐づくtodoを取得
	todos, _ := ts.tr.GetTodosByUser(sess, user.ID)

	// 取得したtodosを、ユーザーのTodosに代入
	user.Todos = todos

	// ユーザーのデータを返す
	return user
}

// todoを作成する処理
func (ts *todoService) TodoSave(w http.ResponseWriter, r *http.Request, sess models.Session) (err error) {
	// フォームの情報を取得
	err = r.ParseForm()

	// エラーの場合はログを吐く
	if err != nil {
		log.Fatalln(err)
	}

	// セッションの情報からユーザーの情報を取得
	user, err := ts.ar.GetUserBySession(sess)

	// エラーの場合はログを吐く
	if err != nil {
		log.Fatalln(err)
	}

	// POSTされたtodoの内容を変数contentに代入
	content := r.PostFormValue("content")

	// 新たなtodoを保存
	if err := ts.tr.CreateTodo(content, user.ID); err != nil {
		log.Fatalln(err)
	}

	// エラーを返す
	return err
}

// todoの編集画面を表示する処理
func (ts *todoService) TodoEdit(w http.ResponseWriter, r *http.Request, sess models.Session, id string) (todo models.Todo) {
	// セッションの情報からユーザーの情報を取得
	_, err := ts.ar.GetUserBySession(sess)

	// エラーの場合はログを吐く
	if err != nil {
		log.Fatalln(err)
	}

	// 渡されたIDを数値型にする
	todoId, _ := strconv.Atoi(id)

	// todoを取得する
	editTodo, err := ts.tr.GetTodo(todoId)

	// エラーの場合はログを吐く
	if err != nil {
		log.Fatalln(err)
	}

	return editTodo
}

// todoの編集の処理
func (ts *todoService) TodoUpdate(w http.ResponseWriter, r *http.Request, sess models.Session, id string) {
	// セッションの情報からユーザーの情報を取得
	user, err := ts.ar.GetUserBySession(sess)

	// エラーの場合はログを吐く
	if err != nil {
		log.Fatalln(err)
	}

	// POSTされたtodoの内容を変数contentに代入
	content := r.PostFormValue("content")

	// 渡されたIDを数値型にする
	todoId, _ := strconv.Atoi(id)

	// 編集後のtodoの情報を変数newTodoDataに代入
	newTodoData := &models.Todo{ID: todoId, Content: content, UserID: user.ID}

	// todoのupdateの実行
	if err := ts.tr.UpdateTodo(newTodoData); err != nil {
		log.Fatalln(err)
	}
}

// todoの削除の処理
func (ts *todoService) TodoDelete(w http.ResponseWriter, r *http.Request, sess models.Session, id string) {
	// セッションの情報からユーザーの情報を取得
	_, err := ts.ar.GetUserBySession(sess)

	// エラーの場合はログを吐く
	if err != nil {
		log.Fatalln(err)
	}

	// 渡されたIDを数値型にする
	todoId, _ := strconv.Atoi(id)

	// todoを取得する
	deleteTodo, err := ts.tr.GetTodo(todoId)

	// todoのdeleteの実行
	if err := ts.tr.DeleteTodo(deleteTodo); err != nil {
		log.Fatalln(err)
	}
}
