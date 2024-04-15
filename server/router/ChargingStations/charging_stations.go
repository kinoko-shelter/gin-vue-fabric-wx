package ChargingStations

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ChargingStationsRouter struct {
}

// InitChargingStationsRouter 初始化 充电站信息表 路由信息
func (s *ChargingStationsRouter) InitChargingStationsRouter(Router *gin.RouterGroup) {
	ChargingStationsInfoRouter := Router.Group("ChargingStationsInfo").Use(middleware.OperationRecord())
	ChargingStationsInfoRouterWithoutRecord := Router.Group("ChargingStationsInfo")
	var ChargingStationsInfoApi = v1.ApiGroupApp.ChargingStationsApiGroup.ChargingStationsApi
	{
		ChargingStationsInfoRouter.POST("createChargingStations", ChargingStationsInfoApi.CreateChargingStations)   // 新建充电站信息表
		ChargingStationsInfoRouter.DELETE("deleteChargingStations", ChargingStationsInfoApi.DeleteChargingStations) // 删除充电站信息表
		ChargingStationsInfoRouter.DELETE("deleteChargingStationsByIds", ChargingStationsInfoApi.DeleteChargingStationsByIds) // 批量删除充电站信息表
		ChargingStationsInfoRouter.PUT("updateChargingStations", ChargingStationsInfoApi.UpdateChargingStations)    // 更新充电站信息表
	}
	{
		ChargingStationsInfoRouterWithoutRecord.GET("findChargingStations", ChargingStationsInfoApi.FindChargingStations)        // 根据ID获取充电站信息表
		ChargingStationsInfoRouterWithoutRecord.GET("getChargingStationsList", ChargingStationsInfoApi.GetChargingStationsList)  // 获取充电站信息表列表
	}
}
