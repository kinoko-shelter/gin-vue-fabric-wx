package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/ChargingStations"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/ChargingStatus"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Roads"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Vehicles"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/cmap"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/vehiclesinfo"
)

type ApiGroup struct {
	SystemApiGroup           system.ApiGroup
	ExampleApiGroup          example.ApiGroup
	VehiclesApiGroup         Vehicles.ApiGroup
	RoadsApiGroup            Roads.ApiGroup
	ChargingStationsApiGroup ChargingStations.ApiGroup
	ChargingStatusApiGroup   ChargingStatus.ApiGroup
	VehiclesinfoApiGroup     vehiclesinfo.ApiGroup
	CmapApiGroup             cmap.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
