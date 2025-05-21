package router

import (
	"bts-tech-test-go/controller"
	"bts-tech-test-go/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Auth
	app.Post("/login", controller.Login)
	app.Post("/register", controller.Register)

	// Checklist
	checklist := app.Group("/checklist")
	checklist.Use(middleware.AuthMiddleware)
	checklist.Get("/", controller.GetChecklist)
	checklist.Post("/", controller.CreateNewChecklist)
	checklist.Delete("/:checklistId", controller.DeleteChecklistById)

	// Item
	checklist.Get("/:checklistId/item", controller.GetAllItemByChecklistId)
	checklist.Post("/:checklistId/item", controller.CreateNewItemInChecklist)
	checklist.Get("/:checklistId/item/:checklistItemId", controller.GetDetailItem)
	checklist.Put("/:checklistId/item/:checklistItemId", controller.ToggleStatusItem)
	checklist.Delete("/:checklistId/item/:checklistItemId", controller.DeleteItem)
	checklist.Put("/:checklistId/item/rename/:checklistItemId", controller.RenameItem)
}
