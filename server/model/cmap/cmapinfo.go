// 自动生成模板Cmapinfo
package cmap

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 地图 结构体  Cmapinfo
type Cmapinfo struct {
	global.GVA_MODEL
	Name       string `json:"name" form:"name" gorm:"column:name;comment:;"`
	Paths      string `json:"paths" form:"paths" gorm:"column:paths;comment:;type:text;"`
	Location_x string `json:"location_x" form:"location_x" gorm:"column:location_x;comment:;"` //位置x
	Location_y string `json:"location_y" form:"location_y" gorm:"column:location_y;comment:;"` //位置y
	CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 地图 Cmapinfo自定义表名 cmapinfo
func (Cmapinfo) TableName() string {
	return "cmapinfo"
}
