import service from '@/utils/request'

// @Tags Cmapinfo
// @Summary 创建地图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cmapinfo true "创建地图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmapinfo/createCmapinfo [post]
export const createCmapinfo = (data) => {
  return service({
    url: '/cmapinfo/createCmapinfo',
    method: 'post',
    data
  })
}

// @Tags Cmapinfo
// @Summary 删除地图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cmapinfo true "删除地图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmapinfo/deleteCmapinfo [delete]
export const deleteCmapinfo = (params) => {
  return service({
    url: '/cmapinfo/deleteCmapinfo',
    method: 'delete',
    params
  })
}

// @Tags Cmapinfo
// @Summary 批量删除地图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除地图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmapinfo/deleteCmapinfo [delete]
export const deleteCmapinfoByIds = (params) => {
  return service({
    url: '/cmapinfo/deleteCmapinfoByIds',
    method: 'delete',
    params
  })
}

// @Tags Cmapinfo
// @Summary 更新地图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cmapinfo true "更新地图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmapinfo/updateCmapinfo [put]
export const updateCmapinfo = (data) => {
  return service({
    url: '/cmapinfo/updateCmapinfo',
    method: 'put',
    data
  })
}

// @Tags Cmapinfo
// @Summary 用id查询地图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Cmapinfo true "用id查询地图"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmapinfo/findCmapinfo [get]
export const findCmapinfo = (params) => {
  return service({
    url: '/cmapinfo/findCmapinfo',
    method: 'get',
    params
  })
}

// @Tags Cmapinfo
// @Summary 分页获取地图列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取地图列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmapinfo/getCmapinfoList [get]
export const getCmapinfoList = (params) => {
  return service({
    url: '/cmapinfo/getCmapinfoList',
    method: 'get',
    params
  })
}
