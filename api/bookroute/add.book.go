package bookroute

import (
	"fiber-stats/models"
	"fiber-stats/util"

	"github.com/gofiber/fiber/v2"
)

type RequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h handler) AddBook(c *fiber.Ctx) error {
	body := RequestBody{}

	//parse request
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	book := models.Book{
		Title:       body.Title,
		Author:      body.Author,
		Description: body.Description,
	}

	// insert
	if result := h.DB.Create(&book); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	resp := util.Response{Code: 200}
	resp.Set("books", book)

	return c.Status(fiber.StatusCreated).JSON(resp)
}
