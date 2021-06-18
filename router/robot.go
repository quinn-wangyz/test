package router

import (
	"github.com/kataras/iris/v12"
)

func RobotRoute(route iris.Party) {
	route.Get("/list", testFunc1) //机器人列表

}

func testFunc1(ctx iris.Context) {
	return
}
