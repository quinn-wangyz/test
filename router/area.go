package router

import (
	"fmt"
	"io/ioutil"
	"robot-api/libs/cache"
	"robot-api/libs/config"
	"robot-api/libs/httpUtil"

	"github.com/kataras/iris/v12"
)

func AreaRoute(route iris.Party) {
	route.Post("/list", listAreaHandler)                       // 获取区域列表
	route.Post("/add", areaAddHandler)                         // 创建区域基本信息
	route.Get("/get/{areaId:string}", getAreaDetailHandler)    // 获取区域详情
	route.Put("/update", updateAreaHandler)                    // 更新区域基本信息
	route.Delete("/delete/{areaId:string}", removeAreaHandler) // 删除区域

}

// 获取区域列表
func listAreaHandler(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	resp := httpUtil.Post(fmt.Sprintf("%s%s", cache.Get(config.Conf.Self.AreaServiceName), "/area/list"), data, "application/json")
	ctx.Write([]byte(resp))
}

// 创建区域
func areaAddHandler(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	resp := httpUtil.Put(fmt.Sprintf("%s%s", cache.Get(config.Conf.Self.AreaServiceName), "/area/"), data, "application/json")
	ctx.Write([]byte(resp))
}

// 获取区域详情
func getAreaDetailHandler(ctx iris.Context) {
	areaId := ctx.Params().Get("areaId")
	resp := httpUtil.Get(fmt.Sprintf("%s%s%s", cache.Get(config.Conf.Self.AreaServiceName), "/area/", areaId))
	ctx.Write([]byte(resp))
}

// 更新区域基本信息
func updateAreaHandler(ctx iris.Context) {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	data := jsonUnmarshal(body)
	areaId, _ := data["id"].(string)
	resp := httpUtil.Put(fmt.Sprintf("%s%s%s", cache.Get(config.Conf.Self.AreaServiceName), "/area/", areaId), data, "application/json")
	ctx.Write([]byte(resp))
}

// 删除区域
func removeAreaHandler(ctx iris.Context) {
	areaId := ctx.Params().Get("areaId")
	resp := httpUtil.Delete(fmt.Sprintf("%s%s%s", cache.Get(config.Conf.Self.MapServiceName), "/area/", areaId))
	ctx.Write([]byte(resp))
}
