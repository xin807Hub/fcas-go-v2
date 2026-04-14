package configuration

import (
	"errors"
	"fcas_server/model/configuration"
	"fcas_server/model/configuration/req"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DimUserCrowdSvc struct {
	Log   *zap.Logger
	Mysql *gorm.DB
}

func NewDimUserCrowdSvc(log *zap.Logger, mysql *gorm.DB) *DimUserCrowdSvc {
	return &DimUserCrowdSvc{
		Log:   log.Named("[DimUserCrowd-用户组管理]"),
		Mysql: mysql,
	}
}

func (svc DimUserCrowdSvc) GetById(id string) (result configuration.DimUserCrowd, err error) {

	err = svc.Mysql.Model(configuration.DimUserCrowd{}).First(&result, id).Error
	if err != nil {
		svc.Log.Error("根据ID获取信息失败", zap.Error(err))
		return result, err
	}

	users, err := getUsersByCrowdId(svc.Mysql, result.ID)
	if err != nil {
		svc.Log.Error("根据群ID获取用户信息失败", zap.Any("crowd_id", result.ID), zap.Error(err))
		return result, fmt.Errorf("获取该用户组的用户信息失败: %w", err)
	}
	result.Users = users

	return result, nil
}

func (svc DimUserCrowdSvc) List(req req.ListRequest) (result []*configuration.DimUserCrowd, total int64, err error) {
	tx := svc.Mysql.Model(configuration.DimUserCrowd{}).Debug()
	if req.Key != "" {
		tx = tx.Where("crowd_name LIKE ?", fmt.Sprint("%", req.Key, "%"))
	}

	// 分页
	err = tx.Count(&total).Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Find(&result).Error
	if err != nil {
		svc.Log.Error("获取列表信息失败", zap.Error(err))
		return result, total, err
	}

	// 填充用户信息
	for _, crowd := range result {
		users, err := getUsersByCrowdId(svc.Mysql, crowd.ID)
		if err != nil {
			svc.Log.Error("根据群ID获取用户信息失败", zap.Any("crowd_id", crowd.ID), zap.Error(err))
			continue
		}
		crowd.Users = users
	}

	return result, total, nil
}

func (svc DimUserCrowdSvc) Save(req req.DimUserCrowdSaveRequest) error {

	err := svc.Mysql.Transaction(func(tx *gorm.DB) error {
		userCrowd := configuration.DimUserCrowd{
			CrowdName: req.CrowdName,
			Remark:    req.Remark,
		}
		if err := tx.Create(&userCrowd).Error; err != nil {
			return fmt.Errorf("创建组信息失败: %w", err)
		}

		relations := make([]configuration.DimUserCrowdRelation, 0, len(req.UserIds))
		for _, userId := range req.UserIds {
			relations = append(relations, configuration.DimUserCrowdRelation{
				CrowdID: userCrowd.ID,
				UserID:  userId,
			})
		}

		if err := tx.Create(&relations).Error; err != nil {
			return fmt.Errorf("创建组与用户关系失败: %w", err)
		}

		return nil
	})

	if err != nil {
		svc.Log.Error("创建失败", zap.Error(err))
		return err
	}

	return nil
}

func (svc DimUserCrowdSvc) Update(req req.DimUserCrowdSaveRequest) error {

	err := svc.Mysql.Transaction(func(tx *gorm.DB) error {
		crowd := configuration.DimUserCrowd{
			ID:        req.ID,
			CrowdName: req.CrowdName,
			Remark:    req.Remark,
		}
		if err := tx.Save(&crowd).Error; err != nil {
			svc.Log.Error("更新组信息失败", zap.Error(err))
			return fmt.Errorf("更新组信息失败: %w", err)
		}

		// 删除原有关系 dim_user_crowd_relation
		if err := tx.Delete(&configuration.DimUserCrowdRelation{}, "crowd_id = ?", crowd.ID).Error; err != nil {
			return fmt.Errorf("删除原有组与用户关系失败: %w", err)
		}

		// 新增关系 dim_user_crowd_relation
		relations := make([]configuration.DimUserCrowdRelation, 0, len(req.UserIds))
		for _, userId := range req.UserIds {
			relations = append(relations, configuration.DimUserCrowdRelation{
				CrowdID: crowd.ID,
				UserID:  userId,
			})
		}

		if err := tx.Create(&relations).Error; err != nil {
			return fmt.Errorf("更新组与用户关系失败: %w", err)
		}

		return nil
	})

	if err != nil {
		svc.Log.Error("更新失败", zap.Error(err))
		return err
	}

	return nil
}

func (svc DimUserCrowdSvc) Delete(ids []int) error {

	err := svc.Mysql.Transaction(func(tx *gorm.DB) error {
		var count int64
		if err := tx.Raw(`SELECT COUNT(*) FROM dim_user_crowd_group_relation WHERE crowd_id IN ?`, ids).Scan(&count).Error; err != nil {
			return fmt.Errorf("查询待删除用户组是否被用户组群引用失败: %w", err)
		}
		if count != 0 {
			return errors.New("该用户组被用户组群引用，请先解除引用")
		}
		if err := tx.Delete(&configuration.DimUserCrowd{}, ids).Error; err != nil {
			return fmt.Errorf("删除组信息失败: %w", err)
		}
		if err := tx.Delete(&configuration.DimUserCrowdRelation{}, "crowd_id IN (?)", ids).Error; err != nil {
			return fmt.Errorf("删除组与用户关系失败: %w", err)
		}
		return nil
	})

	if err != nil {
		svc.Log.Error("删除失败", zap.Any("ids", ids), zap.Error(err))
		return err
	}

	return nil
}

// ValidateUnique 校验用户组名称是否存在, id为0表示新增操作, id不为0表示更新操作（更新时需要排除当前记录）
func (svc DimUserCrowdSvc) ValidateUnique(name string, id int) error {
	tx := svc.Mysql.Model(configuration.DimUserCrowd{})
	if id != 0 {
		tx = tx.Where("id != ?", id)
	}
	var count int64
	err := tx.Where("crowd_name = ?", name).Count(&count).Error
	if err != nil {
		svc.Log.Error("校验用户组名称时出现异常", zap.Error(err))
		return fmt.Errorf("校验用户组名称时出现异常: %w", err)
	}

	if count > 0 {
		return errors.New("用户组名称已存在")
	}

	return nil
}

// getUsersByCrowdId 根据crowdId获取该用户组的所有用户信息
func getUsersByCrowdId(db *gorm.DB, crowdId int) (result []configuration.DimUserCrowdUser, err error) {
	sql := `
			SELECT u.id AS id, user_name 
			FROM dim_user_crowd_relation r 
			    LEFT JOIN dim_user_info u ON u.id = r.user_id 
			WHERE r.crowd_id = ?`

	if err := db.Raw(sql, crowdId).Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
