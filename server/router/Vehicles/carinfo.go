package Vehicles

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CarinfoRouter struct {
}

// InitCarinfoRouter 初始化 车辆信息表 路由信息
func (s *CarinfoRouter) InitCarinfoRouter(Router *gin.RouterGroup) {
	carinfoRouter := Router.Group("carinfo").Use(middleware.OperationRecord())
	carinfoRouterWithoutRecord := Router.Group("carinfo")
	var carinfoApi = v1.ApiGroupApp.VehiclesApiGroup.CarinfoApi
	{
		carinfoRouter.POST("createCarinfo", carinfoApi.CreateCarinfo)   // 新建车辆信息表
		carinfoRouter.DELETE("deleteCarinfo", carinfoApi.DeleteCarinfo) // 删除车辆信息表
		carinfoRouter.DELETE("deleteCarinfoByIds", carinfoApi.DeleteCarinfoByIds) // 批量删除车辆信息表
		carinfoRouter.PUT("updateCarinfo", carinfoApi.UpdateCarinfo)    // 更新车辆信息表
	}
	{
		carinfoRouterWithoutRecord.GET("findCarinfo", carinfoApi.FindCarinfo)        // 根据ID获取车辆信息表
		carinfoRouterWithoutRecord.GET("getCarinfoList", carinfoApi.GetCarinfoList)  // 获取车辆信息表列表
	}
}
