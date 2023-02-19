package controllers

import (
	"go_echo_rest_api/services"
	"net/http"
)

type TodoController interface {
	Top(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
	TodoNew(w http.ResponseWriter, r *http.Request)
	TodoSave(w http.ResponseWriter, r *http.Request)
	TodoEdit(w http.ResponseWriter, r *http.Request)
	TodoUpdate(w http.ResponseWriter, r *http.Request)
	TodoDelete(w http.ResponseWriter, r *http.Request)
}

type todoController struct {
	ts services.TodoService
	as services.AuthService
	bs services.BaseService
}

func NewTodoController(ts services.TodoService, as services.AuthService, bs services.BaseService) TodoController {
	return &todoController{ts, as, bs}
}

// トップページの表示の処理
func (tc *todoController) Top(w http.ResponseWriter, r *http.Request) {
	// セッションチェック
	_, err := tc.bs.Session(w, r)

	if err != nil {
		// セッションがない場合
		// htmlファイルを表示する関数を呼ぶ
		services.GenerateHTML(w, "hello", "layout", "public_navbar", "top")
	} else {
		// セッションがある場合
		// todo一覧ページに遷移
		http.Redirect(w, r, "/todos", 302)
	}
}

// todo一覧ページの処理
func (tc *todoController) Index(w http.ResponseWriter, r *http.Request) {
	// セッションチェック
	sess, err := tc.bs.Session(w, r)

	if err != nil {
		// セッションがない場合
		// トップページに遷移
		http.Redirect(w, r, "/", 302)
	} else {
		// セッションがある場合
		// todo.servise.goのIndexを呼ぶ
		userData := tc.ts.Index(w, r, sess)

		// htmlファイルを表示する関数を呼ぶ
		services.GenerateHTML(w, userData, "layout", "private_navbar", "index")
	}
}

// todo作成画面を表示する処理
func (tc *todoController) TodoNew(w http.ResponseWriter, r *http.Request) {
	// セッションチェック
	_, err := tc.bs.Session(w, r)

	if err != nil {
		// セッションがない場合は、ログインページに遷移
		http.Redirect(w, r, "/login", 302)
	} else {
		// セッションがある場合
		// htmlファイルを表示する関数を呼ぶ
		services.GenerateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

// todoの作成の処理
func (tc *todoController) TodoSave(w http.ResponseWriter, r *http.Request) {
	// セッションチェック
	sess, err := tc.bs.Session(w, r)

	if err != nil {
		// セッションがない場合は、ログインページに遷移
		http.Redirect(w, r, "/login", 302)
	} else {
		// セッションがある場合
		// todo.servise.goのTodoSaveを呼ぶ
		tc.ts.TodoSave(w, r, sess)

		http.Redirect(w, r, "/todos", 302)
	}
}

// todoの編集画面を表示する処理
func (tc *todoController) TodoEdit(w http.ResponseWriter, r *http.Request) {
	// セッションチェック
	sess, err := tc.bs.Session(w, r)

	if err != nil {
		// セッションがない場合は、ログインページに遷移
		http.Redirect(w, r, "/login", 302)
	} else {
		// セッションがある場合
		// todo.servise.goのTodoEditを呼ぶ
		todo := tc.ts.TodoEdit(w, r, sess, r.FormValue("id"))

		// htmlファイルを表示する関数を呼ぶ
		services.GenerateHTML(w, todo, "layout", "private_navbar", "todo_edit")
	}
}

// todoの編集の処理
func (tc *todoController) TodoUpdate(w http.ResponseWriter, r *http.Request) {
	// セッションチェック
	sess, err := tc.bs.Session(w, r)

	if err != nil {
		// セッションがない場合は、ログインページに遷移
		http.Redirect(w, r, "/login", 302)
	} else {
		// セッションがある場合
		// todo.servise.goのTodoUpdateを呼ぶ
		tc.ts.TodoUpdate(w, r, sess, r.FormValue("id"))

		// htmlファイルを表示する関数を呼ぶ
		http.Redirect(w, r, "/todos", 302)
	}
}

// todoの削除の処理
func (tc *todoController) TodoDelete(w http.ResponseWriter, r *http.Request) {
	// セッションチェック
	sess, err := tc.bs.Session(w, r)

	if err != nil {
		// セッションがない場合は、ログインページに遷移
		http.Redirect(w, r, "/login", 302)
	} else {
		// セッションがある場合
		// todo.servise.goのTodoUpdateを呼ぶ
		tc.ts.TodoDelete(w, r, sess, r.FormValue("id"))

		// htmlファイルを表示する関数を呼ぶ
		http.Redirect(w, r, "/todos", 302)
	}
}
