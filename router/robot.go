package router

import (
	"fmt"
	"io/ioutil"
	"robot-api/libs/cache"
	"robot-api/libs/config"
	"robot-api/libs/httpUtil"

	"github.com/kataras/iris/v12"
)

func RobotRoute(route iris.Party) {
	route.Post("/list", listRobotHandler)                           // 机器人列表
	route.Get("/get/{deviceId:string}", getRobotInfoHandler)        // 获取机器人详细信息
	route.Post("/create", createRobotHandler)                       // 创建机器人记录
	route.Post("/update", updateRobotHandler)                       // 更新机器人基本信息
	route.Get("/get/{deviceSn:string}/state", getRobotStateHandler) // 获取机器人最近状态
	route.Post("/deploy", setDeployHandler)                         // 设置机器人部署信息
	route.Post("/undeploy", undeployHandler)                        // 重置机器人部署信息
	route.Get("/delete/{deviceId:string}", removeRobotHandler)      // 删除机器人信息（物理删除）
	route.Post("/disable", disableRobotHandler)                     // 禁用机器人
	route.Post("/enable", enableRobotHandler)                       // 启用机器人

}

func listRobotHandler(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	resp := httpUtil.Post(fmt.Sprintf("%s%s", cache.Get(config.Conf.Self.RobotServiceName), "/robot/list"), data, "application/json")
	ctx.Write([]byte(resp))
}

func getRobotInfoHandler(ctx iris.Context) {
	deviceId := ctx.Params().Get("deviceId")
	resp := httpUtil.Get(fmt.Sprintf("%s%s%s", cache.Get(config.Conf.Self.RobotServiceName), "/robot/", deviceId))
	ctx.Write([]byte(resp))
}

func createRobotHandler(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	resp := httpUtil.Put(fmt.Sprintf("%s%s", cache.Get(config.Conf.Self.RobotServiceName), "/robot/"), data, "application/json")
	ctx.Write([]byte(resp))
}

func updateRobotHandler(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	deviceId, _ := data["id"].(string)
	resp := httpUtil.Put(fmt.Sprintf("%s%s%s", cache.Get(config.Conf.Self.RobotServiceName), "/robot/", deviceId), data, "application/json")
	ctx.Write([]byte(resp))
}

func getRobotStateHandler(ctx iris.Context) {
	deviceId := ctx.Params().Get("deviceId")
	resp := httpUtil.Get(fmt.Sprintf("%s%s%s%s", cache.Get(config.Conf.Self.RobotServiceName), "/robot/", deviceId, "/state"))
	ctx.Write([]byte(resp))
}

func setDeployHandler(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	deviceId, _ := data["deviceId"].(string)
	resp := httpUtil.Put(fmt.Sprintf("%s%s%s%s", cache.Get(config.Conf.Self.RobotServiceName), "/robot/", deviceId, "/deploy"), nil, "application/json")
	ctx.Write([]byte(resp))
}

func undeployHandler(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	deviceId, _ := data["deviceId"].(string)
	resp := httpUtil.Put(fmt.Sprintf("%s%s%s%s", cache.Get(config.Conf.Self.RobotServiceName), "/robot/", deviceId, "/undeploy"), nil, "application/json")
	ctx.Write([]byte(resp))
}

func removeRobotHandler(ctx iris.Context) {
	deviceId := ctx.Params().Get("deviceId")
	resp := httpUtil.Delete(fmt.Sprintf("%s%s%s", cache.Get(config.Conf.Self.RobotServiceName), "/robot/", deviceId))
	ctx.Write([]byte(resp))
}

func disableRobotHandler(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	deviceId, _ := data["deviceId"].(string)
	resp := httpUtil.Post(fmt.Sprintf("%s%s%s%s", cache.Get(config.Conf.Self.RobotServiceName), "/robot/", deviceId, "/disable"), nil, "application/json")
	ctx.Write([]byte(resp))
}

func enableRobotHandler(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	deviceId, _ := data["deviceId"].(string)
	resp := httpUtil.Post(fmt.Sprintf("%s%s%s%s", cache.Get(config.Conf.Self.RobotServiceName), "/robot/", deviceId, "/enable"), nil, "application/json")
	ctx.Write([]byte(resp))
}
