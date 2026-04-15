package addr_set

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// 迁移旧地址
func TestTransferOldAddr(t *testing.T) {
	inputFilePath := "dim_user_info.sql"
	outputFilePath := "new_dim_user_info.sql"

	userInfos, err := ExtractSQLValues(inputFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	addrSet := NewAddrSet()

	var newUserInfos []UserInfo
	for _, user := range userInfos {
		addrs, availableAddrCount, err := addrSet.Add(user.IPAddress)
		if err != nil {
			fmt.Printf("用户ID：%d, 用户名称：%s, 原始地址：%s\n", user.ID, user.UserName, user.IPAddress)
			return
		}
		newUserInfos = append(newUserInfos, UserInfo{
			ID:           user.ID,
			UserName:     user.UserName,
			IPAddress:    addrs,
			IpaddressNum: availableAddrCount,
			UserType:     user.UserType,
			Remark:       user.Remark,
		})
	}

	insertSQL, err := GenerateInsertSQL("fcas_service.dim_user_info", newUserInfos)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := os.WriteFile(outputFilePath, []byte(insertSQL), os.ModePerm); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("迁移完成")
}

// UserInfo 用户信息结构体
type UserInfo struct {
	ID           int      // 用户ID
	UserName     string   `json:"user_name"` // 用户名称
	IPAddress    []string // ip地址段
	IpaddressNum string   // ip个数
	UserType     int      // 用户类型 0：正常 1：监测用户
	Remark       string   // 备注
}

// ExtractSQLValues 从SQL文件中提取INSERT语句中的VALUES内容
func ExtractSQLValues(filePath string) ([]UserInfo, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("打开文件失败: %v", err)
	}
	defer file.Close()

	// 创建正则表达式匹配VALUES内容
	valuesRegex := regexp.MustCompile(`VALUES\s*\((.*?)\);`)

	// 存储提取的结果
	var results []UserInfo

	// 逐行读取文件
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// 跳过注释行和空行
		if strings.HasPrefix(line, "--") || strings.TrimSpace(line) == "" {
			continue
		}

		// 匹配VALUES内容
		matches := valuesRegex.FindStringSubmatch(line)
		if len(matches) > 1 {
			// 提取括号内的内容
			valuesStr := matches[1]
			// 处理字符串中的引号，避免分割错误
			valuesStr = strings.ReplaceAll(valuesStr, "\\'", "\u0001")

			// 使用正则表达式匹配字段，先匹配所有被单引号包围的内容
			fields := strings.Split(valuesStr, ", ") // 分割字段，使用逗号作为分隔符

			// 将字符串转换为对应的类型
			id := 0
			userType := 0

			var err1, err2 error
			id, err1 = strconv.Atoi(fields[0])
			userType, err2 = strconv.Atoi(fields[3])
			if err1 != nil || err2 != nil {
				continue // 跳过无效的数据行
			}

			// 处理字符串字段，移除单引号
			userName := strings.Trim(fields[1], "'")
			ipAddress := strings.Trim(fields[2], "'")
			remark := strings.Trim(fields[4], "'")

			results = append(results, UserInfo{
				ID:        id,
				UserName:  userName,
				IPAddress: strings.Split(ipAddress, ","),
				UserType:  userType,
				Remark:    remark,
			})
		}
	}

	return results, nil
}

// GenerateInsertSQL 根据UserInfo数组生成MySQL INSERT语句
func GenerateInsertSQL(tableName string, users []UserInfo) (string, error) {
	if len(users) == 0 {
		return "", fmt.Errorf("用户列表不能为空")
	}

	var sql strings.Builder
	sql.WriteString(fmt.Sprintf("INSERT INTO %s (id, user_name, ip_address, ip_address_num, user_type, remark) VALUES ", tableName))

	for i, user := range users {
		if i > 0 {
			sql.WriteString(", ")
		}
		// 将IPAddress数组转为JSON格式
		ipJSON, err := json.Marshal(user.IPAddress)
		if err != nil {
			return "", fmt.Errorf("IP地址序列化失败: %v", err)
		}
		ipStr := string(ipJSON)
		// 转义字符串中的单引号
		userName := strings.ReplaceAll(user.UserName, "'", "\\'")
		remark := strings.ReplaceAll(user.Remark, "'", "\\'")

		sql.WriteString(fmt.Sprintf("(%d, '%s', '%s', '%s', %d, '%s')",
			user.ID, userName, ipStr, user.IpaddressNum, user.UserType, remark))
	}

	sql.WriteString(";")
	return sql.String(), nil
}
