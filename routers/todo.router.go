package routers

import (
	"go_echo_rest_api/controllers"
	"net/http"

	"github.com/labstack/echo"
)

type TodoRouter interface {
	SetTodoRouting(e *echo.Echo)
}

type todoRouter struct {
	tc controllers.TodoController
}

func NewTodoRouter(tc controllers.TodoController) TodoRouter {
	return &todoRouter{tc}
}

func (tr *todoRouter) SetTodoRouting(e *echo.Echo) {
	// 「/」（トップページ）に訪れた場合
	e.GET("/", echo.WrapHandler(http.HandlerFunc(tr.tc.Top)))

	// 「/todos」（todo一覧ページ）に訪れた場合
	e.GET("/todos", echo.WrapHandler(http.HandlerFunc(tr.tc.Index)))

	// 「/todos/new」（todo作成）に訪れた場合
	e.GET("/todos/new", echo.WrapHandler(http.HandlerFunc(tr.tc.TodoNew)))

	// 「/todos/save」（todo作成ページ）でtodoを登録した場合
	e.POST("/todos/save", echo.WrapHandler(http.HandlerFunc(tr.tc.TodoSave)))

	// 「/todos/edit/」（todo編集ページ）に訪れた場合
	e.GET("/todos/edit/", echo.WrapHandler(http.HandlerFunc(tr.tc.TodoEdit)))

	// 「/todos/update/」（todo編集ページ）でtodoを編集した場合
	e.POST("/todos/update/", echo.WrapHandler(http.HandlerFunc(tr.tc.TodoUpdate)))

	// 「/todos/delete/」（todo編集ページ）でtodoを編集した場合
	e.GET("/todos/delete/", echo.WrapHandler(http.HandlerFunc(tr.tc.TodoDelete)))

}
