package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/webronin26/shopping-server/pkg/entities"
	"github.com/webronin26/shopping-server/pkg/presenter"
	"github.com/webronin26/shopping-server/pkg/store"
)

// 檢查 tolken 部分
func memberMiddleware(hfc echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		var result presenter.Result

		// 從 Header 當中取出 token
		token := ctx.Request().Header.Get("Authorization")
		if token == "" {
			result.Code = presenter.StatusTokenError
			return ctx.JSON(http.StatusUnauthorized, result)
		}

		account := new(entities.Account)

		// 檢查 token
		query := store.DB.Model(entities.Account{}).
			Where("token = ?", token).
			First(account)
		if query.RecordNotFound() {
			result.Code = presenter.StatusTokenError
			return ctx.JSON(http.StatusUnauthorized, result)
		}
		if query.Error != nil {
			result.Code = presenter.StatusServerError
			return ctx.JSON(http.StatusInternalServerError, result)
		}

		ctx.Set("account", account)

		return hfc(ctx)
	}
}
