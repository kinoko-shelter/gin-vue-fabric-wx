package cmap

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cmap"
    cmapReq "github.com/flipped-aurora/gin-vue-admin/server/model/cmap/request"
    "gorm.io/gorm"
)

type CmapinfoService struct {
}

// CreateCmapinfo 创建地图记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmapinfoService *CmapinfoService) CreateCmapinfo(cmapinfo *cmap.Cmapinfo) (err error) {
	err = global.GVA_DB.Create(cmapinfo).Error
	return err
}

// DeleteCmapinfo 删除地图记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmapinfoService *CmapinfoService)DeleteCmapinfo(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&cmap.Cmapinfo{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&cmap.Cmapinfo{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteCmapinfoByIds 批量删除地图记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmapinfoService *CmapinfoService)DeleteCmapinfoByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&cmap.Cmapinfo{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&cmap.Cmapinfo{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateCmapinfo 更新地图记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmapinfoService *CmapinfoService)UpdateCmapinfo(cmapinfo cmap.Cmapinfo) (err error) {
	err = global.GVA_DB.Save(&cmapinfo).Error
	return err
}

// GetCmapinfo 根据ID获取地图记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmapinfoService *CmapinfoService)GetCmapinfo(ID string) (cmapinfo cmap.Cmapinfo, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cmapinfo).Error
	return
}

// GetCmapinfoInfoList 分页获取地图记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmapinfoService *CmapinfoService)GetCmapinfoInfoList(info cmapReq.CmapinfoSearch) (list []cmap.Cmapinfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cmap.Cmapinfo{})
    var cmapinfos []cmap.Cmapinfo
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
	
	err = db.Find(&cmapinfos).Error
	return  cmapinfos, total, err
}
