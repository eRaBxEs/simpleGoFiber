package router

import (
	"fmt"
	"strconv"
	"strings"

	"fiber-stats/models"
	"fiber-stats/util"

	"github.com/gofiber/fiber/v2"
)

type PageInfo struct {
	Page int
	Size int
}

type List struct {
	Data       []models.PgStatStatement
	Pagination PageInfo
}

func (h handler) SlowQuery(c *fiber.Ctx) error {
	var stats []models.PgStatStatement

	page := c.Query("page")
	offSet := c.Query("offset")
	filter := c.Query("filter")

	if len(page) == 0 || len(offSet) == 0 || len(filter) == 0 {
		if result := h.DB.Order("mean_time desc").Find(&stats); result.Error != nil {
			return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
		}
	}

	pageNo, err := strconv.Atoi(page)
	if err != nil {
		return err
	}

	offset, err := strconv.Atoi(offSet)
	if err != nil {
		return err
	}

	filter = strings.ToUpper(filter)

	filter = fmt.Sprint("%", filter, "%")

	pageNum := 0

	if pageNo >= 0 && len(filter) == 0 {
		if pageNo > 1 {
			pageNum = (pageNo - 1) * offset
		}

		err = h.DB.Offset(pageNum).Order("mean_time desc").Limit(offset).Find(&stats).Error

	}

	if pageNo >= 0 && len(filter) > 0 {
		if pageNo > 1 {
			pageNum = (pageNo - 1) * offset
		}
		err = h.DB.Offset(pageNum).Order("mean_time desc").Limit(offset).Where("query like ?", filter).Find(&stats).Error

	}

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	pagination := PageInfo{Page: pageNum, Size: offset}
	list := List{Data: stats, Pagination: pagination}
	resp := util.Response{Code: 200}
	resp.Set("queries-stats", list)

	return c.Status(fiber.StatusOK).JSON(resp)
}
