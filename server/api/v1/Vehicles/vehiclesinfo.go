package Vehicles

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Vehicles"
    VehiclesReq "github.com/flipped-aurora/gin-vue-admin/server/model/Vehicles/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type VehiclesinfoApi struct {
}

var VehiclesinfoCarService = service.ServiceGroupApp.VehiclesServiceGroup.VehiclesinfoService


// CreateVehiclesinfo 创建车辆信息表
// @Tags Vehiclesinfo
// @Summary 创建车辆信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Vehicles.Vehiclesinfo true "创建车辆信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /VehiclesinfoCar/createVehiclesinfo [post]
func (VehiclesinfoCarApi *VehiclesinfoApi) CreateVehiclesinfo(c *gin.Context) {
	var VehiclesinfoCar Vehicles.Vehiclesinfo
	err := c.ShouldBindJSON(&VehiclesinfoCar)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    VehiclesinfoCar.CreatedBy = utils.GetUserID(c)

	if err := VehiclesinfoCarService.CreateVehiclesinfo(&VehiclesinfoCar); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteVehiclesinfo 删除车辆信息表
// @Tags Vehiclesinfo
// @Summary 删除车辆信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Vehicles.Vehiclesinfo true "删除车辆信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /VehiclesinfoCar/deleteVehiclesinfo [delete]
func (VehiclesinfoCarApi *VehiclesinfoApi) DeleteVehiclesinfo(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := VehiclesinfoCarService.DeleteVehiclesinfo(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteVehiclesinfoByIds 批量删除车辆信息表
// @Tags Vehiclesinfo
// @Summary 批量删除车辆信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /VehiclesinfoCar/deleteVehiclesinfoByIds [delete]
func (VehiclesinfoCarApi *VehiclesinfoApi) DeleteVehiclesinfoByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := VehiclesinfoCarService.DeleteVehiclesinfoByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateVehiclesinfo 更新车辆信息表
// @Tags Vehiclesinfo
// @Summary 更新车辆信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Vehicles.Vehiclesinfo true "更新车辆信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /VehiclesinfoCar/updateVehiclesinfo [put]
func (VehiclesinfoCarApi *VehiclesinfoApi) UpdateVehiclesinfo(c *gin.Context) {
	var VehiclesinfoCar Vehicles.Vehiclesinfo
	err := c.ShouldBindJSON(&VehiclesinfoCar)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    VehiclesinfoCar.UpdatedBy = utils.GetUserID(c)

	if err := VehiclesinfoCarService.UpdateVehiclesinfo(VehiclesinfoCar); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindVehiclesinfo 用id查询车辆信息表
// @Tags Vehiclesinfo
// @Summary 用id查询车辆信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Vehicles.Vehiclesinfo true "用id查询车辆信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /VehiclesinfoCar/findVehiclesinfo [get]
func (VehiclesinfoCarApi *VehiclesinfoApi) FindVehiclesinfo(c *gin.Context) {
	ID := c.Query("ID")
	if reVehiclesinfoCar, err := VehiclesinfoCarService.GetVehiclesinfo(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reVehiclesinfoCar": reVehiclesinfoCar}, c)
	}
}

// GetVehiclesinfoList 分页获取车辆信息表列表
// @Tags Vehiclesinfo
// @Summary 分页获取车辆信息表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query VehiclesReq.VehiclesinfoSearch true "分页获取车辆信息表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /VehiclesinfoCar/getVehiclesinfoList [get]
func (VehiclesinfoCarApi *VehiclesinfoApi) GetVehiclesinfoList(c *gin.Context) {
	var pageInfo VehiclesReq.VehiclesinfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := VehiclesinfoCarService.GetVehiclesinfoInfoList(pageInfo); err != nil {
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
