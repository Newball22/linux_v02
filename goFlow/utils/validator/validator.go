package validator

import (
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTans "github.com/go-playground/validator/v10/translations/zh"
	"goFlow/utils/errmsg"
	"reflect"
)

func Validate(data interface{}) (string, int) {
	/*本来这里需要一步对data进行断言，因为传进来知道是结构体，所以忽略*/

	validate := validator.New()

	uni := unTrans.New(zh_Hans_CN.New()) //转换成汉语中文
	tans, _ := uni.GetTranslator("zh_Hans_CN")

	//对汉语翻译进行注册
	err := zhTans.RegisterDefaultTranslations(validate, tans)
	if err != nil {
		fmt.Println("注册失败:", err)
	}

	//反射label标签
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(tans), errmsg.ERROR //返回错误信息
		}
	}
	return "", errmsg.SUCCESS
}
