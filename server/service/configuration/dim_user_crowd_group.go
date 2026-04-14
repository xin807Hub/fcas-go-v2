package configuration

import (
	"errors"
	"fcas_server/model/configuration"
	"fcas_server/model/configuration/req"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DimUserCrowdGroupSvc struct {
	Log   *zap.Logger
	Mysql *gorm.DB
}

func NewDimUserCrowdGroupSvc(log *zap.Logger, mysql *gorm.DB) *DimUserCrowdGroupSvc {
	return &DimUserCrowdGroupSvc{
		Log:   log.Named("[DimUserCrowdGroup-用户组群管理]"),
		Mysql: mysql,
	}
}

func (svc DimUserCrowdGroupSvc) GetById(id string) (result configuration.DimUserCrowdGroup, err error) {

	err = svc.Mysql.Model(configuration.DimUserCrowdGroup{}).First(&result, id).Error
	if err != nil {
		svc.Log.Error("根据ID获取信息失败", zap.Error(err))
		return result, err
	}

	crowds, err := getCrowdByGroupId(svc.Mysql, result.ID)
	if err != nil {
		svc.Log.Error("根据用户组群ID获取用户组信息失败", zap.Any("group_id", result.ID), zap.Error(err))
		return result, fmt.Errorf("获取该用户组群的用户组信息失败: %w", err)
	}
	result.Crowds = crowds

	return result, nil
}

func (svc DimUserCrowdGroupSvc) List(req req.ListRequest) (result []*configuration.DimUserCrowdGroup, total int64, err error) {
	tx := svc.Mysql.Model(configuration.DimUserCrowdGroup{}).Debug()
	if req.Key != "" {
		tx = tx.Where("group_name LIKE ?", fmt.Sprint("%", req.Key, "%"))
	}

	// 分页
	err = tx.Count(&total).Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Find(&result).Error
	if err != nil {
		svc.Log.Error("获取列表信息失败", zap.Error(err))
		return result, total, err
	}

	// 填充组信息
	for _, group := range result {
		crowds, err := getCrowdByGroupId(svc.Mysql, group.ID)
		if err != nil {
			svc.Log.Error("根据用户组群ID获取用户组信息失败", zap.Any("group_id", group.ID), zap.Error(err))
			continue
		}
		group.Crowds = crowds
	}

	return result, total, nil
}

func (svc DimUserCrowdGroupSvc) Save(req req.DimUserCrowdGroupSaveRequest) error {

	err := svc.Mysql.Debug().Transaction(func(tx *gorm.DB) error {
		group := configuration.DimUserCrowdGroup{
			GroupName: req.GroupName,
			Remark:    req.Remark,
		}
		if err := tx.Create(&group).Error; err != nil {
			return fmt.Errorf("创建用户组群信息失败: %w", err)
		}

		relations := make([]configuration.DimUserCrowdGroupRelation, 0, len(req.CrowdIds))
		for _, id := range req.CrowdIds {
			relations = append(relations, configuration.DimUserCrowdGroupRelation{
				GroupID: group.ID,
				CrowdID: id,
			})
		}

		if err := tx.Create(&relations).Error; err != nil {
			return fmt.Errorf("创建用户组群与用户组关系失败: %w", err)
		}

		return nil
	})

	if err != nil {
		svc.Log.Error("创建失败", zap.Error(err))
		return err
	}

	return nil
}

func (svc DimUserCrowdGroupSvc) Update(req req.DimUserCrowdGroupSaveRequest) error {

	err := svc.Mysql.Transaction(func(tx *gorm.DB) error {
		group := configuration.DimUserCrowdGroup{
			ID:        req.ID,
			GroupName: req.GroupName,
			Remark:    req.Remark,
		}
		if err := tx.Save(&group).Error; err != nil {
			return fmt.Errorf("更新用户组群信息失败: %w", err)
		}

		// 删除原有关系 dim_user_crowd_group_relation
		if err := tx.Delete(&configuration.DimUserCrowdGroupRelation{}, "group_id = ?", group.ID).Error; err != nil {
			return fmt.Errorf("删除原有用户组群与用户组关系失败: %w", err)
		}

		// 创建新的关系 dim_user_crowd_group_relation
		relations := make([]configuration.DimUserCrowdGroupRelation, 0, len(req.CrowdIds))
		for _, id := range req.CrowdIds {
			relations = append(relations, configuration.DimUserCrowdGroupRelation{
				GroupID: group.ID,
				CrowdID: id,
			})
		}
		if err := tx.Create(&relations).Error; err != nil {
			return fmt.Errorf("更新用户组群组与用户组关系失败: %w", err)
		}

		return nil
	})

	if err != nil {
		svc.Log.Error("更新失败", zap.Error(err))
		return err
	}

	return nil
}

func (svc DimUserCrowdGroupSvc) Delete(ids []int) error {

	err := svc.Mysql.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&configuration.DimUserCrowdGroup{}, ids).Error; err != nil {
			return fmt.Errorf("删除用户组群信息失败: %w", err)
		}
		if err := tx.Delete(&configuration.DimUserCrowdGroupRelation{}, "group_id IN (?)", ids).Error; err != nil {
			return fmt.Errorf("删除用户组群与用户组关系失败: %w", err)
		}
		return nil
	})

	if err != nil {
		svc.Log.Error("删除失败", zap.Error(err))
		return err
	}

	return nil
}

// ValidateUnique 校验用户群组名称是否存在, id为0表示新增操作, id不为0表示更新操作（更新时需要排除当前记录）
func (svc DimUserCrowdGroupSvc) ValidateUnique(name string, id int) error {
	var count int64
	tx := svc.Mysql.Model(configuration.DimUserCrowdGroup{})
	if id != 0 {
		tx = tx.Where("id != ?", id)
	}
	err := tx.Where("group_name = ?", name).Count(&count).Error
	if err != nil {
		svc.Log.Error("校验用户组群名称时出现异常", zap.Error(err))
		return fmt.Errorf("校验用户组群名称时出现异常: %w", err)
	}

	if count > 0 {
		return errors.New("用户组群名称已存在")
	}

	return nil
}

// getCrowdByGroupId 根据crowdId获取该用户组的所有用户信息
func getCrowdByGroupId(db *gorm.DB, groupId int) (result []configuration.DimUserCrowdGroupCrowd, err error) {
	sql := `
			SELECT c.id, crowd_name 
			FROM dim_user_crowd_group_relation r 
			    LEFT JOIN dim_user_crowd c on c.id = r.crowd_id  
			WHERE r.group_id = ?`

	if err := db.Raw(sql, groupId).Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// getCrowdNotInGroup 获取未分配的用户组
func getCrowdNotInGroup(db *gorm.DB) (result []configuration.DimUserCrowdGroupCrowd, err error) {
	sql := `
			SELECT id, crowd_name
			FROM dim_user_crowd crowd
			WHERE NOT EXISTS (
				SELECT 1
				FROM dim_user_crowd_group_relation r
				WHERE r.crowd_id = crowd.id
                  )`

	if err := db.Raw(sql).Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// getUserNotInCrowd 获取未分配的用户
func getUserNotInCrowd(db *gorm.DB) (result []configuration.DimUserCrowdUser, err error) {
	sql := `
			SELECT id, user_name
			FROM dim_user_info u
			WHERE NOT EXISTS (
				SELECT 1
				FROM dim_user_crowd_relation r
				WHERE r.user_id = u.id
                  )`

	if err := db.Raw(sql).Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// GetGroupTree 获取用户群组树，level为1表示只显示群组, 为2表示显示群组和用户，为3表示显示群组、用户和用户组；
// 使用数据组装的方式（便于后期维护），返回的树结构为GroupTreeNode类型，包含ID、Label、Children三个字段；
// ID为节点的唯一标识，Label为节点的显示名称，Children为子节点列表；
func (svc DimUserCrowdGroupSvc) GetGroupTree(level int) (tree configuration.GroupTreeNode, err error) {
	// 获取所有群组
	var groups []configuration.DimUserCrowdGroup
	if err := svc.Mysql.Model(configuration.DimUserCrowdGroup{}).Find(&groups).Error; err != nil {
		svc.Log.Error("获取所有群组信息失败", zap.Error(err))
		return tree, err
	}

	tree = configuration.GroupTreeNode{
		ID:       "0",
		Label:    "全部",
		Children: nil,
	}

	childrenNode := make([]configuration.GroupTreeNode, 0, len(groups))
	for _, g := range groups {
		groupNode := svc.createGroupNode(g, level)
		childrenNode = append(childrenNode, groupNode)
	}

	if level > 1 {
		// 添加未分配的组和用户
		node := svc.createUnassignedGroupNode(level)
		childrenNode = append(childrenNode, node)
	}

	tree.Children = childrenNode

	return tree, nil
}

// createGroupNode 创建用户组节点, level为1表示只显示群组, 为2表示显示群组和用户
func (svc DimUserCrowdGroupSvc) createGroupNode(group configuration.DimUserCrowdGroup, level int) configuration.GroupTreeNode {
	node := configuration.GroupTreeNode{
		ID:       fmt.Sprint(group.ID),
		Label:    group.GroupName,
		Children: nil,
	}

	if level <= 1 {
		return node
	}

	crowds, err := getCrowdByGroupId(svc.Mysql, group.ID)
	if err != nil {
		svc.Log.Error("根据用户组群ID获取用户组信息失败", zap.Any("group_id", group.ID), zap.Error(err))
		return node
	}

	childrenNode := make([]configuration.GroupTreeNode, 0, len(crowds))
	for _, crowd := range crowds {
		crowdNode := svc.createCrowdNode(group.ID, crowd, level)
		childrenNode = append(childrenNode, crowdNode)
	}
	node.Children = childrenNode

	return node
}

// createCrowdNode 创建用户组节点, level为1表示只显示群组, 为2表示显示群组和用户
func (svc DimUserCrowdGroupSvc) createCrowdNode(groupId int, crowd configuration.DimUserCrowdGroupCrowd, level int) configuration.GroupTreeNode {
	node := configuration.GroupTreeNode{
		ID:       fmt.Sprintf("%d-%d", groupId, crowd.ID),
		Label:    crowd.CrowdName,
		Children: nil,
	}

	if level <= 2 {
		return node
	}

	users, err := getUsersByCrowdId(svc.Mysql, crowd.ID)
	if err != nil {
		svc.Log.Error("根据用户组ID获取用户信息失败", zap.Any("crowd_id", crowd.ID), zap.Error(err))
		return node
	}

	childrenNode := make([]configuration.GroupTreeNode, 0, len(users))
	for _, user := range users {
		childrenNode = append(childrenNode, configuration.GroupTreeNode{
			ID:    fmt.Sprintf("%d-%d-%d", groupId, crowd.ID, user.ID),
			Label: user.UserName,
		})
	}
	node.Children = childrenNode

	return node
}

// createUnassignedGroupNode 创建未分配的群组节点, level为1表示只显示群组, 为2表示显示群组和用户
func (svc DimUserCrowdGroupSvc) createUnassignedGroupNode(level int) configuration.GroupTreeNode {
	node := configuration.GroupTreeNode{
		ID:       "00",
		Label:    "未分配群",
		Children: nil,
	}

	// 未分配的组
	crowds, err := getCrowdNotInGroup(svc.Mysql)
	if err != nil {
		svc.Log.Error("获取未分配的用户组失败", zap.Error(err))
		return node
	}

	childrenCrowdNode := make([]configuration.GroupTreeNode, 0, len(crowds))
	for _, crowd := range crowds {
		childrenCrowdNode = append(childrenCrowdNode, svc.createCrowdNode(0, crowd, level))
	}

	// 未分配的用户
	if level > 2 {
		users, err := getUserNotInCrowd(svc.Mysql)
		if err != nil {
			svc.Log.Error("获取未分配的用户失败", zap.Error(err))
			return node
		}
		childrenUserNode := make([]configuration.GroupTreeNode, 0, len(users))
		for _, user := range users {
			childrenUserNode = append(childrenUserNode, configuration.GroupTreeNode{
				ID:    fmt.Sprintf("0-0-%d", user.ID),
				Label: user.UserName,
			})
		}

		// 将未分配的用户作为单独的crowd节点添加到childrenCrowdNode中
		childrenCrowdNode = append(childrenCrowdNode, configuration.GroupTreeNode{
			ID:       "0-0",
			Label:    "未分配用户",
			Children: childrenUserNode,
		})
	}

	node.Children = childrenCrowdNode

	return node
}
