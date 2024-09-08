package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lareii/copl.uk/server/models"
)

func GetUsers(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 10)
	offset := c.QueryInt("offset", 0)

	users, err := models.GetUsers(int64(limit), int64(offset))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching users.",
		})
	}

	var responseUsers []models.User
	for _, user := range users {
		user.Email = ""
		user.Password = ""

		responseUsers = append(responseUsers, user)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Users found.",
		"users":   responseUsers,
	})
}
