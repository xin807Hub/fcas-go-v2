package object

import (
	"bufio"
	"fcas_server/global"
	"fcas_server/model/object"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

type AppClassifySvc struct {
	Log   *zap.Logger
	Mysql *gorm.DB
}

func NewAppClassifySvc(log *zap.Logger, mysql *gorm.DB) *AppClassifySvc {
	return &AppClassifySvc{
		Log:   log.Named("[AppClassify-应用分类与自定义]"),
		Mysql: mysql,
	}
}

func (svc AppClassifySvc) List(params object.AppClassifyListRequest) (output []*object.AppClassify, total int64, err error) {
	tx := svc.Mysql.Model(&object.AppClassify{})

	if params.AppTypeId != 0 {
		tx = tx.Where("app_type_id = ?", params.AppTypeId)
	}
	if params.AppId != 0 {
		tx = tx.Where("app_id = ?", params.AppId)
	}

	err = tx.Count(&total).Limit(params.Limit).Offset((params.Page - 1) * params.Limit).Find(&output).Error
	if err != nil {
		svc.Log.Error("获取列表信息失败", zap.Any("params", params), zap.Error(err))
		return nil, 0, err
	}

	return output, total, nil
}

func (svc AppClassifySvc) Import(fileHeader *multipart.FileHeader) error {
	open, err := fileHeader.Open()
	if err != nil {
		svc.Log.Error("打开文件失败", zap.Error(err))
		return err
	}
	defer open.Close()

	appM := make(map[string]string)
	scanner := bufio.NewScanner(open)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) != 6 {
			svc.Log.Error("导入数据行格式错误", zap.String("line", line))
			continue
		}
		// fields[2]: 0100001001
		appM[fmt.Sprintf("%s_%s", fields[2][0:2], fields[2][3:7])] = fmt.Sprintf("%s_%s", fields[5], fields[4]) // 01_0001: 通讯_QQ
	}
	if err := scanner.Err(); err != nil {
		svc.Log.Error("读取文件内容时发生错误", zap.Error(err))
		return err
	}

	apps := make([]object.AppClassify, 0, len(appM))
	for key, val := range appM {
		ids := strings.Split(key, "_")
		names := strings.Split(val, "_")

		appTypeId, _ := strconv.Atoi(ids[0])
		appId, _ := strconv.Atoi(fmt.Sprintf("%d%s", appTypeId, ids[1]))

		apps = append(apps, object.AppClassify{
			AppTypeId:   appTypeId,
			AppId:       appId,
			AppTypeName: names[0],
			AppName:     names[1],
		})
	}

	// 入库
	if err := svc.Mysql.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("1=1").Delete(&object.AppClassify{}).Error; err != nil {
			return fmt.Errorf("删除原始数据失败: %w", err)
		}

		if err := tx.Create(&apps).Error; err != nil {
			return fmt.Errorf("插入新数据失败: %w", err)
		}

		return nil
	}); err != nil {
		svc.Log.Error("导入数据到数据库时发生错误", zap.Error(err))
		return err
	}

	time.Sleep(time.Second)

	// 更新全局Map
	if err := svc.InitGlobalAppMap(); err != nil {
		svc.Log.Error("更新AppMap失败", zap.Error(err))
		return fmt.Errorf("更新AppTypeMap失败，请联系管理员重启系统")
	}

	if err := svc.InitGlobalAppTypeMap(); err != nil {
		svc.Log.Error("更新AppTypeMap失败", zap.Error(err))
		return fmt.Errorf("更新AppTypeMap失败，请联系管理员重启系统")
	}

	return nil
}

func (svc AppClassifySvc) InitGlobalAppMap() error {

	var apps []struct {
		AppId   int
		AppName string
	}
	if err := svc.Mysql.Table("dim_app_classify").Select("app_id", "app_name").Find(&apps).Error; err != nil {
		return err
	}

	appMap := make(map[int]string)
	for _, item := range apps {
		appMap[item.AppId] = item.AppName
	}
	global.AppMap = appMap

	return nil
}

func (svc AppClassifySvc) InitGlobalAppTypeMap() error {

	var appTypes []struct {
		AppTypeId   int
		AppTypeName string
	}
	if err := svc.Mysql.Table("dim_app_classify").Select("distinct app_type_id", "app_type_name").Find(&appTypes).Error; err != nil {
		return nil
	}

	appTypeMap := make(map[int]string, len(appTypes))
	for _, item := range appTypes {
		appTypeMap[item.AppTypeId] = item.AppTypeName
	}
	global.AppTypeMap = appTypeMap

	return nil
}
