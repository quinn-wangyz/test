package router

import (
	"fmt"
	"io/ioutil"
	"robot-api/libs/cache"
	"robot-api/libs/config"
	"robot-api/libs/httpUtil"

	"github.com/kataras/iris/v12"
)

func BuildingRoute(route iris.Party) {
	route.Post("/list", listBuildingHandler)                               //楼宇列表
	route.Post("/add", addBuildingHandler)                                 //添加楼宇
	route.Get("/get/{buildingId:string}", getBuildingHandler)              //获取楼宇详情
	route.Post("/update", updateBuildingHandler)                           //更新楼宇信息
	route.Get("/get/{buildingId:string}/floors", getBuildingFloorsHandler) //获取楼宇的楼层列表
	route.Post("/update/floors", updateBuildingFloorsHandler)              //更新楼层信息
	route.Get("/delete/{buildingId:string}", deleteBuildingHandler)        //删除楼宇
}

//楼宇列表
func listBuildingHandler(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	resp := httpUtil.Post(fmt.Sprintf("%s%s", cache.Get(config.Conf.DependentService.MapService.ServiceName), "/building/list"), data, "application/json")
	ctx.Write([]byte(resp))
}

//楼宇添加
func addBuildingHandler(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	resp := httpUtil.Put(fmt.Sprintf("%s%s", cache.Get(config.Conf.DependentService.MapService.ServiceName), "/building/"), data, "application/json")
	ctx.Write([]byte(resp))
}

//获取楼宇详情
func getBuildingHandler(ctx iris.Context) {
	buildingId := ctx.Params().Get("buildingId")
	resp := httpUtil.Get(fmt.Sprintf("%s%s%s", cache.Get(config.Conf.DependentService.MapService.ServiceName), "/building/", buildingId))
	ctx.Write([]byte(resp))
}

//更新楼宇信息
func updateBuildingHandler(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	buildingId, _ := data["id"].(string)
	resp := httpUtil.Put(fmt.Sprintf("%s%s%s", cache.Get(config.Conf.DependentService.MapService.ServiceName), "/building/", buildingId), data, "application/json")
	ctx.Write([]byte(resp))
}

//获取楼宇的楼层列表
func getBuildingFloorsHandler(ctx iris.Context) {
	buildingId := ctx.Params().Get("buildingId")
	resp := httpUtil.Get(fmt.Sprintf("%s%s%s%s", cache.Get(config.Conf.DependentService.MapService.ServiceName), "/building/", buildingId, "/floors"))
	ctx.Write([]byte(resp))
}

//更新楼层信息
func updateBuildingFloorsHandler(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	buildingId, _ := data["buildingId"].(string)
	resp := httpUtil.Put(fmt.Sprintf("%s%s%s%s", cache.Get(config.Conf.DependentService.MapService.ServiceName), "/building/", buildingId, "/floors"), data, "application/json")
	ctx.Write([]byte(resp))
}

//删除楼宇
func deleteBuildingHandler(ctx iris.Context) {
	buildingId := ctx.Params().Get("buildingId")
	resp := httpUtil.Delete(fmt.Sprintf("%s%s%s", cache.Get(config.Conf.DependentService.MapService.ServiceName), "/building/", buildingId))
	ctx.Write([]byte(resp))
}
