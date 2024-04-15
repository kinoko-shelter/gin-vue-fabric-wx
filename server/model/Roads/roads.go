// 自动生成模板Roads
package Roads

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 道路信息表 结构体  Roads
type Roads struct {
 global.GVA_MODEL 
      Road_id  *int `json:"road_id" form:"road_id" gorm:"column:road_id;comment:道路id;"`  //道路id 
      Road_len  string `json:"road_len" form:"road_len" gorm:"column:road_len;comment:路长;"`  //路长 
      Road_limit  string `json:"road_limit" form:"road_limit" gorm:"column:road_limit;comment:道路限速;"`  //道路限速 
      R_startx  string `json:"r_startx" form:"r_startx" gorm:"column:r_startx;comment:路段起点x;"`  //路段起点x 
      R_starty  string `json:"r_starty" form:"r_starty" gorm:"column:r_starty;comment:路段起点y;"`  //路段起点y 
      R_endx  string `json:"r_endx" form:"r_endx" gorm:"column:r_endx;comment:路段终点;"`  //路段终点x 
      R_endy  string `json:"r_endy" form:"r_endy" gorm:"column:r_endy;comment:路段终点;"`  //路段终点y 
      R_JamIndex  string `json:"r_JamIndex" form:"r_JamIndex" gorm:"column:r__jam_index;comment:路段拥挤指数;"`  //路段拥挤指数 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 道路信息表 Roads自定义表名 Roads
func (Roads) TableName() string {
  return "Roads"
}

