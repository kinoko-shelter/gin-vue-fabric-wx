package cmap

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmap"
	cmapReq "github.com/flipped-aurora/gin-vue-admin/server/model/cmap/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CarmapApi struct {
}

var carmapService = service.ServiceGroupApp.CmapServiceGroup.CarmapService

// CreateCarmap 创建carmap
// @Tags Carmap
// @Summary 创建carmap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cmap.Carmap true "创建carmap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /carmap/createCarmap [post]
func (carmapApi *CarmapApi) CreateCarmap(c *gin.Context) {
	var carmap cmap.Carmap
	err := c.ShouldBindJSON(&carmap)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	carmap.CreatedBy = utils.GetUserID(c)

	if err := carmapService.CreateCarmap(&carmap); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCarmap 删除carmap
// @Tags Carmap
// @Summary 删除carmap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cmap.Carmap true "删除carmap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /carmap/deleteCarmap [delete]
func (carmapApi *CarmapApi) DeleteCarmap(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := carmapService.DeleteCarmap(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCarmapByIds 批量删除carmap
// @Tags Carmap
// @Summary 批量删除carmap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /carmap/deleteCarmapByIds [delete]
func (carmapApi *CarmapApi) DeleteCarmapByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := carmapService.DeleteCarmapByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCarmap 更新carmap
// @Tags Carmap
// @Summary 更新carmap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cmap.Carmap true "更新carmap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /carmap/updateCarmap [put]
func (carmapApi *CarmapApi) UpdateCarmap(c *gin.Context) {
	var carmap cmap.Carmap
	err := c.ShouldBindJSON(&carmap)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	carmap.UpdatedBy = utils.GetUserID(c)

	if err := carmapService.UpdateCarmap(carmap); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCarmap 用id查询carmap
// @Tags Carmap
// @Summary 用id查询carmap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmap.Carmap true "用id查询carmap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /carmap/findCarmap [get]
func (carmapApi *CarmapApi) FindCarmap(c *gin.Context) {
	ID := c.Query("ID")
	if recarmap, err := carmapService.GetCarmap(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recarmap": recarmap}, c)
	}
}

// GetCarmapList 分页获取carmap列表
// @Tags Carmap
// @Summary 分页获取carmap列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmapReq.CarmapSearch true "分页获取carmap列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /carmap/getCarmapList [get]
func (carmapApi *CarmapApi) GetCarmapList(c *gin.Context) {
	var pageInfo cmapReq.CarmapSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := carmapService.GetCarmapInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
