package router

import "github.com/kataras/iris/v12"

func AreaRouter(route iris.Party) {
	route.Get("/", testFunc)
}
