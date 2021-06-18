package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"robot-api/libs/cache"
	"robot-api/libs/config"
	"robot-api/libs/httpUtil"

	"github.com/kataras/iris/v12"
)

func BuildingRouter(route iris.Party) {
	route.Post("/list", BuildingListFunc)                               //楼宇列表
	route.Post("/add", BuildingAddFunc)                                 //添加楼宇
	route.Get("/get/{buildingId:string}", BuildingGetFunc)              //获取楼宇详情
	route.Post("/update", BuildingUpdateFunc)                           //更新楼宇信息
	route.Get("/get/{buildingId:string}/floors", BuildingGetFloorsFunc) //获取楼宇的楼层列表
	route.Post("/update/floors", BuildingUpdateFloorsFunc)              //更新楼层信息
	route.Get("/delete/{buildingId:string}", BuildingDeleteFunc)        //删除楼宇
}

//楼宇列表
func BuildingListFunc(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	resp := httpUtil.Post(fmt.Sprintf("%s%s", cache.Get(config.Conf.Self.MapSerivceName), "/building/list"), data, "application/json")
	ctx.Write([]byte(resp))
}

//楼宇添加
func BuildingAddFunc(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	resp := httpUtil.Put(fmt.Sprintf("%s%s", cache.Get(config.Conf.Self.MapSerivceName), "/building/"), data, "application/json")
	ctx.Write([]byte(resp))
}

//获取楼宇详情
func BuildingGetFunc(ctx iris.Context) {
	buildingId := ctx.Params().Get("buildingId")
	resp := httpUtil.Get(fmt.Sprintf("%s%s%s", cache.Get(config.Conf.Self.MapSerivceName), "/building/", buildingId))
	ctx.Write([]byte(resp))
}

//更新楼宇信息
func BuildingUpdateFunc(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	buildingId, _ := data["id"].(string)
	resp := httpUtil.Put(fmt.Sprintf("%s%s%s", cache.Get(config.Conf.Self.MapSerivceName), "/building/", buildingId), data, "application/json")
	ctx.Write([]byte(resp))
}

//获取楼宇的楼层列表
func BuildingGetFloorsFunc(ctx iris.Context) {
	buildingId := ctx.Params().Get("buildingId")
	resp := httpUtil.Get(fmt.Sprintf("%s%s%s%s", cache.Get(config.Conf.Self.MapSerivceName), "/building/", buildingId, "/floors"))
	ctx.Write([]byte(resp))
}

//更新楼层信息
func BuildingUpdateFloorsFunc(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	buildingId, _ := data["buildingId"].(string)
	resp := httpUtil.Put(fmt.Sprintf("%s%s%s%s", cache.Get(config.Conf.Self.MapSerivceName), "/building/", buildingId, "/floors"), data, "application/json")
	ctx.Write([]byte(resp))
}

//删除楼宇
func BuildingDeleteFunc(ctx iris.Context) {
	buildingId := ctx.Params().Get("buildingId")
	resp := httpUtil.Delete(fmt.Sprintf("%s%s%s", cache.Get(config.Conf.Self.MapSerivceName), "/building/", buildingId))
	ctx.Write([]byte(resp))
}

//解析boby数据
func jsonUnmarshal(data []byte) map[string]interface{} {
	var i map[string]interface{}
	err := json.Unmarshal(data, &i)
	if err != nil {
		fmt.Println(err)
	}
	return i
}
