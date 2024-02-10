package database

import "github.com/samridhi-sahu/pich-booking-system/WithFiber/structs"

var BlockHours []structs.BlockHour

func AddBlockHour() {
	BlockHours = append(BlockHours,
		structs.BlockHour{Id: "block_1", ResourceId: "res_1", StartTime: "2023-08-04T17:00:00Z", EndTime: "2023-08-04T18:00:00Z"},

		structs.BlockHour{Id: "block_2", ResourceId: "res_2", StartTime: "2023-08-04T13:00:00Z", EndTime: "2023-08-04T14:00:00Z"},

		structs.BlockHour{Id: "block_3", ResourceId: "res_3", StartTime: "2023-08-04T12:00:00Z", EndTime: "2023-08-04T13:00:00Z"},
	)
}
