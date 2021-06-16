package router

import (
	"github.com/kataras/iris/v12"
)

func ElevatorRoute(route iris.Party) {
	route.Get("/", testFunc)
}

func testFunc(ctx iris.Context) {
	return
}
