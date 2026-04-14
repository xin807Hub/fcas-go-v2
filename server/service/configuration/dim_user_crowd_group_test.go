package configuration

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
)

type DimUserCrowdGroupEntity struct {
	ID   int
	Name string
}

type DimUserCrowdEntity struct {
	ID   int
	Name string
}

type DimUserInfoEntity struct {
	ID       int
	UserName string
}

// 模拟服务
type GroupService struct{}
type CrowdService struct{}
type UserService struct{}

func (g *GroupService) ListGroups() []DimUserCrowdGroupEntity {
	return []DimUserCrowdGroupEntity{
		{ID: 1, Name: "Group A"},
		{ID: 2, Name: "Group B"},
	}
}

func (c *CrowdService) GetUserCrowdsByGroupID(groupID int) []DimUserCrowdEntity {
	return []DimUserCrowdEntity{
		{ID: 1, Name: "Crowd X"},
		{ID: 2, Name: "Crowd Y"},
	}
}

func (c *CrowdService) GetUserCrowdsNotInGroup() []DimUserCrowdEntity {
	return []DimUserCrowdEntity{
		{ID: 3, Name: "Unassigned Crowd"},
	}
}

func (u *UserService) GetUsersByCrowdID(crowdID int) []DimUserInfoEntity {
	return []DimUserInfoEntity{
		{ID: 1, UserName: "User 1"},
		{ID: 2, UserName: "User 2"},
	}
}

func (u *UserService) GetUsersNotInAnyCrowd() []DimUserInfoEntity {
	return []DimUserInfoEntity{
		{ID: 3, UserName: "Unassigned User"},
	}
}

// 构建用户树的主函数
func getUserTree(level int) []map[string]interface{} {
	groupService := GroupService{}
	crowdService := CrowdService{}
	userService := UserService{}

	var tree []map[string]interface{}

	// 获取所有的分组
	groups := groupService.ListGroups()
	for _, group := range groups {
		groupNode := createGroupNode(group, level, &crowdService, &userService)
		tree = append(tree, groupNode)
	}

	// 添加未分配群和未分配用户
	if level > 1 {
		unassignedGroup := createUnassignedGroupNode(level, &crowdService, &userService)
		tree = append(tree, unassignedGroup)
	}

	// 最外层节点
	rootNode := map[string]interface{}{
		"id":       "0",
		"label":    "全部",
		"children": tree,
	}

	return []map[string]interface{}{rootNode}
}

// 创建分组节点
func createGroupNode(group DimUserCrowdGroupEntity, level int, crowdService *CrowdService, userService *UserService) map[string]interface{} {
	groupNode := map[string]interface{}{
		"id":    strconv.Itoa(group.ID),
		"label": group.Name,
	}

	if level > 1 {
		var crowds []map[string]interface{}
		for _, crowd := range crowdService.GetUserCrowdsByGroupID(group.ID) {
			crowdNode := createCrowdNode(crowd, group.ID, level, userService)
			crowds = append(crowds, crowdNode)
		}
		groupNode["children"] = crowds
	}

	return groupNode
}

// 创建用户群节点
func createCrowdNode(crowd DimUserCrowdEntity, groupID int, level int, userService *UserService) map[string]interface{} {
	crowdNode := map[string]interface{}{
		"id":    fmt.Sprintf("%d-%d", groupID, crowd.ID),
		"label": crowd.Name,
	}

	if level > 2 {
		var users []map[string]interface{}
		for _, user := range userService.GetUsersByCrowdID(crowd.ID) {
			users = append(users, map[string]interface{}{
				"id":    fmt.Sprintf("%d-%d-%d", groupID, crowd.ID, user.ID),
				"label": user.UserName,
			})
		}
		crowdNode["children"] = users
	}

	return crowdNode
}

// 创建未分配群节点
func createUnassignedGroupNode(level int, crowdService *CrowdService, userService *UserService) map[string]interface{} {
	unassignedGroup := map[string]interface{}{
		"id":    "00",
		"label": "未分配群",
	}

	var unassignedCrowds []map[string]interface{}
	for _, crowd := range crowdService.GetUserCrowdsNotInGroup() {
		crowdNode := createCrowdNode(crowd, 0, level, userService)
		unassignedCrowds = append(unassignedCrowds, crowdNode)
	}

	if level > 2 {
		var unassignedUsers []map[string]interface{}
		for _, user := range userService.GetUsersNotInAnyCrowd() {
			unassignedUsers = append(unassignedUsers, map[string]interface{}{
				"id":    fmt.Sprintf("0-0-%d", user.ID),
				"label": user.UserName,
			})
		}
		unassignedCrowds = append(unassignedCrowds, map[string]interface{}{
			"id":       "0-0",
			"label":    "未分配用户",
			"children": unassignedUsers,
		})
	}
	unassignedGroup["children"] = unassignedCrowds

	return unassignedGroup
}

func TestGetUserTree(t *testing.T) {
	tree := getUserTree(3)

	marshal, _ := json.Marshal(tree)
	if err := os.WriteFile("user_tree.json", marshal, os.ModePerm); err != nil {
		fmt.Println(err)
		return
	}
}

func TestGetGroupTree(t *testing.T) {

	mysql := setupMysql()
	log := setupLog()

	groupService := NewDimUserCrowdGroupSvc(log, mysql)

	tree, err := groupService.GetGroupTree(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	marshal, _ := json.Marshal(tree)
	if err := os.WriteFile("group_tree.json", marshal, os.ModePerm); err != nil {
		fmt.Println(err)
		return
	}

}
