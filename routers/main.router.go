package routers

import (
	"go_echo_rest_api/controllers"
	"net/http"

	"github.com/labstack/echo"
)

// func GenerateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
// 	// filesを定義
// 	var files []string

// 	// filenamesを回す
// 	for _, file := range filenames {
// 		// filesに代入
// 		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
// 	}

// 	// Mustはあらかじめtemplateをキャッシュしておいて効率良くする。ParseFilesは失敗の際にパニック状態になる。
// 	template := template.Must(template.ParseFiles(files...))

// 	// defineを使用している場合は、ExecuteTemplateを使用して明示的に宣言する必要がある
// 	template.ExecuteTemplate(w, "layout", data)
// }

func NewMainRouter() *echo.Echo {
	e := echo.New()

	// 「/」（トップページ）に訪れた場合
	e.GET("/", echo.WrapHandler(http.HandlerFunc(controllers.Top)))

	// 「/signup」（サインアップページ）に訪れた場合
	e.GET("/signup", echo.WrapHandler(http.HandlerFunc(controllers.ShowSignUpFrom)))

	// 「/signup」（サインアップページ）でユーザーを登録した場合
	e.POST("/signup", echo.WrapHandler(http.HandlerFunc(controllers.CreateUser)))

	// 「/login」（ログインページ）に訪れた場合
	e.GET("/login", echo.WrapHandler(http.HandlerFunc(controllers.ShowLoginFrom)))

	// 「/logout」（todo一覧ページ）に訪れた場合
	e.GET("/logout", echo.WrapHandler(http.HandlerFunc(controllers.Logout)))

	// 「/signup」（サインアップページ）でユーザーを登録した場合
	e.POST("/authenticate", echo.WrapHandler(http.HandlerFunc(controllers.Authenticate)))

	// 「/todos」（todo一覧ページ）に訪れた場合
	e.GET("/todos", echo.WrapHandler(http.HandlerFunc(controllers.Index)))

	// 「/todos/new」（todo作成）に訪れた場合
	e.GET("/todos/new", echo.WrapHandler(http.HandlerFunc(controllers.TodoNew)))

	// 「/todos/save」（todo作成ページ）でtodoを登録した場合
	e.POST("/todos/save", echo.WrapHandler(http.HandlerFunc(controllers.TodoSave)))

	// 「/todos/edit/」（todo編集ページ）に訪れた場合
	e.GET("/todos/edit/", echo.WrapHandler(http.HandlerFunc(controllers.TodoEdit)))

	// 「/todos/update/」（todo編集ページ）でtodoを編集した場合
	e.POST("/todos/update/", echo.WrapHandler(http.HandlerFunc(controllers.TodoUpdate)))

	// 「/todos/delete/」（todo編集ページ）でtodoを編集した場合
	e.GET("/todos/delete/", echo.WrapHandler(http.HandlerFunc(controllers.TodoDelete)))

	return e
}
