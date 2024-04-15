package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type ChargingStatusSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      User_id  string `json:"user_id" form:"user_id" `
                      Car_id  string `json:"car_id" form:"car_id" `
                      Station_Id  string `json:"station_Id" form:"station_Id" `
                      Arrival_status  *int `json:"arrival_status" form:"arrival_status" `
                      Charging_status  *int `json:"charging_status" form:"charging_status" `
    request.PageInfo
}
