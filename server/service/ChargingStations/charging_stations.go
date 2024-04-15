package ChargingStations

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ChargingStations"
    ChargingStationsReq "github.com/flipped-aurora/gin-vue-admin/server/model/ChargingStations/request"
    "gorm.io/gorm"
)

type ChargingStationsService struct {
}

// CreateChargingStations 创建充电站信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (ChargingStationsInfoService *ChargingStationsService) CreateChargingStations(ChargingStationsInfo *ChargingStations.ChargingStations) (err error) {
	err = global.GVA_DB.Create(ChargingStationsInfo).Error
	return err
}

// DeleteChargingStations 删除充电站信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (ChargingStationsInfoService *ChargingStationsService)DeleteChargingStations(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&ChargingStations.ChargingStations{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&ChargingStations.ChargingStations{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteChargingStationsByIds 批量删除充电站信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (ChargingStationsInfoService *ChargingStationsService)DeleteChargingStationsByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&ChargingStations.ChargingStations{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&ChargingStations.ChargingStations{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateChargingStations 更新充电站信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (ChargingStationsInfoService *ChargingStationsService)UpdateChargingStations(ChargingStationsInfo ChargingStations.ChargingStations) (err error) {
	err = global.GVA_DB.Save(&ChargingStationsInfo).Error
	return err
}

// GetChargingStations 根据ID获取充电站信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (ChargingStationsInfoService *ChargingStationsService)GetChargingStations(ID string) (ChargingStationsInfo ChargingStations.ChargingStations, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&ChargingStationsInfo).Error
	return
}

// GetChargingStationsInfoList 分页获取充电站信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (ChargingStationsInfoService *ChargingStationsService)GetChargingStationsInfoList(info ChargingStationsReq.ChargingStationsSearch) (list []ChargingStations.ChargingStations, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&ChargingStations.ChargingStations{})
    var ChargingStationsInfos []ChargingStations.ChargingStations
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Wait_num != nil {
        db = db.Where("wait_num = ?",info.Wait_num)
    }
    if info.Charge_num != "" {
        db = db.Where("charge_num = ?",info.Charge_num)
    }
    if info.Total_service != "" {
        db = db.Where("total_service = ?",info.Total_service)
    }
    if info.Station_locationx != "" {
        db = db.Where("station_locationx = ?",info.Station_locationx)
    }
    if info.Station_locationy != "" {
        db = db.Where("station_locationy = ?",info.Station_locationy)
    }
    if info.Station_name != "" {
        db = db.Where("station_name LIKE ?","%"+ info.Station_name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&ChargingStationsInfos).Error
	return  ChargingStationsInfos, total, err
}
