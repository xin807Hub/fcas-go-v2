package initialize

import (
	"fcas_server/global"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

func TransInit() {
	// 注册翻译器
	zhT := zh.New()
	uni := ut.New(zhT, zhT)
	trans, _ := uni.GetTranslator("zh")
	global.Trans = trans

	// 获取gin的校验器
	ginValidate := binding.Validator.Engine().(*validator.Validate)

	// 给gin校验器注册翻译器
	err := zhTranslations.RegisterDefaultTranslations(ginValidate, global.Trans)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
