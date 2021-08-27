package validator

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	chTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
	"testing"
)

/**
Validator 是基于 tag（标记）实现结构体和单个字段的值验证库，它包含以下功能：

使用验证 tag（标记）或自定义验证器进行跨字段和跨结构体验证。
关于 slice、数组和 map，允许验证多维字段的任何或所有级别。
能够深入 map 键和值进行验证。
通过在验证之前确定接口的基础类型来处理类型接口。
处理自定义字段类型（如 sql 驱动程序 Valuer）。
别名验证标记，它允许将多个验证映射到单个标记，以便更轻松地定义结构体上的验证。
提取自定义的字段名称，例如，可以指定在验证时提取 JSON 名称，并在生成的 FieldError 中使用该名称。
Web 框架 gin 的默认验证器。
*/
func Test_Validator_01(t *testing.T) {
	validate := validator.New()
	//验证变量
	email := "admin*admin.com"
	//email := ""
	/**
	Var 方法使用 tag（标记）验证方式验证单个变量。
	func (*validator.Validate).Var(field interface{}, tag string) error
	它接收一个 interface{} 空接口类型的 field 和一个 string 类型的 tag，返回传递的非法值得无效验证错误，否则将 nil 或 ValidationErrors 作为错误。如果错误不是 nil，则需要断言错误去访问错误数组，
	*/
	err := validate.Var(email, "required,email")

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		fmt.Println(validationErrors)
		return
	}

}

func Test_Validator_02(t *testing.T) {
	/**
	验证结构体
	*/
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return "j"
		}
		return name
	})
	type User struct {
		ID     int64  `json:"id" validate:"gt=0"`
		Name   string `json:"name" validate:"required"`
		Gender string `json:"gender" validate:"required,oneof=man woman"`
		Age    uint8  `json:"age" validate:"required,gte=0,lte=130"`
		Email  string `json:"email" validate:"required,email"`
	}
	user := &User{
		ID:     1,
		Name:   "liam",
		Gender: "body",
		Age:    135,
		Email:  "zyy@qq.com",
	}

	err := validate.Struct(user)
	zh := zh.New()
	uni := ut.New(zh)
	trans, _ := uni.GetTranslator("zh")
	_ = chTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		fmt.Println(validationErrors.Translate(trans))
		return
	}

}
