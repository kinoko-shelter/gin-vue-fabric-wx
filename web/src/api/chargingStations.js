import service from '@/utils/request'

// @Tags ChargingStations
// @Summary 创建充电站信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ChargingStations true "创建充电站信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ChargingStationsInfo/createChargingStations [post]
export const createChargingStations = (data) => {
  return service({
    url: '/ChargingStationsInfo/createChargingStations',
    method: 'post',
    data
  })
}

// @Tags ChargingStations
// @Summary 删除充电站信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ChargingStations true "删除充电站信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ChargingStationsInfo/deleteChargingStations [delete]
export const deleteChargingStations = (params) => {
  return service({
    url: '/ChargingStationsInfo/deleteChargingStations',
    method: 'delete',
    params
  })
}

// @Tags ChargingStations
// @Summary 批量删除充电站信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除充电站信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ChargingStationsInfo/deleteChargingStations [delete]
export const deleteChargingStationsByIds = (params) => {
  return service({
    url: '/ChargingStationsInfo/deleteChargingStationsByIds',
    method: 'delete',
    params
  })
}

// @Tags ChargingStations
// @Summary 更新充电站信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ChargingStations true "更新充电站信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ChargingStationsInfo/updateChargingStations [put]
export const updateChargingStations = (data) => {
  return service({
    url: '/ChargingStationsInfo/updateChargingStations',
    method: 'put',
    data
  })
}

// @Tags ChargingStations
// @Summary 用id查询充电站信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ChargingStations true "用id查询充电站信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ChargingStationsInfo/findChargingStations [get]
export const findChargingStations = (params) => {
  return service({
    url: '/ChargingStationsInfo/findChargingStations',
    method: 'get',
    params
  })
}

// @Tags ChargingStations
// @Summary 分页获取充电站信息表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取充电站信息表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ChargingStationsInfo/getChargingStationsList [get]
export const getChargingStationsList = (params) => {
  return service({
    url: '/ChargingStationsInfo/getChargingStationsList',
    method: 'get',
    params
  })
}
