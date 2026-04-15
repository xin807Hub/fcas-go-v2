package initialize

import (
	sysModel "fcas_server/model/system"
	"fcas_server/service/system"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

func SyncApi(mysql *gorm.DB, engine *gin.Engine) error {
	// 获取gin.Engine中的router信息
	routers := getAllRouter(engine)

	// 获取角色id
	var authorityIds []string
	if err := mysql.Table("sys_authorities").Select("authority_id").Find(&authorityIds).Error; err != nil {
		return fmt.Errorf("failed to get all authorities id: %w", err)
	}

	// 组装casbin信息
	authorityCasbinMap := make(map[string][][]string, len(authorityIds))
	for _, authorityId := range authorityIds {
		for _, r := range routers {
			authorityCasbinMap[authorityId] = append(authorityCasbinMap[authorityId], []string{authorityId, r.Path, r.Method})
		}
	}

	var casbinSvc system.CasbinService
	casbin := casbinSvc.Casbin()

	err := mysql.Transaction(func(tx *gorm.DB) error {
		// 软删除
		if err := tx.Where("1 = 1").Delete(&sysModel.SysApi{}).Error; err != nil {
			return fmt.Errorf("failed to delete old api: %w", err)
		}

		//永久删除
		//if err := tx.Where("1 = 1").Unscoped().Delete(&sysModel.SysApi{}).Error; err != nil {
		//	return fmt.Errorf("failed to delete old api: %w", err)
		//}

		if err := tx.Create(&routers).Error; err != nil {
			return fmt.Errorf("failed to creat new api: %w", err)
		}

		for authorityId, rules := range authorityCasbinMap {
			casbinSvc.ClearCasbin(0, authorityId)
			succ, _ := casbin.AddPolicies(rules)
			if !succ {
				return fmt.Errorf("角色（%s）存在相同API，添加失败", authorityId)
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// 获取gin Engine中所有路由信息
func getAllRouter(engine *gin.Engine) []sysModel.SysApi {
	routers := make([]sysModel.SysApi, 0, len(engine.Routes()))
	for _, r := range engine.Routes() {
		group, desc := extractPath(r.Path)
		routers = append(routers, sysModel.SysApi{
			Method:      r.Method,
			Path:        r.Path,
			ApiGroup:    group,
			Description: desc,
		})
	}

	return routers
}

func extractPath(path string) (group string, desc string) {
	parts := strings.SplitN(strings.TrimPrefix(path, "/"), "/", 2)
	switch len(parts) {
	case 1:
		group = parts[0]
		desc = ""
	case 2:
		group = parts[0]
		desc = parts[1]
	default:
		panic("invalid route path: " + path)
	}

	return group, desc
}
