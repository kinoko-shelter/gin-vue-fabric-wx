package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/ChargingStations"
	"github.com/flipped-aurora/gin-vue-admin/server/router/ChargingStatus"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Roads"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Vehicles"
	"github.com/flipped-aurora/gin-vue-admin/server/router/cmap"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/vehiclesinfo"
)

type RouterGroup struct {
	System           system.RouterGroup
	Example          example.RouterGroup
	Vehicles         Vehicles.RouterGroup
	Roads            Roads.RouterGroup
	ChargingStations ChargingStations.RouterGroup
	ChargingStatus   ChargingStatus.RouterGroup
	Vehiclesinfo     vehiclesinfo.RouterGroup
	Cmap             cmap.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
