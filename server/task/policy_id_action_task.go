package task

import (
	"bufio"
	"bytes"
	"fcas_server/global"
	"fcas_server/model/policy"
	policyService "fcas_server/service/policy"
	"fmt"
	"github.com/jlaffaye/ftp"
	"go.uber.org/zap"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var dimControlPolicyActionService = policyService.DimControlPolicyActionService{}

func RunPolicyActionIdTask() {
	// 登录ftp、切换目录、获取上报文件列表
	ftpClient, files, fail := loginFtpChangeDirAndGetFiles()
	if fail {
		return
	}
	defer func(ftpClient *ftp.ServerConn) {
		err := ftpClient.Quit()
		if err != nil {
			global.Log.Error("ftpClient.Quit fail", zap.Error(err))
			return
		}
	}(ftpClient)

	// 处理各个上报文件
	for _, file := range files {
		fileName := file.Name

		result, notOk := readFileAndRename(fileName, ftpClient)
		if notOk {
			continue
		}

		for i := range result {
			err := dimControlPolicyActionService.SaveOrUpdate(setEntity(result[i]))
			if err != nil {
				continue
			}
		}
	}

}

func loginFtpChangeDirAndGetFiles() (*ftp.ServerConn, []*ftp.Entry, bool) {
	host := global.CONFIG.Policy.Ftp.IP
	port := global.CONFIG.Policy.Ftp.Port
	user := global.CONFIG.Policy.Ftp.UserName
	password := global.CONFIG.Policy.Ftp.UserPwd
	path := global.CONFIG.Policy.Ftp.Path

	// 1 获取FTP连接
	ftpClient, err := ftp.Dial(host+":"+port, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		global.Log.Error(fmt.Sprintf("连接%s:%s失败, err: %s", host, port, err))
		return nil, nil, true
	}

	// 2 登录Ftp客户端
	err = ftpClient.Login(user, password)
	if err != nil {
		global.Log.Error(fmt.Sprintf("ftp登录失败：%s %s,  err: %s", user, password, err))
		return nil, nil, true
	}

	global.Log.Debug("ftp登录成功!")

	// 3 切换目录
	err = ftpClient.ChangeDir(path)
	if err != nil {
		global.Log.Error("ftp 切换目录失败", zap.Error(err))
		return nil, nil, true
	}

	// 4 查看路径下的文件
	files, listErr := ftpClient.List(".")
	if listErr != nil {
		global.Log.Error(fmt.Sprintf("ftp读取路径上的文件失败：path=%s ,  err: %s", path, err))
		return nil, nil, true
	}
	return ftpClient, files, false
}

func readFileAndRename(fileName string, ftpClient *ftp.ServerConn) ([]string, bool) {
	// 1 检查文件扩展名是否是.txt
	if filepath.Ext(fileName) != ".txt" {
		global.Log.Debug(fmt.Sprintf("文件 %s 不是 .txt 文件， 跳过处理", fileName))
		return nil, true
	}

	// 2 读取文件内容
	reader, retrErr := ftpClient.Retr(fileName)
	if retrErr != nil {
		global.Log.Error("读取FTP文件失败：", zap.Error(retrErr))
		return nil, true
	}

	// 3 读取所有内容到内存中（文件不太大）
	buf := new(bytes.Buffer)
	if _, readErr := buf.ReadFrom(reader); readErr != nil {
		global.Log.Error("读取FTP文件中的内容失败：", zap.Error(readErr))
		return nil, true
	}

	err := reader.Close()
	if err != nil {
		global.Log.Error("reader.Close() fail", zap.Error(err))
	}

	// 4 按照行分割
	scanner := bufio.NewScanner(strings.NewReader(buf.String()))

	// 5 移除可能的空行
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" { // 跳过空行
			lines = append(lines, line)
		}
	}
	if err = scanner.Err(); err != nil {
		global.Log.Error("scanner 读取文件行 失败", zap.Error(err))
		return nil, true
	}

	// 6 重命名文件
	newFileName := fileName + ".end"
	if renameErr := ftpClient.Rename(fileName, newFileName); renameErr != nil {
		global.Log.Error("重命名文件失败", zap.Error(renameErr))
		return nil, true
	}
	return lines, false
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
