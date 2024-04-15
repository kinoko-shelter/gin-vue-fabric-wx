package Roads

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Roads"
    RoadsReq "github.com/flipped-aurora/gin-vue-admin/server/model/Roads/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type RoadsApi struct {
}

var RoadsInfoService = service.ServiceGroupApp.RoadsServiceGroup.RoadsService


// CreateRoads 创建道路信息表
// @Tags Roads
// @Summary 创建道路信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Roads.Roads true "创建道路信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /RoadsInfo/createRoads [post]
func (RoadsInfoApi *RoadsApi) CreateRoads(c *gin.Context) {
	var RoadsInfo Roads.Roads
	err := c.ShouldBindJSON(&RoadsInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    RoadsInfo.CreatedBy = utils.GetUserID(c)

	if err := RoadsInfoService.CreateRoads(&RoadsInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRoads 删除道路信息表
// @Tags Roads
// @Summary 删除道路信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Roads.Roads true "删除道路信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /RoadsInfo/deleteRoads [delete]
func (RoadsInfoApi *RoadsApi) DeleteRoads(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := RoadsInfoService.DeleteRoads(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRoadsByIds 批量删除道路信息表
// @Tags Roads
// @Summary 批量删除道路信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /RoadsInfo/deleteRoadsByIds [delete]
func (RoadsInfoApi *RoadsApi) DeleteRoadsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := RoadsInfoService.DeleteRoadsByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRoads 更新道路信息表
// @Tags Roads
// @Summary 更新道路信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Roads.Roads true "更新道路信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /RoadsInfo/updateRoads [put]
func (RoadsInfoApi *RoadsApi) UpdateRoads(c *gin.Context) {
	var RoadsInfo Roads.Roads
	err := c.ShouldBindJSON(&RoadsInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    RoadsInfo.UpdatedBy = utils.GetUserID(c)

	if err := RoadsInfoService.UpdateRoads(RoadsInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRoads 用id查询道路信息表
// @Tags Roads
// @Summary 用id查询道路信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Roads.Roads true "用id查询道路信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /RoadsInfo/findRoads [get]
func (RoadsInfoApi *RoadsApi) FindRoads(c *gin.Context) {
	ID := c.Query("ID")
	if reRoadsInfo, err := RoadsInfoService.GetRoads(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reRoadsInfo": reRoadsInfo}, c)
	}
}

// GetRoadsList 分页获取道路信息表列表
// @Tags Roads
// @Summary 分页获取道路信息表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query RoadsReq.RoadsSearch true "分页获取道路信息表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /RoadsInfo/getRoadsList [get]
func (RoadsInfoApi *RoadsApi) GetRoadsList(c *gin.Context) {
	var pageInfo RoadsReq.RoadsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := RoadsInfoService.GetRoadsInfoList(pageInfo); err != nil {
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
