// 自动生成模板Carmap
package cmap

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// carmap 结构体  Carmap
type Carmap struct {
 global.GVA_MODEL 
      Location_x  string `json:"location_x" form:"location_x" gorm:"column:location_x;comment:;"`  //位置x 
      Location_y  string `json:"location_y" form:"location_y" gorm:"column:location_y;comment:;"`  //位置y 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName carmap Carmap自定义表名 carmap
func (Carmap) TableName() string {
  return "carmap"
}

