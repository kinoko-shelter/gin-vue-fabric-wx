package Roads

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Roads"
    RoadsReq "github.com/flipped-aurora/gin-vue-admin/server/model/Roads/request"
    "gorm.io/gorm"
)

type RoadsService struct {
}

// CreateRoads 创建道路信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (RoadsInfoService *RoadsService) CreateRoads(RoadsInfo *Roads.Roads) (err error) {
	err = global.GVA_DB.Create(RoadsInfo).Error
	return err
}

// DeleteRoads 删除道路信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (RoadsInfoService *RoadsService)DeleteRoads(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&Roads.Roads{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&Roads.Roads{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteRoadsByIds 批量删除道路信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (RoadsInfoService *RoadsService)DeleteRoadsByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&Roads.Roads{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&Roads.Roads{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateRoads 更新道路信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (RoadsInfoService *RoadsService)UpdateRoads(RoadsInfo Roads.Roads) (err error) {
	err = global.GVA_DB.Save(&RoadsInfo).Error
	return err
}

// GetRoads 根据ID获取道路信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (RoadsInfoService *RoadsService)GetRoads(ID string) (RoadsInfo Roads.Roads, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&RoadsInfo).Error
	return
}

// GetRoadsInfoList 分页获取道路信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (RoadsInfoService *RoadsService)GetRoadsInfoList(info RoadsReq.RoadsSearch) (list []Roads.Roads, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Roads.Roads{})
    var RoadsInfos []Roads.Roads
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Road_id != nil {
        db = db.Where("road_id = ?",info.Road_id)
    }
    if info.Road_len != "" {
        db = db.Where("road_len LIKE ?","%"+ info.Road_len+"%")
    }
    if info.Road_limit != "" {
        db = db.Where("road_limit LIKE ?","%"+ info.Road_limit+"%")
    }
    if info.R_startx != "" {
        db = db.Where("r_startx = ?",info.R_startx)
    }
    if info.R_starty != "" {
        db = db.Where("r_starty = ?",info.R_starty)
    }
    if info.R_endx != "" {
        db = db.Where("r_endx = ?",info.R_endx)
    }
    if info.R_endy != "" {
        db = db.Where("r_endy = ?",info.R_endy)
    }
    if info.R_JamIndex != "" {
        db = db.Where("r__jam_index = ?",info.R_JamIndex)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&RoadsInfos).Error
	return  RoadsInfos, total, err
}
