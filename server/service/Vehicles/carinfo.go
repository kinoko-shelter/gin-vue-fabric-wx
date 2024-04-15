package Vehicles

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Vehicles"
    VehiclesReq "github.com/flipped-aurora/gin-vue-admin/server/model/Vehicles/request"
)

type CarinfoService struct {
}

// CreateCarinfo 创建车辆信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (carinfoService *CarinfoService) CreateCarinfo(carinfo *Vehicles.Carinfo) (err error) {
	err = global.GVA_DB.Create(carinfo).Error
	return err
}

// DeleteCarinfo 删除车辆信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (carinfoService *CarinfoService)DeleteCarinfo(ID string) (err error) {
	err = global.GVA_DB.Delete(&Vehicles.Carinfo{},"id = ?",ID).Error
	return err
}

// DeleteCarinfoByIds 批量删除车辆信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (carinfoService *CarinfoService)DeleteCarinfoByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]Vehicles.Carinfo{},"id in ?",IDs).Error
	return err
}

// UpdateCarinfo 更新车辆信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (carinfoService *CarinfoService)UpdateCarinfo(carinfo Vehicles.Carinfo) (err error) {
	err = global.GVA_DB.Save(&carinfo).Error
	return err
}

// GetCarinfo 根据ID获取车辆信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (carinfoService *CarinfoService)GetCarinfo(ID string) (carinfo Vehicles.Carinfo, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&carinfo).Error
	return
}

// GetCarinfoInfoList 分页获取车辆信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (carinfoService *CarinfoService)GetCarinfoInfoList(info VehiclesReq.CarinfoSearch) (list []Vehicles.Carinfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Vehicles.Carinfo{})
    var carinfos []Vehicles.Carinfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Car_id != nil {
        db = db.Where("car__id = ?",info.Car_id)
    }
    if info.Speed != "" {
        db = db.Where("speed > ?",info.Speed)
    }
    if info.MileagePre != "" {
        db = db.Where("mileage_pre = ?",info.MileagePre)
    }
    if info.CarLocationx != "" {
        db = db.Where("car_locationx = ?",info.CarLocationx)
    }
    if info.CarLocationy != "" {
        db = db.Where("car_locationy = ?",info.CarLocationy)
    }
    if info.Soc != "" {
        db = db.Where("soc LIKE ?","%"+ info.Soc+"%")
    }
    if info.TotalMileage != "" {
        db = db.Where("total_mileage > ?",info.TotalMileage)
    }
    if info.Type != "" {
        db = db.Where("type = ?",info.Type)
    }
    if info.CurrentState != "" {
        db = db.Where("current_state = ?",info.CurrentState)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&carinfos).Error
	return  carinfos, total, err
}
