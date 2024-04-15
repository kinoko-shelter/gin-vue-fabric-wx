// 自动生成模板ChargingStations
package ChargingStations

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 充电站信息表 结构体  ChargingStations
type ChargingStations struct {
 global.GVA_MODEL 
      Station_Id  *int `json:"station_Id" form:"station_Id" gorm:"column:station__id;comment:充电站id;"`  //充电站id 
      Pile_num  *int `json:"pile_num" form:"pile_num" gorm:"column:pile_num;comment:充电桩个数;"`  //充电桩个数 
      Wait_num  *int `json:"wait_num" form:"wait_num" gorm:"column:wait_num;comment:排队车辆数;"`  //排队车辆数 
      Charge_num  string `json:"charge_num" form:"charge_num" gorm:"column:charge_num;comment:充电车辆数;"`  //充电车辆数 
      Total_service  string `json:"total_service" form:"total_service" gorm:"column:total_service;comment:累计服务车次;"`  //累计服务车次 
      Station_locationx  string `json:"station_locationx" form:"station_locationx" gorm:"column:station_locationx;comment:充电站位置x;"`  //充电站位置x 
      Station_locationy  string `json:"station_locationy" form:"station_locationy" gorm:"column:station_locationy;comment:充电站位置y;"`  //充电站位置y 
      Station_name  string `json:"station_name" form:"station_name" gorm:"column:station_name;comment:充电站名称;"`  //充电站名称 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 充电站信息表 ChargingStations自定义表名 ChargingStations
func (ChargingStations) TableName() string {
  return "ChargingStations"
}

