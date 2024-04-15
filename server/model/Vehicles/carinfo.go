// 自动生成模板Carinfo
package Vehicles

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 车辆信息表 结构体  Carinfo
type Carinfo struct {
 global.GVA_MODEL 
      Car_id  string `json:"car_id" form:"car_id" gorm:"column:car__id;comment:车辆id;size:255;"`  //车辆id 
      Speed  string `json:"speed" form:"speed" gorm:"column:speed;comment:速度;size:191;"`  //速度 
      MileagePre  string `json:"mileagePre" form:"mileagePre" gorm:"column:mileage_pre;comment:车辆日前预测里程;size:191;"`  //车辆日前预测里程 
      CarLocationx  string `json:"carLocationx" form:"carLocationx" gorm:"column:car_locationx;comment:车辆位置x;size:191;"`  //车辆位置x 
      CarLocationy  string `json:"carLocationy" form:"carLocationy" gorm:"column:car_locationy;comment:车辆位置y;size:191;"`  //车辆位置y 
      Soc  string `json:"soc" form:"soc" gorm:"column:soc;comment:电量;size:191;"`  //电量 
      TotalMileage  string `json:"totalMileage" form:"totalMileage" gorm:"column:total_mileage;comment:累计行驶里程;size:191;"`  //累计行驶里程 
      CreatedBy  *int `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"`  //创建者 
      UpdatedBy  *int `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"`  //更新者 
      DeletedBy  *int `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"`  //删除者 
      Type  string `json:"type" form:"type" gorm:"column:type;comment:车辆类型;"`  //车辆类型 
      CurrentState  string `json:"currentState" form:"currentState" gorm:"column:current_state;comment:当前状态;"`  //当前状态 
}


// TableName 车辆信息表 Carinfo自定义表名 carinfo
func (Carinfo) TableName() string {
  return "carinfo"
}

