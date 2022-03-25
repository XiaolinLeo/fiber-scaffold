package middleware

import (
	"fiber-scaffold/common/code"
	"fiber-scaffold/utils"

	"github.com/gofiber/fiber/v2"
)

func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("accessToken")
		if token == "" {
			return c.JSON(fiber.Map{
				"code": code.VERYTY_Token_NULL,
				"msg":  code.MsgInfo(code.VERYTY_Token_NULL),
			})
		}
		payload, err := utils.Verity([]byte(token))
		if err != nil {
			return c.JSON(fiber.Map{
				"code": code.VERYTY_Token_Fail,
				"msg":  code.MsgInfo(code.VERYTY_Token_Fail),
			})
		}
		c.Set("username", payload.Username)
		return c.Next()
	}
}
