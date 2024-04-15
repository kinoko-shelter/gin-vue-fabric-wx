package Vehicles

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VehiclesinfoRouter struct {
}

// InitVehiclesinfoRouter 初始化 车辆信息表 路由信息
func (s *VehiclesinfoRouter) InitVehiclesinfoRouter(Router *gin.RouterGroup) {
	VehiclesinfoCarRouter := Router.Group("VehiclesinfoCar").Use(middleware.OperationRecord())
	VehiclesinfoCarRouterWithoutRecord := Router.Group("VehiclesinfoCar")
	var VehiclesinfoCarApi = v1.ApiGroupApp.VehiclesApiGroup.VehiclesinfoApi
	{
		VehiclesinfoCarRouter.POST("createVehiclesinfo", VehiclesinfoCarApi.CreateVehiclesinfo)   // 新建车辆信息表
		VehiclesinfoCarRouter.DELETE("deleteVehiclesinfo", VehiclesinfoCarApi.DeleteVehiclesinfo) // 删除车辆信息表
		VehiclesinfoCarRouter.DELETE("deleteVehiclesinfoByIds", VehiclesinfoCarApi.DeleteVehiclesinfoByIds) // 批量删除车辆信息表
		VehiclesinfoCarRouter.PUT("updateVehiclesinfo", VehiclesinfoCarApi.UpdateVehiclesinfo)    // 更新车辆信息表
	}
	{
		VehiclesinfoCarRouterWithoutRecord.GET("findVehiclesinfo", VehiclesinfoCarApi.FindVehiclesinfo)        // 根据ID获取车辆信息表
		VehiclesinfoCarRouterWithoutRecord.GET("getVehiclesinfoList", VehiclesinfoCarApi.GetVehiclesinfoList)  // 获取车辆信息表列表
	}
}
