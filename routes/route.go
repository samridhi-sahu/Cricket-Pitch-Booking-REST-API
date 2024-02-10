package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samridhi-sahu/pich-booking-system/WithFiber/controllers"
)

func Routes(app *fiber.App) error {
	app.Get("/resources", controllers.GetAllResouces)
	app.Get("/durations", controllers.GetAllDurations)
	app.Get("/businesshours", controllers.GetBusinessHours)
	app.Get("/blockhours", controllers.GetBlockHours)
	app.Get("/appointments", controllers.GetAllAppointments)
	app.Get("/resources/page", controllers.GetAllResoucesHTML)
	app.Get("/businesshours/page", controllers.GetBusinessHoursHTML)
	app.Get("/blockhours/page", controllers.GetBlockHoursHTML)
	app.Get("/appointments/page", controllers.GetAllAppointmentsHTML)

	return nil
}
