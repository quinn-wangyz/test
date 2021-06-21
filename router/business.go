package router

import (
	"fmt"
	"io/ioutil"
	"robot-api/libs/cache"
	"robot-api/libs/config"
	"robot-api/libs/httpUtil"

	"github.com/kataras/iris/v12"
)

func BusinessRoute(route iris.Party) {
	route.Post("/list", businessListHandler)                        // 获取业务列表
	route.Post("/add", businessAddHandler)                          // 创建业务
	route.Get("/get/{businessId:string}", getBusinessHandler)       // 获取业务信息
	route.Post("/update", updateBusinessHandler)                    // 更新业务信息
	route.Post("/disable", disableBusinessHandler)                  // 禁用业务
	route.Post("/enable", enableBusinessHandler)                    // 启用业务
	route.Get("/delete/{businessId:string}", removeBusinessHandler) // 删除业务
}

// 获取业务列表
func businessListHandler(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	resp := httpUtil.Post(fmt.Sprintf("%s%s", cache.Get(config.Conf.Self.BusinessServiceName), "/business/list"), data, "application/json")
	ctx.Write([]byte(resp))
}

// 创建业务
func businessAddHandler(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	resp := httpUtil.Put(fmt.Sprintf("%s%s", cache.Get(config.Conf.Self.BusinessServiceName), "/business/"), data, "application/json")
	ctx.Write([]byte(resp))
}

// 获取业务信息
func getBusinessHandler(ctx iris.Context) {
	businessId := ctx.Params().Get("businessId")
	resp := httpUtil.Get(fmt.Sprintf("%s%s%s", cache.Get(config.Conf.Self.BusinessServiceName), "/business/", businessId))
	ctx.Write([]byte(resp))
}

// 更新业务信息
func updateBusinessHandler(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	businessId, _ := data["id"].(string)
	resp := httpUtil.Put(fmt.Sprintf("%s%s%s", cache.Get(config.Conf.Self.BusinessServiceName), "/business/", businessId), data, "application/json")
	ctx.Write([]byte(resp))
}

// 禁用业务
func disableBusinessHandler(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	businessId, _ := data["id"].(string)
	resp := httpUtil.Post(fmt.Sprintf("%s%s%s%s", cache.Get(config.Conf.Self.BusinessServiceName), "/business/", businessId, "/disable"), nil, "application/json")
	ctx.Write([]byte(resp))
}

// 启用业务
func enableBusinessHandler(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	businessId, _ := data["id"].(string)
	resp := httpUtil.Post(fmt.Sprintf("%s%s%s%s", cache.Get(config.Conf.Self.BusinessServiceName), "/business/", businessId, "/enable"), nil, "application/json")
	ctx.Write([]byte(resp))
}

// 删除业务
func removeBusinessHandler(ctx iris.Context) {
	businessId := ctx.Params().Get("businessId")
	resp := httpUtil.Delete(fmt.Sprintf("%s%s%s", cache.Get(config.Conf.Self.BusinessServiceName), "/business/", businessId))
	ctx.Write([]byte(resp))
}
