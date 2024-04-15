package ChargingStatus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ChargingStatus"
    ChargingStatusReq "github.com/flipped-aurora/gin-vue-admin/server/model/ChargingStatus/request"
    "gorm.io/gorm"
)

type ChargingStatusService struct {
}

// CreateChargingStatus 创建用户充电状态表记录
// Author [piexlmax](https://github.com/piexlmax)
func (ChargingStatusInfoService *ChargingStatusService) CreateChargingStatus(ChargingStatusInfo *ChargingStatus.ChargingStatus) (err error) {
	err = global.GVA_DB.Create(ChargingStatusInfo).Error
	return err
}

// DeleteChargingStatus 删除用户充电状态表记录
// Author [piexlmax](https://github.com/piexlmax)
func (ChargingStatusInfoService *ChargingStatusService)DeleteChargingStatus(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&ChargingStatus.ChargingStatus{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&ChargingStatus.ChargingStatus{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteChargingStatusByIds 批量删除用户充电状态表记录
// Author [piexlmax](https://github.com/piexlmax)
func (ChargingStatusInfoService *ChargingStatusService)DeleteChargingStatusByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&ChargingStatus.ChargingStatus{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&ChargingStatus.ChargingStatus{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateChargingStatus 更新用户充电状态表记录
// Author [piexlmax](https://github.com/piexlmax)
func (ChargingStatusInfoService *ChargingStatusService)UpdateChargingStatus(ChargingStatusInfo ChargingStatus.ChargingStatus) (err error) {
	err = global.GVA_DB.Save(&ChargingStatusInfo).Error
	return err
}

// GetChargingStatus 根据ID获取用户充电状态表记录
// Author [piexlmax](https://github.com/piexlmax)
func (ChargingStatusInfoService *ChargingStatusService)GetChargingStatus(ID string) (ChargingStatusInfo ChargingStatus.ChargingStatus, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&ChargingStatusInfo).Error
	return
}

// GetChargingStatusInfoList 分页获取用户充电状态表记录
// Author [piexlmax](https://github.com/piexlmax)
func (ChargingStatusInfoService *ChargingStatusService)GetChargingStatusInfoList(info ChargingStatusReq.ChargingStatusSearch) (list []ChargingStatus.ChargingStatus, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&ChargingStatus.ChargingStatus{})
    var ChargingStatusInfos []ChargingStatus.ChargingStatus
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.User_id != "" {
        db = db.Where("user_id = ?",info.User_id)
    }
    if info.Car_id != "" {
        db = db.Where("car_id = ?",info.Car_id)
    }
    if info.Station_Id != "" {
        db = db.Where("station__id = ?",info.Station_Id)
    }
    if info.Arrival_status != nil {
        db = db.Where("arrival_status = ?",info.Arrival_status)
    }
    if info.Charging_status != nil {
        db = db.Where("charging_status = ?",info.Charging_status)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&ChargingStatusInfos).Error
	return  ChargingStatusInfos, total, err
}
