package books

import (
	"github.com/eRaBxEs/go-fiber-api/pkg/common/models"
	"github.com/eRaBxEs/go-fiber-api/pkg/common/util"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetBooks(c *fiber.Ctx) error {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	resp := util.Response{Code: 200}
	resp.Set("book", books)

	return c.Status(fiber.StatusOK).JSON(resp)
}
