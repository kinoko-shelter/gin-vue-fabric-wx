import service from '@/utils/request'

// @Tags Carmap
// @Summary 创建carmap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Carmap true "创建carmap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /carmap/createCarmap [post]
export const createCarmap = (data) => {
  return service({
    url: '/carmap/createCarmap',
    method: 'post',
    data
  })
}

// @Tags Carmap
// @Summary 删除carmap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Carmap true "删除carmap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /carmap/deleteCarmap [delete]
export const deleteCarmap = (params) => {
  return service({
    url: '/carmap/deleteCarmap',
    method: 'delete',
    params
  })
}

// @Tags Carmap
// @Summary 批量删除carmap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除carmap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /carmap/deleteCarmap [delete]
export const deleteCarmapByIds = (params) => {
  return service({
    url: '/carmap/deleteCarmapByIds',
    method: 'delete',
    params
  })
}

// @Tags Carmap
// @Summary 更新carmap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Carmap true "更新carmap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /carmap/updateCarmap [put]
export const updateCarmap = (data) => {
  return service({
    url: '/carmap/updateCarmap',
    method: 'put',
    data
  })
}

// @Tags Carmap
// @Summary 用id查询carmap
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Carmap true "用id查询carmap"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /carmap/findCarmap [get]
export const findCarmap = (params) => {
  return service({
    url: '/carmap/findCarmap',
    method: 'get',
    params
  })
}

// @Tags Carmap
// @Summary 分页获取carmap列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取carmap列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /carmap/getCarmapList [get]
export const getCarmapList = (params) => {
  return service({
    url: '/carmap/getCarmapList',
    method: 'get',
    params
  })
}
