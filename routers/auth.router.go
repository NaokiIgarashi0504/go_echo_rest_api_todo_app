package routers

import (
	"go_echo_rest_api/controllers"
	"net/http"

	"github.com/labstack/echo"
)

type AuthRouter interface {
	SetAuthRouting(e *echo.Echo)
}

type authRouter struct {
	ac controllers.AuthController
}

func NewAuthRouter(ac controllers.AuthController) AuthRouter {
	return &authRouter{ac}
}

func (ar *authRouter) SetAuthRouting(e *echo.Echo) {
	// 「/signup」（サインアップページ）に訪れた場合
	e.GET("/signup", echo.WrapHandler(http.HandlerFunc(ar.ac.ShowSignUpFrom)))

	// 「/signup」（サインアップページ）でユーザーを登録した場合
	e.POST("/signup", echo.WrapHandler(http.HandlerFunc(ar.ac.CreateUser)))

	// 「/login」（ログインページ）に訪れた場合
	e.GET("/login", echo.WrapHandler(http.HandlerFunc(ar.ac.ShowLoginFrom)))

	// 「/logout」（todo一覧ページ）に訪れた場合
	e.GET("/logout", echo.WrapHandler(http.HandlerFunc(ar.ac.Logout)))

	// 「/signup」（サインアップページ）でユーザーを登録した場合
	e.POST("/authenticate", echo.WrapHandler(http.HandlerFunc(ar.ac.Authenticate)))
}
