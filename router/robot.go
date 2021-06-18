package router

import (
	"github.com/kataras/iris/v12"
)

func RobotRoute(route iris.Party) {
	route.Post("/", createRobotHandler)                           // 创建机器人记录
	route.Post("/list", listRobotHandler)                         // 机器人列表
	route.Post("/{deviceId:string}/deploy", setDeployHandler)     // 设置机器人部署信息
	route.Post("/{deviceId:string}/undeploy", undeployHandler)    // 重置机器人部署信息
	route.Get("/{deviceId:string}", removeRobotHandler)           // 删除机器人信息（物理删除）
	route.Post("/{deviceId:string}", updateRobotHandler)          // 更新机器人基本信息
	route.Post("/{deviceId:string}/disable", disableRobotHandler) // 禁用机器人
	route.Post("/{deviceId:string}/enable", enableRobotHandler)   // 启用机器人
	route.Get("/{deviceId:string}", getRobotInfoHandler)          // 获取机器人详细信息
	route.Get("/{deviceSn:string}/state", getRobotStateHandler)   // 获取机器人最近状态
}

func testFunc1(ctx iris.Context) {
	return
}
