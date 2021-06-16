package router

import "github.com/kataras/iris/v12"

func BuildingRouter(route iris.Party) {
	route.Get("/", testFunc)
}
