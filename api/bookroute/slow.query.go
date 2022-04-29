package bookroute

import (
	"fmt"
	"strconv"
	"strings"

	"fiber-stats/models"
	"fiber-stats/util"

	"github.com/gofiber/fiber/v2"
)

func (h handler) SlowQuery(c *fiber.Ctx) error {
	var stats []models.PgStatStatement
	var cnt int64

	page := c.Query("page")
	offSet := c.Query("offset")
	query := c.Query("filter")

	if len(page) == 0 || len(offSet) == 0 || len(query) == 0 {
		if result := h.DB.Order("mean_time desc").Find(&stats); result.Error != nil {
			return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
		}
		// do the counting
		if err := h.DB.Model(&stats).Count(&cnt).Error; err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
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

	query = strings.ToUpper(query)

	query = fmt.Sprintf("%%%s%%", query)

	pageNum := 0

	if pageNo >= 0 && len(query) == 0 {
		if pageNo > 1 {
			pageNum = (pageNo - 1) * offset
		}

		err = h.DB.Offset(pageNum).Order("mean_time desc").Limit(offset).Find(&stats).Error

		// do the counting
		if err := h.DB.Model(&stats).Count(&cnt).Error; err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
	}

	if pageNo >= 0 && len(query) > 0 {
		if pageNo > 1 {
			pageNum = (pageNo - 1) * offset
		}
		err = h.DB.Offset(pageNum).Order("mean_time desc").Limit(offset).Where("query like ?", query).Find(&stats).Error

		// do the counting
		if err := h.DB.Model(&stats).Where("query like ?", query).Count(&cnt).Error; err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
	}

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	pagination := models.PageInfo{Page: pageNum, Size: offset, TotalCount: cnt}
	list := models.List{Data: stats, Pagination: pagination}
	resp := util.Response{Code: 200}
	resp.Set("queries-stats", list)

	return c.Status(fiber.StatusOK).JSON(resp)
}
