package controller

import (
	"bts-tech-test-go/database"
	"bts-tech-test-go/model"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetAllItemByChecklistId(c *fiber.Ctx) error {
	checklistID := c.Params("checklistId")
	var items []model.Item
	if err := database.DB.Where("checklist_id = ?", checklistID).Find(&items).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get items"})
	}
	return c.JSON(items)
}

func CreateNewItemInChecklist(c *fiber.Ctx) error {
	checklistID, _ := strconv.ParseUint(c.Params("checklistId"), 10, 64)
	var input struct {
		ItemName string `json:"itemName"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	item := model.Item{
		ItemName:    input.ItemName,
		ChecklistId: uint(checklistID),
	}
	if err := database.DB.Create(&item).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create item"})
	}
	return c.JSON(item)
}

func GetDetailItem(c *fiber.Ctx) error {
	itemID := c.Params("checklistItemId")
	var item model.Item
	if err := database.DB.First(&item, itemID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
	}
	return c.JSON(item)
}

func ToggleStatusItem(c *fiber.Ctx) error {
	itemID := c.Params("checklistItemId")

	var item model.Item
	if err := database.DB.First(&item, itemID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Item not found",
		})
	}

	// Toggle status: 0 ↔ 1
	if item.Status == 0 {
		item.Status = 1
	} else {
		item.Status = 0
	}

	if err := database.DB.Save(&item).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update item status",
		})
	}

	statusStr := "undone"
	if item.Status == 1 {
		statusStr = "done"
	}

	return c.JSON(fiber.Map{
		"message": "Item marked as " + statusStr,
		"item":    item,
	})
}

func DeleteItem(c *fiber.Ctx) error {
	itemID := c.Params("checklistItemId")
	if err := database.DB.Delete(&model.Item{}, itemID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete item"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func RenameItem(c *fiber.Ctx) error {
	itemID := c.Params("checklistItemId")
	var input struct {
		ItemName string `json:"itemName"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	var item model.Item
	if err := database.DB.First(&item, itemID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
	}
	item.ItemName = input.ItemName
	database.DB.Save(&item)
	return c.JSON(item)
}
