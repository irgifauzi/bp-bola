package url

import (
	"github.com/irgifauzi/bp-bola/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Web(page *fiber.App) {
	// page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	// page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth

	page.Get("/", controller.Sink)
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)
	page.Options("/", controller.Sink)

	page.Get("/checkip", controller.Homepage)
	page.Get("/pemain", controller.GetAllPemain)
	page.Get("/pemain/:id", controller.GetPemainID)
	page.Post("/insert", controller.InsertDataPemain)
	page.Put("/update/:id", controller.UpdateDataPemain)
	page.Delete("/delete/:id", controller.DeletePemainByID)
	page.Get("/docs/*", swagger.HandlerDefault)
}
