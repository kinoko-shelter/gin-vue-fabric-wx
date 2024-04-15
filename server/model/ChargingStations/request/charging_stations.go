package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type ChargingStationsSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      Wait_num  *int `json:"wait_num" form:"wait_num" `
                      Charge_num  string `json:"charge_num" form:"charge_num" `
                      Total_service  string `json:"total_service" form:"total_service" `
                      Station_locationx  string `json:"station_locationx" form:"station_locationx" `
                      Station_locationy  string `json:"station_locationy" form:"station_locationy" `
                      Station_name  string `json:"station_name" form:"station_name" `
    request.PageInfo
}
