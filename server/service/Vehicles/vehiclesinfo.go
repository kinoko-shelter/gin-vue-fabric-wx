package Vehicles

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Vehicles"
    VehiclesReq "github.com/flipped-aurora/gin-vue-admin/server/model/Vehicles/request"
    "gorm.io/gorm"
)

type VehiclesinfoService struct {
}

// CreateVehiclesinfo 创建车辆信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (VehiclesinfoCarService *VehiclesinfoService) CreateVehiclesinfo(VehiclesinfoCar *Vehicles.Vehiclesinfo) (err error) {
	err = global.GVA_DB.Create(VehiclesinfoCar).Error
	return err
}

// DeleteVehiclesinfo 删除车辆信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (VehiclesinfoCarService *VehiclesinfoService)DeleteVehiclesinfo(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&Vehicles.Vehiclesinfo{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&Vehicles.Vehiclesinfo{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteVehiclesinfoByIds 批量删除车辆信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (VehiclesinfoCarService *VehiclesinfoService)DeleteVehiclesinfoByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&Vehicles.Vehiclesinfo{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&Vehicles.Vehiclesinfo{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateVehiclesinfo 更新车辆信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (VehiclesinfoCarService *VehiclesinfoService)UpdateVehiclesinfo(VehiclesinfoCar Vehicles.Vehiclesinfo) (err error) {
	err = global.GVA_DB.Save(&VehiclesinfoCar).Error
	return err
}

// GetVehiclesinfo 根据ID获取车辆信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (VehiclesinfoCarService *VehiclesinfoService)GetVehiclesinfo(ID string) (VehiclesinfoCar Vehicles.Vehiclesinfo, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&VehiclesinfoCar).Error
	return
}

// GetVehiclesinfoInfoList 分页获取车辆信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (VehiclesinfoCarService *VehiclesinfoService)GetVehiclesinfoInfoList(info VehiclesReq.VehiclesinfoSearch) (list []Vehicles.Vehiclesinfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Vehicles.Vehiclesinfo{})
    var VehiclesinfoCars []Vehicles.Vehiclesinfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Car_Id != nil {
        db = db.Where("car__id = ?",info.Car_Id)
    }
    if info.Speed != "" {
        db = db.Where("speed = ?",info.Speed)
    }
    if info.Mileage_pre != "" {
        db = db.Where("mileage_pre = ?",info.Mileage_pre)
    }
    if info.Car_locationx != "" {
        db = db.Where("car_locationx = ?",info.Car_locationx)
    }
    if info.Car_locationy != "" {
        db = db.Where("car_locationy = ?",info.Car_locationy)
    }
    if info.Soc != "" {
        db = db.Where("soc = ?",info.Soc)
    }
    if info.Total_mileage != "" {
        db = db.Where("total_mileage = ?",info.Total_mileage)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&VehiclesinfoCars).Error
	return  VehiclesinfoCars, total, err
}
