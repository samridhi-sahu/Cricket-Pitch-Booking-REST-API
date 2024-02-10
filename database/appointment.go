package database

import "github.com/samridhi-sahu/pich-booking-system/WithFiber/structs"

var Appointments []structs.Appointment

func AddAppointment() {
	Appointments = append(Appointments,
		structs.Appointment{Id: "apt_1", ResourceId: "res_1", Quantity: 10, StartTime: "2023-08-04T08:00:00Z", EndTime: "2023-08-04T10:00:00Z"},
		structs.Appointment{Id: "apt_2", ResourceId: "res_1", Quantity: 12, StartTime: "2023-08-04T10:00:00Z", EndTime: "2023-08-04T12:00:00Z"},
		structs.Appointment{Id: "apt_3", ResourceId: "res_1", Quantity: 10, StartTime: "2023-08-04T15:00:00Z", EndTime: "2023-08-04T17:00:00Z"},
		structs.Appointment{Id: "apt_4", ResourceId: "res_1", Quantity: 12, StartTime: "2023-08-04T18:00:00Z", EndTime: "2023-08-04T20:00:00Z"},

		structs.Appointment{Id: "apt_5", ResourceId: "res_2", Quantity: 12, StartTime: "2023-08-04T09:00:00Z", EndTime: "2023-08-04T11:00:00Z"},
		structs.Appointment{Id: "apt_6", ResourceId: "res_2", Quantity: 15, StartTime: "2023-08-04T11:00:00Z", EndTime: "2023-08-04T13:00:00Z"},
		structs.Appointment{Id: "apt_7", ResourceId: "res_2", Quantity: 12, StartTime: "2023-08-04T16:00:00Z", EndTime: "2023-08-04T18:00:00Z"},
		structs.Appointment{Id: "apt_8", ResourceId: "res_2", Quantity: 15, StartTime: "2023-08-04T18:00:00Z", EndTime: "2023-08-04T20:00:00Z"},

		structs.Appointment{Id: "apt_9", ResourceId: "res_3", Quantity: 10, StartTime: "2023-08-04T10:00:00Z", EndTime: "2023-08-04T12:00:00Z"},
		structs.Appointment{Id: "apt_10", ResourceId: "res_3", Quantity: 15, StartTime: "2023-08-04T13:00:00Z", EndTime: "2023-08-04T15:00:00Z"},
		structs.Appointment{Id: "apt_11", ResourceId: "res_3", Quantity: 10, StartTime: "2023-08-04T19:00:00Z", EndTime: "2023-08-04T21:00:00Z"},
	)
}
