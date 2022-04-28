package books

import (
	"github.com/eRaBxEs/go-fiber-api/pkg/common/models"
	"github.com/eRaBxEs/go-fiber-api/pkg/common/util"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	resp := util.Response{Code: 200}
	resp.Set("book", book)

	return c.Status(fiber.StatusOK).JSON(resp)
}
