package helpers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/samridhi-sahu/pich-booking-system/WithFiber/database"
	"github.com/samridhi-sahu/pich-booking-system/WithFiber/structs"
)

func StringToTime(timeStr string) time.Time {
	t, _ := time.Parse(time.RFC3339, timeStr)
	return t
}

func HeaderSet(w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json")
}

func Params(c *fiber.Ctx) (string, time.Time, time.Time) {
	resourceId := c.Query("resourceId")
	st := c.Query("startTime")
	et := c.Query("endTime")
	startTime := StringToTime(st)
	endTime := StringToTime(et)
	return resourceId, startTime, endTime
}
func GetBusinessHoursList(resourceId string, startTime, endTime time.Time) []structs.BusinessHour {
	businessHours := []structs.BusinessHour{}
	for _, bh := range database.BusinessHours {
		businessSt := bh.StartTime
		businessEt := bh.EndTime
		businessStartTime := StringToTime(businessSt)
		businessEndTime := StringToTime(businessEt)
		if resourceId == bh.ResourceId && businessStartTime.After(startTime) && businessEndTime.Before(endTime) {
			businessHours = append(businessHours, bh)
		}
	}
	return businessHours
}
func GetBlockHoursList(resourceId string, startTime, endTime time.Time) []structs.BlockHour {
	blockHours := []structs.BlockHour{}
	for _, blh := range database.BlockHours {
		blockSt := blh.StartTime
		blockEt := blh.EndTime
		blockStartTime := StringToTime(blockSt)
		blockEndTime := StringToTime(blockEt)
		if resourceId == blh.ResourceId && blockStartTime.After(startTime) && blockEndTime.Before(endTime) {
			blockHours = append(blockHours, blh)
		}
	}
	return blockHours
}
func GetAppointmentHoursList(resourceId string, startTime, endTime time.Time) []structs.Appointment {
	appointmentHours := []structs.Appointment{}
	for _, ah := range database.Appointments {
		appointmentSt := ah.StartTime
		appointmentEt := ah.EndTime
		appointmentStartTime := StringToTime(appointmentSt)
		appointmentEndTime := StringToTime(appointmentEt)
		if resourceId == ah.ResourceId && appointmentStartTime.After(startTime) && appointmentEndTime.Before(endTime) {
			appointmentHours = append(appointmentHours, ah)
		}
	}
	return appointmentHours
}
