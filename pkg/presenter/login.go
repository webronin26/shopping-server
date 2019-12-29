package presenter

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/webronin26/shopping-server/pkg/usecases/login"
)

type LoginParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// 登入
func Login(ctx echo.Context) error {

	var result Result

	var param LoginParam
	if err := ctx.Bind(&param); err != nil {
		result.Code = StatusParamError
		return ctx.JSON(http.StatusBadRequest, result)
	}
	if param.Email == "" || param.Password == "" {
		result.Code = StatusParamError
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var input login.Input
	input.Email = param.Email
	input.Password = param.Password

	output, err := login.Exec(input)
	if err != nil {
		result.Code = StatusAccountError
		return ctx.JSON(http.StatusNotFound, result)
	}

	result.Data = output
	result.Code = StatusSuccess

	return ctx.JSON(http.StatusOK, result)
}
