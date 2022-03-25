package user

import (
	"fiber-scaffold/app/repository"
	"fiber-scaffold/app/service/user_service"
	"fiber-scaffold/common/code"
	"fiber-scaffold/common/response"
	"fiber-scaffold/common/validate"
	"fiber-scaffold/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type auth struct {
	Username string `json:"username" validate:"required,min=5,max=32"`
	Password string `json:"password" validate:"required,min=5,max=32"`
}

type tokenInput struct {
	Token string `json:"token"`
}

func UserLogin(c *fiber.Ctx) error {

	loginForm := &auth{}
	if err := c.BodyParser(loginForm); err != nil {
		return err
	}
	validate := validate.ValidateStruct(loginForm)
	if validate != nil {
		return c.JSON(validate)
	}
	authService := user_service.Auth{Username: loginForm.Username, Password: loginForm.Password}
	isCheck := authService.Check()
	if !isCheck {
		return fiber.ErrUnauthorized
	}

	token, err := utils.Sign(loginForm.Username)
	if err != nil {
		return c.JSON(response.Response{Code: code.ERROR, Msg: code.MsgInfo(code.ERROR), Data: nil})
	}
	return c.JSON(response.Response{Code: code.SUCCESS, Msg: code.MsgInfo(code.SUCCESS), Data: map[string]interface{}{
		"accessToken": token,
	}})
}

func UserInfo(c *fiber.Ctx) error {
	tokenForm := &tokenInput{}
	if err := c.BodyParser(tokenForm); err != nil {
		return err
	}
	payload, _ := utils.Verity([]byte(tokenForm.Token))
	permission := repository.UserInfo(payload.Username)
	data := make(map[string]interface{})
	data["code"] = http.StatusOK
	data["msg"] = "success"
	data["data"] = map[string]interface{}{
		"permissions": []string{permission},
		"username":    payload.Username,
	}
	return c.JSON(data)
}
