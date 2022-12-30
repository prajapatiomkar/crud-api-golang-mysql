package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prajapatiomkar/crud-api-golang-mysql/handler"
)

func SetupRoutes(app *fiber.App){

	app.Get("/employee",handler.GetAllEmployee)

	app.Post("/employee",handler.CreateEmployee) 
	app.Put("/employee",handler.UpdateEmployee) 
	app.Delete("/employee/:id",handler.DeleteEmployee) 
	app.Get("/employee/:id",handler.GetEmployeeById) 

}