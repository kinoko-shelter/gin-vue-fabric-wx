import service from '@/utils/request'

// @Tags ChargingStatus
// @Summary 创建用户充电状态表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ChargingStatus true "创建用户充电状态表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ChargingStatusInfo/createChargingStatus [post]
export const createChargingStatus = (data) => {
  return service({
    url: '/ChargingStatusInfo/createChargingStatus',
    method: 'post',
    data
  })
}

// @Tags ChargingStatus
// @Summary 删除用户充电状态表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ChargingStatus true "删除用户充电状态表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ChargingStatusInfo/deleteChargingStatus [delete]
export const deleteChargingStatus = (params) => {
  return service({
    url: '/ChargingStatusInfo/deleteChargingStatus',
    method: 'delete',
    params
  })
}

// @Tags ChargingStatus
// @Summary 批量删除用户充电状态表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户充电状态表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ChargingStatusInfo/deleteChargingStatus [delete]
export const deleteChargingStatusByIds = (params) => {
  return service({
    url: '/ChargingStatusInfo/deleteChargingStatusByIds',
    method: 'delete',
    params
  })
}

// @Tags ChargingStatus
// @Summary 更新用户充电状态表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ChargingStatus true "更新用户充电状态表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ChargingStatusInfo/updateChargingStatus [put]
export const updateChargingStatus = (data) => {
  return service({
    url: '/ChargingStatusInfo/updateChargingStatus',
    method: 'put',
    data
  })
}

// @Tags ChargingStatus
// @Summary 用id查询用户充电状态表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ChargingStatus true "用id查询用户充电状态表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ChargingStatusInfo/findChargingStatus [get]
export const findChargingStatus = (params) => {
  return service({
    url: '/ChargingStatusInfo/findChargingStatus',
    method: 'get',
    params
  })
}

// @Tags ChargingStatus
// @Summary 分页获取用户充电状态表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户充电状态表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ChargingStatusInfo/getChargingStatusList [get]
export const getChargingStatusList = (params) => {
  return service({
    url: '/ChargingStatusInfo/getChargingStatusList',
    method: 'get',
    params
  })
}
