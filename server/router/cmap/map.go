package cmap

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CarmapRouter struct {
}

// InitCarmapRouter 初始化 carmap 路由信息
func (s *CarmapRouter) InitCarmapRouter(Router *gin.RouterGroup) {
	carmapRouter := Router.Group("carmap").Use(middleware.OperationRecord())
	carmapRouterWithoutRecord := Router.Group("carmap")
	var carmapApi = v1.ApiGroupApp.CmapApiGroup.CarmapApi
	{
		carmapRouter.POST("createCarmap", carmapApi.CreateCarmap)   // 新建carmap
		carmapRouter.DELETE("deleteCarmap", carmapApi.DeleteCarmap) // 删除carmap
		carmapRouter.DELETE("deleteCarmapByIds", carmapApi.DeleteCarmapByIds) // 批量删除carmap
		carmapRouter.PUT("updateCarmap", carmapApi.UpdateCarmap)    // 更新carmap
	}
	{
		carmapRouterWithoutRecord.GET("findCarmap", carmapApi.FindCarmap)        // 根据ID获取carmap
		carmapRouterWithoutRecord.GET("getCarmapList", carmapApi.GetCarmapList)  // 获取carmap列表
	}
}
