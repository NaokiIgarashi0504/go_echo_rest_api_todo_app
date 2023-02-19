package routers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type MainRouter interface {
	setupRouting() *echo.Echo
	StartWebServer() error
}

type mainRouter struct {
	ar AuthRouter
	tr TodoRouter
}

func NewMainRouter(ar AuthRouter, tr TodoRouter) MainRouter {
	return &mainRouter{ar, tr}
}

/*
ルーティング定義
*/
func (mr *mainRouter) setupRouting() *echo.Echo {
	e := echo.New()

	mr.ar.SetAuthRouting(e)
	mr.tr.SetTodoRouting(e)

	return e
}

/*
サーバー起動
*/
func (mr *mainRouter) StartWebServer() error {
	fmt.Println("サーバー起動")
	// // ルーティング設定
	// setupRouting()

	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), mr.setupRouting())
}

// func NewMainRouter() *echo.Echo {
// 	e := echo.New()

// 	// 「/」（トップページ）に訪れた場合
// 	e.GET("/", echo.WrapHandler(http.HandlerFunc(controllers.Top)))

// 	// 「/signup」（サインアップページ）に訪れた場合
// 	e.GET("/signup", echo.WrapHandler(http.HandlerFunc(controllers.ShowSignUpFrom)))

// 	// 「/signup」（サインアップページ）でユーザーを登録した場合
// 	e.POST("/signup", echo.WrapHandler(http.HandlerFunc(controllers.CreateUser)))

// 	// 「/login」（ログインページ）に訪れた場合
// 	e.GET("/login", echo.WrapHandler(http.HandlerFunc(controllers.ShowLoginFrom)))

// 	// 「/logout」（todo一覧ページ）に訪れた場合
// 	e.GET("/logout", echo.WrapHandler(http.HandlerFunc(controllers.Logout)))

// 	// 「/signup」（サインアップページ）でユーザーを登録した場合
// 	e.POST("/authenticate", echo.WrapHandler(http.HandlerFunc(controllers.Authenticate)))

// 	todos := e.Group("/todos")
// 	// 「/todos」（todo一覧ページ）に訪れた場合
// 	todos.GET("", echo.WrapHandler(http.HandlerFunc(controllers.Index)))

// 	// 「/todos/new」（todo作成）に訪れた場合
// 	todos.GET("/new", echo.WrapHandler(http.HandlerFunc(controllers.TodoNew)))

// 	// 「/todos/save」（todo作成ページ）でtodoを登録した場合
// 	todos.POST("/save", echo.WrapHandler(http.HandlerFunc(controllers.TodoSave)))

// 	// 「/todos/edit/」（todo編集ページ）に訪れた場合
// 	todos.GET("/edit/", echo.WrapHandler(http.HandlerFunc(controllers.TodoEdit)))

// 	// 「/todos/update/」（todo編集ページ）でtodoを編集した場合
// 	todos.POST("/update/", echo.WrapHandler(http.HandlerFunc(controllers.TodoUpdate)))
// 	// 今後フロントエンドのReactで実装して、PUTで送信されるようにする（現在はHTMLだからGETかPOSTのみ）
// 	// todos.PUT("/update/", echo.WrapHandler(http.HandlerFunc(controllers.TodoUpdate)))

// 	// 「/todos/delete/」（todo編集ページ）でtodoを編集した場合
// 	todos.GET("/delete/", echo.WrapHandler(http.HandlerFunc(controllers.TodoDelete)))
// 	// 今後フロントエンドのReactで実装して、DELETEで送信されるようにする（現在はHTMLだからGETかPOSTのみ）
// 	// todos.DELETE("/delete/", echo.WrapHandler(http.HandlerFunc(controllers.TodoDelete)))

// 	// // 「/todos」（todo一覧ページ）に訪れた場合
// 	// e.GET("/todos", echo.WrapHandler(http.HandlerFunc(controllers.Index)))

// 	// // 「/todos/new」（todo作成）に訪れた場合
// 	// e.GET("/todos/new", echo.WrapHandler(http.HandlerFunc(controllers.TodoNew)))

// 	// // 「/todos/save」（todo作成ページ）でtodoを登録した場合
// 	// e.POST("/todos/save", echo.WrapHandler(http.HandlerFunc(controllers.TodoSave)))

// 	// // 「/todos/edit/」（todo編集ページ）に訪れた場合
// 	// e.GET("/todos/edit/", echo.WrapHandler(http.HandlerFunc(controllers.TodoEdit)))

// 	// // 「/todos/update/」（todo編集ページ）でtodoを編集した場合
// 	// e.POST("/todos/update/", echo.WrapHandler(http.HandlerFunc(controllers.TodoUpdate)))

// 	// // 「/todos/delete/」（todo編集ページ）でtodoを編集した場合
// 	// e.GET("/todos/delete/", echo.WrapHandler(http.HandlerFunc(controllers.TodoDelete)))

// 	return e
// }
