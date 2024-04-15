package Roads

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RoadsRouter struct {
}

// InitRoadsRouter 初始化 道路信息表 路由信息
func (s *RoadsRouter) InitRoadsRouter(Router *gin.RouterGroup) {
	RoadsInfoRouter := Router.Group("RoadsInfo").Use(middleware.OperationRecord())
	RoadsInfoRouterWithoutRecord := Router.Group("RoadsInfo")
	var RoadsInfoApi = v1.ApiGroupApp.RoadsApiGroup.RoadsApi
	{
		RoadsInfoRouter.POST("createRoads", RoadsInfoApi.CreateRoads)   // 新建道路信息表
		RoadsInfoRouter.DELETE("deleteRoads", RoadsInfoApi.DeleteRoads) // 删除道路信息表
		RoadsInfoRouter.DELETE("deleteRoadsByIds", RoadsInfoApi.DeleteRoadsByIds) // 批量删除道路信息表
		RoadsInfoRouter.PUT("updateRoads", RoadsInfoApi.UpdateRoads)    // 更新道路信息表
	}
	{
		RoadsInfoRouterWithoutRecord.GET("findRoads", RoadsInfoApi.FindRoads)        // 根据ID获取道路信息表
		RoadsInfoRouterWithoutRecord.GET("getRoadsList", RoadsInfoApi.GetRoadsList)  // 获取道路信息表列表
	}
}
