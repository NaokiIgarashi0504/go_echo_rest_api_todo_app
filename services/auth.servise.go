package services

import (
	"go_echo_rest_api/models"
	"go_echo_rest_api/repositories"
	"log"
	"net/http"
)

type AuthService interface {
	CreateUser(w http.ResponseWriter, r *http.Request) (err error)
	Authenticate(w http.ResponseWriter, r *http.Request) (result bool)
	Logout(w http.ResponseWriter, r *http.Request) (result bool)
}

type authService struct {
	ar repositories.AuthRepository
}

func NewAuthService(ar repositories.AuthRepository) AuthService {
	return &authService{ar}
}

// ユーザー登録の処理
func (as *authService) CreateUser(w http.ResponseWriter, r *http.Request) (err error) {
	// ParseFormでformの内容を解析
	err = r.ParseForm()

	// エラーの場合はログを吐く
	if err != nil {
		log.Fatalln(err)
	}

	// 変数userにPOSTされてきた値を代入
	user := models.User{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		PassWord: r.PostFormValue("password"),
	}

	// auth.repository.goのCreateUserに変数userを渡す
	if err := as.ar.CreateUser(&user); err != nil {
		log.Fatalln(err)
	}

	// エラーの場合はログを吐く
	if err != nil {
		log.Fatalln(err)
	}

	// エラーを返す
	return err
}

// ユーザーの認証の処理
func (as *authService) Authenticate(w http.ResponseWriter, r *http.Request) (result bool) {
	// フォームの情報を取得
	err := r.ParseForm()

	// エラーの場合はログを吐く
	if err != nil {
		log.Fatalln(err)
	}

	// emailの情報をからユーザー情報を取得する
	user, err := as.ar.GetUserByEmail(r.PostFormValue("email"))

	// エラーの場合はログを吐き、ログインページに遷移
	if err != nil {
		log.Fatalln(err)
		http.Redirect(w, r, "/login", 302)
	}

	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
		// パスワードが正しい場合
		// セッションを作成
		session, err := as.ar.CreateSession(user)

		// エラーの場合はログを吐く
		if err != nil {
			log.Fatalln(err)
		}

		// cookieの作成
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}

		// cookieをセット
		http.SetCookie(w, &cookie)

		// 認証成功
		return true
	} else {
		// パスワードが正くない場合
		// 認証失敗
		return false
	}
}

// ログアウトの処理
func (as *authService) Logout(w http.ResponseWriter, r *http.Request) (result bool) {
	// cookieを取得
	cookie, err := r.Cookie("_cookie")

	if err != nil {
		// エラーの場合はログを吐いて、falseを返す
		log.Fatalln(err)
		return false
	}

	// ErrNoCookieではない場合はセッションのストラクトを作成する
	if err != http.ErrNoCookie {
		// 設定したcookieの値を変数に代入
		session := models.Session{UUID: cookie.Value}

		// セッションを削除
		as.ar.DeleteSessionByUUID(session)
	}

	// セッションの削除が完了
	return true
}
