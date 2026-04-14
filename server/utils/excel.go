package utils

import (
	"bytes"
	"fcas_server/global"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
)

func GetExcelFileStream(fileName string, headers []string, dataList [][]interface{}) (io.Reader, error) {
	// 创建一个缓冲区来存储Excel
	var buf bytes.Buffer

	// 创建excel文件
	f := excelize.NewFile()
	//defer func() {
	//	if err := f.Close(); err != nil {
	//		global.Log.Error(err.Error())
	//	}
	//}()

	// 创建新的工作表
	var sheetName = "Sheet1"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		global.Log.Error(err.Error())
		return &buf, err
	}
	// 设置下载地址以及文件名
	f.Path = fileName
	// 设置活动工作表
	f.SetActiveSheet(index)

	// 写入表头
	for i, header := range headers {
		// 将表头转成excel列名
		colName, _ := excelize.CoordinatesToCellName(i+1, 1) // 第一行，从第一列开始累计
		_ = f.SetCellValue(sheetName, colName, header)
	}

	// 写入数据
	for rowIndex, row := range dataList {
		for colIndex, cell := range row {
			_ = f.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(rune('A'+colIndex)), rowIndex+2), cell)
		}
	}

	// 将Excel文件写入缓冲区
	if err := f.Write(&buf); err != nil {
		return nil, err
	}

	return &buf, nil
}

func ExportToExcel(fields []string, headers []string, data []map[string]interface{}) ([]byte, error) {
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
		for colIndex, field := range fields {
			value := item[field]
			col := string(rune('A' + colIndex))
			_ = f.SetCellValue(sheetName, fmt.Sprintf("%s%d", col, rowIndex+2), value)
		}
	}

	// 创建一个缓冲区来存储Excel
	var buf bytes.Buffer

	// 将Excel文件写入缓冲区
	if err := f.Write(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
