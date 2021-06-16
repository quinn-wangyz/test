package router

import (
	"github.com/kataras/iris/v12"
)

func RobotRoute(route iris.Party) {
	route.Get("/", testFunc1)
}

func testFunc1(ctx iris.Context) {
	return
}
