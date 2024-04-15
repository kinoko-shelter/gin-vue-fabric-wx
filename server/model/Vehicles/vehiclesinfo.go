// 自动生成模板Vehiclesinfo
package Vehicles

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 车辆信息表 结构体  Vehiclesinfo
type Vehiclesinfo struct {
 global.GVA_MODEL 
      Car_Id  *int `json:"car_Id" form:"car_Id" gorm:"column:car__id;comment:车辆id;"`  //车辆id 
      Speed  string `json:"speed" form:"speed" gorm:"column:speed;comment:速度;"`  //速度 
      Mileage_pre  string `json:"mileage_pre" form:"mileage_pre" gorm:"column:mileage_pre;comment:车辆日前预测里程;"`  //车辆日前预测里程 
      Car_locationx  string `json:"car_locationx" form:"car_locationx" gorm:"column:car_locationx;comment:车辆位置x;"`  //车辆位置x 
      Car_locationy  string `json:"car_locationy" form:"car_locationy" gorm:"column:car_locationy;comment:车辆位置y;"`  //车辆位置y 
      Soc  string `json:"soc" form:"soc" gorm:"column:soc;comment:电量;"`  //电量 
      Total_mileage  string `json:"total_mileage" form:"total_mileage" gorm:"column:total_mileage;comment:累计行驶里程;"`  //累计行驶里程 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 车辆信息表 Vehiclesinfo自定义表名 Vehiclesinfo
func (Vehiclesinfo) TableName() string {
  return "Vehiclesinfo"
}

