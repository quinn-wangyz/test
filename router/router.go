package router

import "github.com/kataras/iris/v12"

func Router(route iris.Party) {
	route.PartyFunc("/robot", RobotRoute)
	route.PartyFunc("/elevator", ElevatorRoute)
	route.PartyFunc("/building", BuildingRouter)
}
