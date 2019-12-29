package server

import (
	"github.com/labstack/echo"
	"github.com/webronin26/shopping-server/pkg/presenter"
)

// 不必登入的 route
func generalRoute(e *echo.Echo) {
	e.GET("/main", presenter.MainPage)
	e.GET("/search", presenter.Search)
	e.GET("/query/product/:product_id", presenter.QueryProduct)

	e.POST("/login", presenter.Login)
}

// 會員的 route
func memberRoute(e *echo.Echo, middlewares ...echo.MiddlewareFunc) {
	routers := e.Group("/member", middlewares...)

	routers.POST("/buy", presenter.BuyProduct)
	routers.DELETE("/logout", presenter.Logout)
}
