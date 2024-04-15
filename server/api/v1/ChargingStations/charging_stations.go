package ChargingStations

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/ChargingStations"
    ChargingStationsReq "github.com/flipped-aurora/gin-vue-admin/server/model/ChargingStations/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type ChargingStationsApi struct {
}

var ChargingStationsInfoService = service.ServiceGroupApp.ChargingStationsServiceGroup.ChargingStationsService


// CreateChargingStations 创建充电站信息表
// @Tags ChargingStations
// @Summary 创建充电站信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ChargingStations.ChargingStations true "创建充电站信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ChargingStationsInfo/createChargingStations [post]
func (ChargingStationsInfoApi *ChargingStationsApi) CreateChargingStations(c *gin.Context) {
	var ChargingStationsInfo ChargingStations.ChargingStations
	err := c.ShouldBindJSON(&ChargingStationsInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    ChargingStationsInfo.CreatedBy = utils.GetUserID(c)

	if err := ChargingStationsInfoService.CreateChargingStations(&ChargingStationsInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteChargingStations 删除充电站信息表
// @Tags ChargingStations
// @Summary 删除充电站信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ChargingStations.ChargingStations true "删除充电站信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ChargingStationsInfo/deleteChargingStations [delete]
func (ChargingStationsInfoApi *ChargingStationsApi) DeleteChargingStations(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := ChargingStationsInfoService.DeleteChargingStations(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteChargingStationsByIds 批量删除充电站信息表
// @Tags ChargingStations
// @Summary 批量删除充电站信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ChargingStationsInfo/deleteChargingStationsByIds [delete]
func (ChargingStationsInfoApi *ChargingStationsApi) DeleteChargingStationsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := ChargingStationsInfoService.DeleteChargingStationsByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateChargingStations 更新充电站信息表
// @Tags ChargingStations
// @Summary 更新充电站信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ChargingStations.ChargingStations true "更新充电站信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ChargingStationsInfo/updateChargingStations [put]
func (ChargingStationsInfoApi *ChargingStationsApi) UpdateChargingStations(c *gin.Context) {
	var ChargingStationsInfo ChargingStations.ChargingStations
	err := c.ShouldBindJSON(&ChargingStationsInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    ChargingStationsInfo.UpdatedBy = utils.GetUserID(c)

	if err := ChargingStationsInfoService.UpdateChargingStations(ChargingStationsInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindChargingStations 用id查询充电站信息表
// @Tags ChargingStations
// @Summary 用id查询充电站信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ChargingStations.ChargingStations true "用id查询充电站信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ChargingStationsInfo/findChargingStations [get]
func (ChargingStationsInfoApi *ChargingStationsApi) FindChargingStations(c *gin.Context) {
	ID := c.Query("ID")
	if reChargingStationsInfo, err := ChargingStationsInfoService.GetChargingStations(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reChargingStationsInfo": reChargingStationsInfo}, c)
	}
}

// GetChargingStationsList 分页获取充电站信息表列表
// @Tags ChargingStations
// @Summary 分页获取充电站信息表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ChargingStationsReq.ChargingStationsSearch true "分页获取充电站信息表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ChargingStationsInfo/getChargingStationsList [get]
func (ChargingStationsInfoApi *ChargingStationsApi) GetChargingStationsList(c *gin.Context) {
	var pageInfo ChargingStationsReq.ChargingStationsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ChargingStationsInfoService.GetChargingStationsInfoList(pageInfo); err != nil {
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
