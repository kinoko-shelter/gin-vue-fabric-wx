package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type VehiclesinfoSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      Car_Id  *int `json:"car_Id" form:"car_Id" `
                      Speed  string `json:"speed" form:"speed" `
                      Mileage_pre  string `json:"mileage_pre" form:"mileage_pre" `
                      Car_locationx  string `json:"car_locationx" form:"car_locationx" `
                      Car_locationy  string `json:"car_locationy" form:"car_locationy" `
                      Soc  string `json:"soc" form:"soc" `
                      Total_mileage  string `json:"total_mileage" form:"total_mileage" `
    request.PageInfo
}
