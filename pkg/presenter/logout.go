package presenter

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/webronin26/shopping-server/pkg/entities"
	"github.com/webronin26/shopping-server/pkg/usecases/logout"
)

func Logout(ctx echo.Context) error {

	var result Result

	account := ctx.Get("account").(*entities.Account)

	var input logout.Input
	input.AccountID = account.ID

	if err := logout.Exec(input); err != nil {
		result.Code = StatusServerError
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Code = StatusSuccess

	return ctx.JSON(http.StatusOK, result)
}
