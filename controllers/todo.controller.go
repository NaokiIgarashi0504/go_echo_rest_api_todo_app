package controllers

import (
	"go_echo_rest_api/services"
	"net/http"
)

// トップページの表示の処理
func Top(w http.ResponseWriter, r *http.Request) {
	// セッションチェック
	_, err := services.Session(w, r)

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
func Index(w http.ResponseWriter, r *http.Request) {
	// セッションチェック
	sess, err := services.Session(w, r)

	if err != nil {
		// セッションがない場合
		// トップページに遷移
		http.Redirect(w, r, "/", 302)
	} else {
		// セッションがある場合
		// todo.servise.goのIndexを呼ぶ
		userData := services.Index(w, r, sess)

		// htmlファイルを表示する関数を呼ぶ
		services.GenerateHTML(w, userData, "layout", "private_navbar", "index")
	}
}

// todo作成画面を表示する処理
func TodoNew(w http.ResponseWriter, r *http.Request) {
	// セッションチェック
	_, err := services.Session(w, r)

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
func TodoSave(w http.ResponseWriter, r *http.Request) {
	// セッションチェック
	sess, err := services.Session(w, r)

	if err != nil {
		// セッションがない場合は、ログインページに遷移
		http.Redirect(w, r, "/login", 302)
	} else {
		// セッションがある場合
		// todo.servise.goのTodoSaveを呼ぶ
		services.TodoSave(w, r, sess)

		http.Redirect(w, r, "/todos", 302)
	}
}

// todoの編集画面を表示する処理
func TodoEdit(w http.ResponseWriter, r *http.Request) {
	// セッションチェック
	sess, err := services.Session(w, r)

	if err != nil {
		// セッションがない場合は、ログインページに遷移
		http.Redirect(w, r, "/login", 302)
	} else {
		// セッションがある場合
		// todo.servise.goのTodoEditを呼ぶ
		todo := services.TodoEdit(w, r, sess, r.FormValue("id"))

		// htmlファイルを表示する関数を呼ぶ
		services.GenerateHTML(w, todo, "layout", "private_navbar", "todo_edit")
	}
}

// todoの編集の処理
func TodoUpdate(w http.ResponseWriter, r *http.Request) {
	// セッションチェック
	sess, err := services.Session(w, r)

	if err != nil {
		// セッションがない場合は、ログインページに遷移
		http.Redirect(w, r, "/login", 302)
	} else {
		// セッションがある場合
		// todo.servise.goのTodoUpdateを呼ぶ
		services.TodoUpdate(w, r, sess, r.FormValue("id"))

		// htmlファイルを表示する関数を呼ぶ
		http.Redirect(w, r, "/todos", 302)
	}
}

// todoの削除の処理
func TodoDelete(w http.ResponseWriter, r *http.Request) {
	// セッションチェック
	sess, err := services.Session(w, r)

	if err != nil {
		// セッションがない場合は、ログインページに遷移
		http.Redirect(w, r, "/login", 302)
	} else {
		// セッションがある場合
		// todo.servise.goのTodoUpdateを呼ぶ
		services.TodoDelete(w, r, sess, r.FormValue("id"))

		// htmlファイルを表示する関数を呼ぶ
		http.Redirect(w, r, "/todos", 302)
	}
}
