package valid

import (
	"reflect"

	"github.com/go-playground/validator"
)

// GetValidMsg 返回结构体中的msg参数 (参数绑定验证)
func GetValidMsg(err error, obj interface{}) string {
	//使用的时候需要传obj的指针
	objType := reflect.TypeOf(obj)
	//将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		//断言成功
		for _, e := range errs {
			//循环每一个错误信息
			//根据报错字段名获取结构体的具体字段
			if f, exits := objType.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}
