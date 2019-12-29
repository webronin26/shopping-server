package presenter

import (
	"net/http"

	"github.com/labstack/echo"
	main "github.com/webronin26/shopping-server/pkg/usecases/main_page"
)

func MainPage(ctx echo.Context) error {

	var result Result

	output, err := main.Exec()
	if err != nil {
		result.Code = StatusServerError
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Code = StatusSuccess
	result.Data = output

	return ctx.JSON(http.StatusOK, result)
}
