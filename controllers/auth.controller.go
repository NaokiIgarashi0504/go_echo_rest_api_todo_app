package controllers

import (
	"go_echo_rest_api/services"
	"log"
	"net/http"
)

type AuthController interface {
	ShowSignUpFrom(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	ShowLoginFrom(w http.ResponseWriter, r *http.Request)
	Authenticate(w http.ResponseWriter, r *http.Request)
}

type authController struct {
	as services.AuthService
	bs services.BaseService
}

func NewAuthController(as services.AuthService, bs services.BaseService) AuthController {
	return &authController{as, bs}
}

// サインアップページを表示する処理
func (ac *authController) ShowSignUpFrom(w http.ResponseWriter, r *http.Request) {
	// セッションチェック
	_, err := ac.bs.Session(w, r)

	if err != nil {
		// セッションがない場合
		// htmlファイルを表示する関数を呼ぶ
		services.GenerateHTML(w, nil, "layout", "public_navbar", "signup")
	} else {
		// セッションがある場合
		// todo一覧ページに遷移
		http.Redirect(w, r, "/todos", 302)
	}
}

// ユーザー登録の処理
func (ac *authController) CreateUser(w http.ResponseWriter, r *http.Request) {
	// auth.servise.goのCreateUserを呼ぶ
	err := ac.as.CreateUser(w, r)

	// エラーの場合はログを吐く
	if err != nil {
		log.Fatalln(err)
	}

	// トップページに遷移
	http.Redirect(w, r, "/", 302)
}

// ログインページを表示する処理
func (ac *authController) ShowLoginFrom(w http.ResponseWriter, r *http.Request) {
	// セッションチェック
	_, err := ac.bs.Session(w, r)

	if err != nil {
		// セッションがない場合
		// htmlファイルを表示する関数を呼ぶ
		services.GenerateHTML(w, nil, "layout", "public_navbar", "login")
	} else {
		// セッションがある場合
		// todo一覧ページに遷移
		http.Redirect(w, r, "/todos", 302)
	}
}

// ログイン処理
func (ac *authController) Authenticate(w http.ResponseWriter, r *http.Request) {
	// auth.servise.goのAuthenticateを呼ぶ
	result := ac.as.Authenticate(w, r)

	if result {
		// 認証成功の場合は、トップページに遷移
		http.Redirect(w, r, "/", 302)
	} else {
		// 認証失敗の場合は、ログインページに遷移
		http.Redirect(w, r, "/login", 302)
	}
}

func (ac *authController) Logout(w http.ResponseWriter, r *http.Request) {
	// auth.servise.goのAuthenticateを呼ぶ
	result := ac.as.Logout(w, r)

	if result {
		// セッションの削除が完了したらログイン画面にリダイレクト
		http.Redirect(w, r, "/login", 302)
	}
}
