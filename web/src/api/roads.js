import service from '@/utils/request'

// @Tags Roads
// @Summary 创建道路信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Roads true "创建道路信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /RoadsInfo/createRoads [post]
export const createRoads = (data) => {
  return service({
    url: '/RoadsInfo/createRoads',
    method: 'post',
    data
  })
}

// @Tags Roads
// @Summary 删除道路信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Roads true "删除道路信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /RoadsInfo/deleteRoads [delete]
export const deleteRoads = (params) => {
  return service({
    url: '/RoadsInfo/deleteRoads',
    method: 'delete',
    params
  })
}

// @Tags Roads
// @Summary 批量删除道路信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除道路信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /RoadsInfo/deleteRoads [delete]
export const deleteRoadsByIds = (params) => {
  return service({
    url: '/RoadsInfo/deleteRoadsByIds',
    method: 'delete',
    params
  })
}

// @Tags Roads
// @Summary 更新道路信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Roads true "更新道路信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /RoadsInfo/updateRoads [put]
export const updateRoads = (data) => {
  return service({
    url: '/RoadsInfo/updateRoads',
    method: 'put',
    data
  })
}

// @Tags Roads
// @Summary 用id查询道路信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Roads true "用id查询道路信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /RoadsInfo/findRoads [get]
export const findRoads = (params) => {
  return service({
    url: '/RoadsInfo/findRoads',
    method: 'get',
    params
  })
}

// @Tags Roads
// @Summary 分页获取道路信息表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取道路信息表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /RoadsInfo/getRoadsList [get]
export const getRoadsList = (params) => {
  return service({
    url: '/RoadsInfo/getRoadsList',
    method: 'get',
    params
  })
}
