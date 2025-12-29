package url

import (
	"github.com/gofiber/fiber/v2"
	"github.com/indrariksa/contactsAPI/controller"
)

func Web(page *fiber.App) {
	page.Get("/", controller.Home)
	page.Get("/contacts", controller.GetAll)
	page.Get("/contacts/:id", controller.GetKontakID)
	page.Post("/insert", controller.InsertData)
	page.Put("/update/:id", controller.UpdateData)
	page.Delete("/delete/:id", controller.DeleteKontak)
	page.Post("/login", controller.Login)

	// page.Post("/insert", middleware.JWTProtected(), controller.InsertData)

}
