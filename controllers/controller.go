package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samridhi-sahu/pich-booking-system/WithFiber/database"
	"github.com/samridhi-sahu/pich-booking-system/WithFiber/helpers"
)

func GetAllResouces(c *fiber.Ctx) error {
	return c.JSON(database.Resources)
}

func GetAllDurations(c *fiber.Ctx) error {
	return c.JSON(database.Durations)
}

func GetBusinessHours(c *fiber.Ctx) error {
	resourceId, startTime, endTime := helpers.Params(c)
	return c.JSON(helpers.GetBusinessHoursList(resourceId, startTime, endTime))
}

func GetBlockHours(c *fiber.Ctx) error {
	resourceId, startTime, endTime := helpers.Params(c)
	return c.JSON(helpers.GetBlockHoursList(resourceId, startTime, endTime))

}

func GetAllAppointments(c *fiber.Ctx) error {
	resourceId, startTime, endTime := helpers.Params(c)
	return c.JSON(helpers.GetAppointmentHoursList(resourceId, startTime, endTime))

}

func GetAllResoucesHTML(c *fiber.Ctx) error {
	c.Render("resourcePage", database.Resources)
	return nil
}

func GetBusinessHoursHTML(c *fiber.Ctx) error {
	resourceId, startTime, endTime := helpers.Params(c)
	c.Render("businessHourPage", helpers.GetBusinessHoursList(resourceId, startTime, endTime))
	return nil
}

func GetBlockHoursHTML(c *fiber.Ctx) error {
	resourceId, startTime, endTime := helpers.Params(c)
	c.Render("blockHourPage", helpers.GetBlockHoursList(resourceId, startTime, endTime))
	return nil
}

func GetAllAppointmentsHTML(c *fiber.Ctx) error {
	resourceId, startTime, endTime := helpers.Params(c)
	c.Render("appointmentsPage", helpers.GetAppointmentHoursList(resourceId, startTime, endTime))
	return nil
}
