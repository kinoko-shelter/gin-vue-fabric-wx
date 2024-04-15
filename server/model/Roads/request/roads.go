package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type RoadsSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      Road_id  *int `json:"road_id" form:"road_id" `
                      Road_len  string `json:"road_len" form:"road_len" `
                      Road_limit  string `json:"road_limit" form:"road_limit" `
                      R_startx  string `json:"r_startx" form:"r_startx" `
                      R_starty  string `json:"r_starty" form:"r_starty" `
                      R_endx  string `json:"r_endx" form:"r_endx" `
                      R_endy  string `json:"r_endy" form:"r_endy" `
                      R_JamIndex  string `json:"r_JamIndex" form:"r_JamIndex" `
    request.PageInfo
}
