package test

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"net/http"
	"testing"
)

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *gorm.DB

// initDB
func initDB() {
	dsn := "root:123456@tcp(192.168.5.246:3306)/fcas_service?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

// ExportToExcel 函数用于将数据导出为Excel文件流
func ExportToExcel(headers []string, data []map[string]interface{}) (io.Reader, error) {
	// 创建一个新的Excel文件
	f := excelize.NewFile()

	// 创建一个工作表
	sheetName := "Sheet1"
	index, _ := f.NewSheet(sheetName)
	// 设置活动的工作表
	f.SetActiveSheet(index)

	// 写入表头
	for i, header := range headers {
		col := string(rune('A' + i))
		_ = f.SetCellValue(sheetName, fmt.Sprintf("%s1", col), header)
	}

	// 写入数据
	for rowIndex, item := range data {
		for colIndex, header := range headers {
			value := item[header]
			col := string(rune('A' + colIndex))
			_ = f.SetCellValue(sheetName, fmt.Sprintf("%s%d", col, rowIndex+2), value)
		}
	}

	// 创建一个缓冲区来存储Excel文件
	var buf bytes.Buffer

	// 将Excel文件写入缓冲区
	if err := f.Write(&buf); err != nil {
		return nil, err
	}

	// 返回文件流
	return &buf, nil
}

func TestExport(t *testing.T) {
	initDB()

	r := gin.Default()

	// 创建一个路由处理函数
	r.GET("/export", func(c *gin.Context) {
		// 定义表头
		headers := []string{"ID", "Name", "Age"}

		// 查询数据库中的所有用户
		var users []User
		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query data"})
			return
		}

		// 将数据库数据放到interface中
		var list interface{}
		list = users

		// 从interface中取出相应类型的数据
		dataList := list.([]User)

		// 将查询结果转换为 []map[string]interface{} 格式
		data := make([]map[string]interface{}, len(dataList))
		for i, user := range dataList {
			data[i] = map[string]interface{}{
				"ID":   user.ID,
				"Name": user.Name,
				"Age":  user.Age,
			}
		}

		// 调用公共方法生成Excel文件流
		fileStream, err := ExportToExcel(headers, data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate file"})
			return
		}

		// 设置HTTP响应头
		c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		c.Header("Content-Disposition", "attachment; filename=export.xlsx")
		c.Header("File-Name", "export.xlsx")
		c.Header("Content-Transfer-Encoding", "binary")

		// 将文件流写入HTTP响应
		if _, err := io.Copy(c.Writer, fileStream); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write file"})
			return
		}
	})

	// 启动服务 监听并在 0.0.0.0:8080 上启动服务
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
