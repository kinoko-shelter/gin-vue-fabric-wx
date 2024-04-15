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

type CmapinfoApi struct {
}

var cmapinfoService = service.ServiceGroupApp.CmapServiceGroup.CmapinfoService

// CreateCmapinfo 创建地图
// @Tags Cmapinfo
// @Summary 创建地图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cmap.Cmapinfo true "创建地图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmapinfo/createCmapinfo [post]
func (cmapinfoApi *CmapinfoApi) CreateCmapinfo(c *gin.Context) {
	var cmapinfo cmap.Cmapinfo
	err := c.ShouldBindJSON(&cmapinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cmapinfo.CreatedBy = utils.GetUserID(c)

	if err := cmapinfoService.CreateCmapinfo(&cmapinfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmapinfo 删除地图
// @Tags Cmapinfo
// @Summary 删除地图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cmap.Cmapinfo true "删除地图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmapinfo/deleteCmapinfo [delete]
func (cmapinfoApi *CmapinfoApi) DeleteCmapinfo(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := cmapinfoService.DeleteCmapinfo(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmapinfoByIds 批量删除地图
// @Tags Cmapinfo
// @Summary 批量删除地图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmapinfo/deleteCmapinfoByIds [delete]
func (cmapinfoApi *CmapinfoApi) DeleteCmapinfoByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := cmapinfoService.DeleteCmapinfoByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmapinfo 更新地图
// @Tags Cmapinfo
// @Summary 更新地图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cmap.Cmapinfo true "更新地图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmapinfo/updateCmapinfo [put]
func (cmapinfoApi *CmapinfoApi) UpdateCmapinfo(c *gin.Context) {
	var cmapinfo cmap.Cmapinfo
	err := c.ShouldBindJSON(&cmapinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cmapinfo.UpdatedBy = utils.GetUserID(c)

	if err := cmapinfoService.UpdateCmapinfo(cmapinfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmapinfo 用id查询地图
// @Tags Cmapinfo
// @Summary 用id查询地图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmap.Cmapinfo true "用id查询地图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmapinfo/findCmapinfo [get]
func (cmapinfoApi *CmapinfoApi) FindCmapinfo(c *gin.Context) {
	ID := c.Query("ID")
	if recmapinfo, err := cmapinfoService.GetCmapinfo(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recmapinfo": recmapinfo}, c)
	}
}

//	分页获取地图列表
//
// @Tags Cmapinfo
// @Summary 分页获取地图列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cmapReq.CmapinfoSearch true "分页获取地图列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmapinfo/getCmapinfoList [get]
func (cmapinfoApi *CmapinfoApi) GetCmapinfoList(c *gin.Context) {
	var pageInfo cmapReq.CmapinfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cmapinfoService.GetCmapinfoInfoList(pageInfo); err != nil {
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
