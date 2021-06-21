package router

import (
	"encoding/json"
	"fmt"

	"github.com/kataras/iris/v12"
)

func Router(route iris.Party) {
	route.PartyFunc("/robot", RobotRoute)
	route.PartyFunc("/elevator", ElevatorRoute)
	route.PartyFunc("/building", BuildingRoute)
	route.PartyFunc("/business", BusinessRoute)
	route.PartyFunc("/area", AreaRoute)
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
