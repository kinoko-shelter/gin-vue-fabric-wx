package cmap

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmap"
    cmapReq "github.com/flipped-aurora/gin-vue-admin/server/model/cmap/request"
    "gorm.io/gorm"
)

type CarmapService struct {
}

// CreateCarmap 创建carmap记录
// Author [piexlmax](https://github.com/piexlmax)
func (carmapService *CarmapService) CreateCarmap(carmap *cmap.Carmap) (err error) {
	err = global.GVA_DB.Create(carmap).Error
	return err
}

// DeleteCarmap 删除carmap记录
// Author [piexlmax](https://github.com/piexlmax)
func (carmapService *CarmapService)DeleteCarmap(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&cmap.Carmap{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&cmap.Carmap{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteCarmapByIds 批量删除carmap记录
// Author [piexlmax](https://github.com/piexlmax)
func (carmapService *CarmapService)DeleteCarmapByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&cmap.Carmap{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&cmap.Carmap{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateCarmap 更新carmap记录
// Author [piexlmax](https://github.com/piexlmax)
func (carmapService *CarmapService)UpdateCarmap(carmap cmap.Carmap) (err error) {
	err = global.GVA_DB.Save(&carmap).Error
	return err
}

// GetCarmap 根据ID获取carmap记录
// Author [piexlmax](https://github.com/piexlmax)
func (carmapService *CarmapService)GetCarmap(ID string) (carmap cmap.Carmap, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&carmap).Error
	return
}

// GetCarmapInfoList 分页获取carmap记录
// Author [piexlmax](https://github.com/piexlmax)
func (carmapService *CarmapService)GetCarmapInfoList(info cmapReq.CarmapSearch) (list []cmap.Carmap, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cmap.Carmap{})
    var carmaps []cmap.Carmap
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Location_x != "" {
        db = db.Where("location_x = ?",info.Location_x)
    }
    if info.Location_y != "" {
        db = db.Where("location_y = ?",info.Location_y)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&carmaps).Error
	return  carmaps, total, err
}
