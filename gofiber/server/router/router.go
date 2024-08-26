package router

import (
	"gofiber/controller"

	"github.com/gofiber/fiber/v2"
)

// setup routing information
func SetupRouters(app *fiber.App) {

	//list => get
	//add => post
	//update => put
	//delete => delete

	app.Get("/", controller.BlogList)
	app.Post("/", controller.BlogCreate)
	app.Put("/:id", controller.BlogUpdate)
	app.Delete("/:id", controller.BlogDelete)

}
