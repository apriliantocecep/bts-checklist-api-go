package controller

import (
	"bts-tech-test-go/database"
	"bts-tech-test-go/model"
	"github.com/gofiber/fiber/v2"
)

func GetChecklist(c *fiber.Ctx) error {
	var checklists []model.Checklist
	database.DB.Find(&checklists)
	return c.JSON(checklists)
}

func CreateNewChecklist(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var input struct {
		Name string `json:"name"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	checklist := model.Checklist{
		Name:   input.Name,
		UserId: userID,
	}
	if err := database.DB.Create(&checklist).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create checklist"})
	}
	return c.JSON(checklist)
}

func DeleteChecklistById(c *fiber.Ctx) error {
	checklistID := c.Params("checklistId")
	if err := database.DB.Delete(&model.Checklist{}, checklistID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete checklist"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
