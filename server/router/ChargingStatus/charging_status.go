package ChargingStatus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ChargingStatusRouter struct {
}

// InitChargingStatusRouter 初始化 用户充电状态表 路由信息
func (s *ChargingStatusRouter) InitChargingStatusRouter(Router *gin.RouterGroup) {
	ChargingStatusInfoRouter := Router.Group("ChargingStatusInfo").Use(middleware.OperationRecord())
	ChargingStatusInfoRouterWithoutRecord := Router.Group("ChargingStatusInfo")
	var ChargingStatusInfoApi = v1.ApiGroupApp.ChargingStatusApiGroup.ChargingStatusApi
	{
		ChargingStatusInfoRouter.POST("createChargingStatus", ChargingStatusInfoApi.CreateChargingStatus)   // 新建用户充电状态表
		ChargingStatusInfoRouter.DELETE("deleteChargingStatus", ChargingStatusInfoApi.DeleteChargingStatus) // 删除用户充电状态表
		ChargingStatusInfoRouter.DELETE("deleteChargingStatusByIds", ChargingStatusInfoApi.DeleteChargingStatusByIds) // 批量删除用户充电状态表
		ChargingStatusInfoRouter.PUT("updateChargingStatus", ChargingStatusInfoApi.UpdateChargingStatus)    // 更新用户充电状态表
	}
	{
		ChargingStatusInfoRouterWithoutRecord.GET("findChargingStatus", ChargingStatusInfoApi.FindChargingStatus)        // 根据ID获取用户充电状态表
		ChargingStatusInfoRouterWithoutRecord.GET("getChargingStatusList", ChargingStatusInfoApi.GetChargingStatusList)  // 获取用户充电状态表列表
	}
}
