package login

import (
	"errors"
	"fmt"

	"github.com/webronin26/shopping-server/pkg/entities"
	"github.com/webronin26/shopping-server/pkg/store"
	"github.com/webronin26/shopping-server/pkg/util"
)

type Input struct {
	Email    string
	Password string
}

type OutPut struct {
	Token string `json:"token"`
}

func Exec(input Input) (OutPut, error) {

	var output OutPut
	var account entities.Account

	query := store.DB.Model(entities.Account{}).
		Where("email = ? AND password = ?", input.Email, input.Password).
		Limit(1).
		Scan(&account)
	if err := query.Error; err != nil {
		fmt.Println(err.Error())
		return output, errors.New("can't find account")
	}

	// 如果帳號密碼都存在並且正確，產生一組 token，寫入資料庫，並且回傳 token 給使用者
	token, err := util.CreateToken(input.Email, input.Password)
	if err != nil {
		fmt.Println(err.Error())
		return output, errors.New("create token error")
	}

	update := store.DB.Model(entities.Account{}).
		Where("email = ? AND password = ?", input.Email, input.Password).
		Update("token", token)
	if err := update.Error; err != nil {
		fmt.Println(err.Error())
		return output, errors.New("update token error")
	}

	output.Token = token

	return output, nil
}
