package blog

import (
	"fiber-scaffold/app/service/blog_service"
	"fiber-scaffold/common/code"
	"fiber-scaffold/common/response"
	"fiber-scaffold/common/validate"
	"github.com/gofiber/fiber/v2"
)

type blogForm struct {
	Title   string `json:"title" validate:"required`
	Content string `json:"content" validate:"required"`
}

func AddBlog(c *fiber.Ctx) error {
	bForm := &blogForm{}
	if err := c.BodyParser(bForm); err != nil {
		return err
	}
	validate := validate.ValidateStruct(bForm)
	if validate != nil {
		return c.JSON(validate)
	}
	blogSer := blog_service.Blog{Title: bForm.Title, Content: bForm.Content}
	if err := blogSer.Add(); err != nil {
		return c.JSON(response.Response{Code: code.ERROR, Msg: code.MsgInfo(code.ERROR), Data: nil})
	}

	return c.JSON(response.Response{Code: code.SUCCESS, Msg: code.MsgInfo(code.SUCCESS)})
}
