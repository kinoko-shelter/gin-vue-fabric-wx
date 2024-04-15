import service from '@/utils/request'

// @Tags Vehiclesinfo
// @Summary 创建车辆信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Vehiclesinfo true "创建车辆信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /VehiclesinfoCar/createVehiclesinfo [post]
export const createVehiclesinfo = (data) => {
  return service({
    url: '/VehiclesinfoCar/createVehiclesinfo',
    method: 'post',
    data
  })
}

// @Tags Vehiclesinfo
// @Summary 删除车辆信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Vehiclesinfo true "删除车辆信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /VehiclesinfoCar/deleteVehiclesinfo [delete]
export const deleteVehiclesinfo = (params) => {
  return service({
    url: '/VehiclesinfoCar/deleteVehiclesinfo',
    method: 'delete',
    params
  })
}

// @Tags Vehiclesinfo
// @Summary 批量删除车辆信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除车辆信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /VehiclesinfoCar/deleteVehiclesinfo [delete]
export const deleteVehiclesinfoByIds = (params) => {
  return service({
    url: '/VehiclesinfoCar/deleteVehiclesinfoByIds',
    method: 'delete',
    params
  })
}

// @Tags Vehiclesinfo
// @Summary 更新车辆信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Vehiclesinfo true "更新车辆信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /VehiclesinfoCar/updateVehiclesinfo [put]
export const updateVehiclesinfo = (data) => {
  return service({
    url: '/VehiclesinfoCar/updateVehiclesinfo',
    method: 'put',
    data
  })
}

// @Tags Vehiclesinfo
// @Summary 用id查询车辆信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Vehiclesinfo true "用id查询车辆信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /VehiclesinfoCar/findVehiclesinfo [get]
export const findVehiclesinfo = (params) => {
  return service({
    url: '/VehiclesinfoCar/findVehiclesinfo',
    method: 'get',
    params
  })
}

// @Tags Vehiclesinfo
// @Summary 分页获取车辆信息表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取车辆信息表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /VehiclesinfoCar/getVehiclesinfoList [get]
export const getVehiclesinfoList = (params) => {
  return service({
    url: '/VehiclesinfoCar/getVehiclesinfoList',
    method: 'get',
    params
  })
}
