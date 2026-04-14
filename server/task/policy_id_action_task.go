package task

import (
	"bufio"
	"fcas_server/global"
	"fcas_server/model/policy"
	policyService "fcas_server/service/policy"
	"fmt"
	"github.com/jlaffaye/ftp"
	"strconv"
	"strings"
	"time"
)

var dimControlPolicyActionService = policyService.DimControlPolicyActionService{}

func RunPolicyIdActionIdTask() {
	user := global.CONFIG.Policy.UserName
	password := global.CONFIG.Policy.UserPwd
	host := global.CONFIG.Policy.IP
	port := global.CONFIG.Policy.Port
	path := global.CONFIG.Policy.Path

	// 获取FTP连接
	client, err := ftp.Dial(host+":"+port, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		global.Log.Error(fmt.Sprintf("连接%s:%s失败, err: %s", host, port, err))
		return
	}

	// 登录Ftp客户端
	err = client.Login(user, password)
	if err != nil {
		global.Log.Error(fmt.Sprintf("ftp登录失败：%s %s,  err: %s", user, password, err))
		return
	}
	defer client.Quit()

	entries, err := client.List(path)
	if err != nil {
		global.Log.Error(fmt.Sprintf("ftp读取文件路径失败：%s ,  err: %s", path, err))
		return
	}
	for _, entry := range entries {
		// 读取文件
		reader, _ := client.Retr(entry.Name)
		defer reader.Close()

		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			line := scanner.Text()
			_ = dimControlPolicyActionService.SaveOrUpdate(setEntity(line))
		}
	}

}

func setEntity(line string) policy.DimControlPolicyAction {
	lines := strings.Split(line, "|")
	policyAction := policy.DimControlPolicyAction{}
	policyAction.PolicyId, _ = strconv.Atoi(lines[0])
	policyAction.VlanId, _ = strconv.Atoi(lines[1])
	policyAction.ShuntIp = lines[2]
	policyAction.UploadActionId, _ = strconv.Atoi(lines[3])
	policyAction.DownloadActionId, _ = strconv.Atoi(lines[4])
	policyAction.UploadDeviceId, _ = strconv.Atoi(lines[5])
	return policyAction
}
