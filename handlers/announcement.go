package handlers

import (
	"absent.com/absentapi/models"
	"absent.com/absentapi/repository"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// @Summary Create a todo.
// @Description create a single todo.
// @Tags todos
// @Accept json
// @Param todo body CreateTodoDTO true "Todo to create"
// @Produce json
// @Success 200 {object} CreateTodoResDTO
// @Router /todos [post]
func HandleAllAnnouncements(c *fiber.Ctx) error {
	limit := c.Query("pageSize")
	var limitNumber int
	var err error
	if limit == "" {
		limitNumber = 10
	} else {
		limitNumber, err = strconv.Atoi(limit)
	}

	results, err := repository.GetAnnouncements(limitNumber)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("")
	}
	return c.Status(200).JSON(results)
}

func HandleCreateAnnouncements(c *fiber.Ctx) error {
	announcement := new(models.AnnouncementRequest)
	err := c.BodyParser(announcement)
	if err != nil {
		err.Error()
	}
	createdAnnouncement, err := repository.CreateAnnouncement(announcement.Message)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(createdAnnouncement)
}
func HandleUpdateAnnouncements(c *fiber.Ctx) error {
	announcement := new(models.AnnouncementRequest)
	err := c.BodyParser(announcement)
	if err != nil {
		err.Error()
	}
	dbAnnouncement := announcement.ToAnnouncement()
	updateAnnouncement, err := repository.UpdateAnnouncement(*dbAnnouncement)

	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(updateAnnouncement)
}
func HandleGetOneAnnouncement(c *fiber.Ctx) error {
	announcement := new(models.Announcement)
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusInternalServerError).JSON("")
	}
	numId, err := strconv.Atoi(id)

	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON("")
	}
	announcement, err = repository.GetAnnouncement(uint(numId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("")
	}
	return c.Status(200).JSON(announcement)

}

func HandleDeleteAnnouncements(c *fiber.Ctx) error {
	item := c.Params("id")
	if item == "" {
		return c.Status(fiber.StatusNotAcceptable).JSON("")
	}

	numId, err := strconv.Atoi(item)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON("")
	}
	err = repository.DeleteAnnouncement(uint(numId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("")
	}

	return c.Status(fiber.StatusNoContent).JSON("")
}
