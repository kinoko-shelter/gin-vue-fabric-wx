import service from '@/utils/request'

// @Tags Carinfo
// @Summary 创建车辆信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Carinfo true "创建车辆信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /carinfo/createCarinfo [post]
export const createCarinfo = (data) => {
  return service({
    url: '/carinfo/createCarinfo',
    method: 'post',
    data
  })
}

// @Tags Carinfo
// @Summary 删除车辆信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Carinfo true "删除车辆信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /carinfo/deleteCarinfo [delete]
export const deleteCarinfo = (params) => {
  return service({
    url: '/carinfo/deleteCarinfo',
    method: 'delete',
    params
  })
}

// @Tags Carinfo
// @Summary 批量删除车辆信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除车辆信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /carinfo/deleteCarinfo [delete]
export const deleteCarinfoByIds = (params) => {
  return service({
    url: '/carinfo/deleteCarinfoByIds',
    method: 'delete',
    params
  })
}

// @Tags Carinfo
// @Summary 更新车辆信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Carinfo true "更新车辆信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /carinfo/updateCarinfo [put]
export const updateCarinfo = (data) => {
  return service({
    url: '/carinfo/updateCarinfo',
    method: 'put',
    data
  })
}

// @Tags Carinfo
// @Summary 用id查询车辆信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Carinfo true "用id查询车辆信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /carinfo/findCarinfo [get]
export const findCarinfo = (params) => {
  return service({
    url: '/carinfo/findCarinfo',
    method: 'get',
    params
  })
}

// @Tags Carinfo
// @Summary 分页获取车辆信息表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取车辆信息表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /carinfo/getCarinfoList [get]
export const getCarinfoList = (params) => {
  return service({
    url: '/carinfo/getCarinfoList',
    method: 'get',
    params
  })
}
