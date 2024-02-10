package database

import "github.com/samridhi-sahu/pich-booking-system/WithFiber/structs"

var Resources []structs.Resource

func AddResource() {
	Resources = append(Resources,
		structs.Resource{Id: "res_1", Name: "Pitch A", Description: "cricket club's - Pitch A"},
		structs.Resource{Id: "res_2", Name: "Pitch B", Description: "cricket club's - Pitch B"},
		structs.Resource{Id: "res_3", Name: "Pitch C", Description: "cricket club's - Pitch C"},
	)
}
