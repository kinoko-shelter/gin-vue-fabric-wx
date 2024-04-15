package ChargingStatus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/ChargingStatus"
    ChargingStatusReq "github.com/flipped-aurora/gin-vue-admin/server/model/ChargingStatus/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type ChargingStatusApi struct {
}

var ChargingStatusInfoService = service.ServiceGroupApp.ChargingStatusServiceGroup.ChargingStatusService


// CreateChargingStatus 创建用户充电状态表
// @Tags ChargingStatus
// @Summary 创建用户充电状态表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ChargingStatus.ChargingStatus true "创建用户充电状态表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ChargingStatusInfo/createChargingStatus [post]
func (ChargingStatusInfoApi *ChargingStatusApi) CreateChargingStatus(c *gin.Context) {
	var ChargingStatusInfo ChargingStatus.ChargingStatus
	err := c.ShouldBindJSON(&ChargingStatusInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    ChargingStatusInfo.CreatedBy = utils.GetUserID(c)

	if err := ChargingStatusInfoService.CreateChargingStatus(&ChargingStatusInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteChargingStatus 删除用户充电状态表
// @Tags ChargingStatus
// @Summary 删除用户充电状态表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ChargingStatus.ChargingStatus true "删除用户充电状态表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ChargingStatusInfo/deleteChargingStatus [delete]
func (ChargingStatusInfoApi *ChargingStatusApi) DeleteChargingStatus(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := ChargingStatusInfoService.DeleteChargingStatus(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteChargingStatusByIds 批量删除用户充电状态表
// @Tags ChargingStatus
// @Summary 批量删除用户充电状态表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ChargingStatusInfo/deleteChargingStatusByIds [delete]
func (ChargingStatusInfoApi *ChargingStatusApi) DeleteChargingStatusByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := ChargingStatusInfoService.DeleteChargingStatusByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateChargingStatus 更新用户充电状态表
// @Tags ChargingStatus
// @Summary 更新用户充电状态表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ChargingStatus.ChargingStatus true "更新用户充电状态表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ChargingStatusInfo/updateChargingStatus [put]
func (ChargingStatusInfoApi *ChargingStatusApi) UpdateChargingStatus(c *gin.Context) {
	var ChargingStatusInfo ChargingStatus.ChargingStatus
	err := c.ShouldBindJSON(&ChargingStatusInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    ChargingStatusInfo.UpdatedBy = utils.GetUserID(c)

	if err := ChargingStatusInfoService.UpdateChargingStatus(ChargingStatusInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindChargingStatus 用id查询用户充电状态表
// @Tags ChargingStatus
// @Summary 用id查询用户充电状态表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ChargingStatus.ChargingStatus true "用id查询用户充电状态表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ChargingStatusInfo/findChargingStatus [get]
func (ChargingStatusInfoApi *ChargingStatusApi) FindChargingStatus(c *gin.Context) {
	ID := c.Query("ID")
	if reChargingStatusInfo, err := ChargingStatusInfoService.GetChargingStatus(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reChargingStatusInfo": reChargingStatusInfo}, c)
	}
}

// GetChargingStatusList 分页获取用户充电状态表列表
// @Tags ChargingStatus
// @Summary 分页获取用户充电状态表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ChargingStatusReq.ChargingStatusSearch true "分页获取用户充电状态表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ChargingStatusInfo/getChargingStatusList [get]
func (ChargingStatusInfoApi *ChargingStatusApi) GetChargingStatusList(c *gin.Context) {
	var pageInfo ChargingStatusReq.ChargingStatusSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ChargingStatusInfoService.GetChargingStatusInfoList(pageInfo); err != nil {
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
