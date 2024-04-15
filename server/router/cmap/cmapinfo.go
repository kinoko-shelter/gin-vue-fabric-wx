package cmap

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CmapinfoRouter struct {
}

// InitCmapinfoRouter 初始化 地图 路由信息
func (s *CmapinfoRouter) InitCmapinfoRouter(Router *gin.RouterGroup) {
	cmapinfoRouter := Router.Group("cmapinfo").Use(middleware.OperationRecord())
	cmapinfoRouterWithoutRecord := Router.Group("cmapinfo")
	var cmapinfoApi = v1.ApiGroupApp.CmapApiGroup.CmapinfoApi
	{
		cmapinfoRouter.POST("createCmapinfo", cmapinfoApi.CreateCmapinfo)   // 新建地图
		cmapinfoRouter.DELETE("deleteCmapinfo", cmapinfoApi.DeleteCmapinfo) // 删除地图
		cmapinfoRouter.DELETE("deleteCmapinfoByIds", cmapinfoApi.DeleteCmapinfoByIds) // 批量删除地图
		cmapinfoRouter.PUT("updateCmapinfo", cmapinfoApi.UpdateCmapinfo)    // 更新地图
	}
	{
		cmapinfoRouterWithoutRecord.GET("findCmapinfo", cmapinfoApi.FindCmapinfo)        // 根据ID获取地图
		cmapinfoRouterWithoutRecord.GET("getCmapinfoList", cmapinfoApi.GetCmapinfoList)  // 获取地图列表
	}
}
