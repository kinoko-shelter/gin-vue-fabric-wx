package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/ChargingStations"
	"github.com/flipped-aurora/gin-vue-admin/server/service/ChargingStatus"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Roads"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Vehicles"
	"github.com/flipped-aurora/gin-vue-admin/server/service/cmap"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/vehiclesinfo"
)

type ServiceGroup struct {
	SystemServiceGroup           system.ServiceGroup
	ExampleServiceGroup          example.ServiceGroup
	VehiclesServiceGroup         Vehicles.ServiceGroup
	RoadsServiceGroup            Roads.ServiceGroup
	ChargingStationsServiceGroup ChargingStations.ServiceGroup
	ChargingStatusServiceGroup   ChargingStatus.ServiceGroup
	VehiclesinfoServiceGroup     vehiclesinfo.ServiceGroup
	CmapServiceGroup             cmap.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
