package router

import (
	"fmt"
	"io/ioutil"
	"robot-api/libs/cache"
	"robot-api/libs/httpUtil"

	"github.com/golang/glog"
	"github.com/kataras/iris/v12"
	response "github.com/zhangsq-ax/iris-response-helper"
)

func BuildingRouter(route iris.Party) {
	route.Post("/list", BuildingListFunc) //楼宇列表
	route.Post("/list/test", BuildingListTestFunc)
}

func BuildingListFunc(ctx iris.Context) {
	fmt.Println(cache.Get("robot-api"))
	res := response.NewResponseHelper(&ctx, glog.Errorln)
	defer res.ResponseJSON("Failed to write response")
	s, _ := ioutil.ReadAll(ctx.Request().Body) //把	body 内容读入字符串 s
	fmt.Println(string(s))
	fmt.Println(cache.Get("robot-api") + "/list/test")
	returnValue := httpUtil.Post(cache.Get("robot-api")+"/building/list/test", ctx.Request().Body, "application/json")
	fmt.Println(returnValue)
}

func BuildingListTestFunc(ctx iris.Context) {
	fmt.Println("----------------------")
	res := response.NewResponseHelper(&ctx, glog.Errorln)
	defer res.ResponseJSON("Failed to write response")
	s, _ := ioutil.ReadAll(ctx.Request().Body) //把	body 内容读入字符串 s
	fmt.Println(string(s))
}
