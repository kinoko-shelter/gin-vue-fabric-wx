// 自动生成模板ChargingStatus
package ChargingStatus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 用户充电状态表 结构体  ChargingStatus
type ChargingStatus struct {
 global.GVA_MODEL 
      User_id  string `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id;"`  //用户id 
      Car_id  string `json:"car_id" form:"car_id" gorm:"column:car_id;comment:车辆id;"`  //车辆id 
      Station_Id  string `json:"station_Id" form:"station_Id" gorm:"column:station__id;comment:充电站id;"`  //充电站id 
      Arrival_status  *int `json:"arrival_status" form:"arrival_status" gorm:"column:arrival_status;comment:到达状态（0:未到达, 1:已到达）;"`  //到达状态 
      Charging_status  *int `json:"charging_status" form:"charging_status" gorm:"column:charging_status;comment:充电状态（0:未充电, 1:正在充电, 2:充电完成）;"`  //充电状态 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 用户充电状态表 ChargingStatus自定义表名 ChargingStatus
func (ChargingStatus) TableName() string {
  return "ChargingStatus"
}

