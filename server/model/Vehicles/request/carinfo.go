package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CarinfoSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      Car_id  *int `json:"car_id" form:"car_id" `
                      Speed  string `json:"speed" form:"speed" `
                      MileagePre  string `json:"mileagePre" form:"mileagePre" `
                      CarLocationx  string `json:"carLocationx" form:"carLocationx" `
                      CarLocationy  string `json:"carLocationy" form:"carLocationy" `
                      Soc  string `json:"soc" form:"soc" `
                      TotalMileage  string `json:"totalMileage" form:"totalMileage" `
                      Type  string `json:"type" form:"type" `
                      CurrentState  string `json:"currentState" form:"currentState" `
    request.PageInfo
}
