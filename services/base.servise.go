package services

import (
	"fmt"
	"go_echo_rest_api/models"
	"go_echo_rest_api/repositories"
	"net/http"
	"text/template"
)

// htmlファイルを表示する処理
func GenerateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	// filesを定義
	var files []string

	// filenamesを回す
	for _, file := range filenames {
		// filesに代入
		files = append(files, fmt.Sprintf("views/templates/%s.html", file))
	}

	// Mustはあらかじめtemplateをキャッシュしておいて効率良くする。ParseFilesは失敗の際にパニック状態になる。
	template := template.Must(template.ParseFiles(files...))

	// defineを使用している場合は、ExecuteTemplateを使用して明示的に宣言する必要がある
	template.ExecuteTemplate(w, "layout", data)
}

// セッションをチェックする処理
func Session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	// 設定しているcookieを取得
	cookie, err := r.Cookie("_cookie")

	// errが空の場合
	if err == nil {
		// 設定しているcookieの値を変数に代入
		sess = models.Session{UUID: cookie.Value}

		// セッションのチェック
		if ok, _ := repositories.CheckSession(&sess); !ok {
			// チェックで正しくなかった場合は、無効なセッション設定
			err = fmt.Errorf("Invalid session")
		}
	}

	return sess, err
}
