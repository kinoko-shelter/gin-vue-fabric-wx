package Vehicles

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Vehicles"
    VehiclesReq "github.com/flipped-aurora/gin-vue-admin/server/model/Vehicles/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type CarinfoApi struct {
}

var carinfoService = service.ServiceGroupApp.VehiclesServiceGroup.CarinfoService


// CreateCarinfo 创建车辆信息表
// @Tags Carinfo
// @Summary 创建车辆信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Vehicles.Carinfo true "创建车辆信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /carinfo/createCarinfo [post]
func (carinfoApi *CarinfoApi) CreateCarinfo(c *gin.Context) {
	var carinfo Vehicles.Carinfo
	err := c.ShouldBindJSON(&carinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := carinfoService.CreateCarinfo(&carinfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCarinfo 删除车辆信息表
// @Tags Carinfo
// @Summary 删除车辆信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Vehicles.Carinfo true "删除车辆信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /carinfo/deleteCarinfo [delete]
func (carinfoApi *CarinfoApi) DeleteCarinfo(c *gin.Context) {
	ID := c.Query("ID")
	if err := carinfoService.DeleteCarinfo(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCarinfoByIds 批量删除车辆信息表
// @Tags Carinfo
// @Summary 批量删除车辆信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /carinfo/deleteCarinfoByIds [delete]
func (carinfoApi *CarinfoApi) DeleteCarinfoByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := carinfoService.DeleteCarinfoByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCarinfo 更新车辆信息表
// @Tags Carinfo
// @Summary 更新车辆信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Vehicles.Carinfo true "更新车辆信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /carinfo/updateCarinfo [put]
func (carinfoApi *CarinfoApi) UpdateCarinfo(c *gin.Context) {
	var carinfo Vehicles.Carinfo
	err := c.ShouldBindJSON(&carinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := carinfoService.UpdateCarinfo(carinfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCarinfo 用id查询车辆信息表
// @Tags Carinfo
// @Summary 用id查询车辆信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Vehicles.Carinfo true "用id查询车辆信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /carinfo/findCarinfo [get]
func (carinfoApi *CarinfoApi) FindCarinfo(c *gin.Context) {
	ID := c.Query("ID")
	if recarinfo, err := carinfoService.GetCarinfo(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recarinfo": recarinfo}, c)
	}
}

// GetCarinfoList 分页获取车辆信息表列表
// @Tags Carinfo
// @Summary 分页获取车辆信息表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query VehiclesReq.CarinfoSearch true "分页获取车辆信息表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /carinfo/getCarinfoList [get]
func (carinfoApi *CarinfoApi) GetCarinfoList(c *gin.Context) {
	var pageInfo VehiclesReq.CarinfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := carinfoService.GetCarinfoInfoList(pageInfo); err != nil {
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
