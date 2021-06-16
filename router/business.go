package router

import "github.com/kataras/iris/v12"

func BusinessRoute(route iris.Party) {
	route.Get("/", testFunc)
}
