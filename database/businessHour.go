package database

import "github.com/samridhi-sahu/pich-booking-system/WithFiber/structs"

var BusinessHours []structs.BusinessHour

func AddBusinessHour() {
	BusinessHours = append(BusinessHours,
		structs.BusinessHour{Id: "bus_1", ResourceId: "res_1", Quantity: 15, StartTime: "2023-08-04T08:00:00Z", EndTime: "2023-08-04T13:00:00Z"},
		structs.BusinessHour{Id: "bus_2", ResourceId: "res_1", Quantity: 15, StartTime: "2023-08-04T15:00:00Z", EndTime: "2023-08-04T20:00:00Z"},

		structs.BusinessHour{Id: "bus_3", ResourceId: "res_2", Quantity: 15, StartTime: "2023-08-04T09:00:00Z", EndTime: "2023-08-04T14:00:00Z"},
		structs.BusinessHour{Id: "bus_4", ResourceId: "res_2", Quantity: 15, StartTime: "2023-08-04T16:00:00Z", EndTime: "2023-08-04T21:00:00Z"},

		structs.BusinessHour{Id: "bus_5", ResourceId: "res_3", Quantity: 15, StartTime: "2023-08-04T10:00:00Z", EndTime: "2023-08-04T15:00:00Z"},
		structs.BusinessHour{Id: "bus_6", ResourceId: "res_3", Quantity: 15, StartTime: "2023-08-04T17:00:00Z", EndTime: "2023-08-04T22:00:00Z"},
	)
}
