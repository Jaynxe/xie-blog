package utils

import (
	"reflect"
	"strings"
	"unicode"

	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func BuildQuerySQL(tx *gorm.DB, query *model.QueryRequest, role ...string) *gorm.DB {
	var where []string
	var params []any
	if query.Email != "" {
		where = append(where, "email = ?")
		params = append(params, query.Email)
	}
	if query.Name != "" {
		where = append(where, "name = ?")
		params = append(params, query.Name)
	}

	if len(where) == 0 {
		return nil
	}

	whereStatement := strings.Join(where, " AND ")

	if len(role) > 0 {
		whereStatement += " AND role = ?"
		params = append(params, role[0])
	}
	return tx.Where(whereStatement, params...)
}
func BuildLoginSQL(tx *gorm.DB, query *model.UserLoginRequest, role ...string) *gorm.DB {
	var where []string
	var params []any
	if query.Name != "" {
		where = append(where, "name = ?")
		params = append(params, query.Name)
	}
	if len(where) == 0 {
		return nil
	}
	whereStatement := strings.Join(where, " AND ")

	if len(role) > 0 {
		whereStatement += " AND role = ?"
		params = append(params, role[0])
	}
	return tx.Where(whereStatement, params...)
}
func IgnoreStructCopy(to, from any, ignore string) {
	copier.CopyWithOption(to, from, copier.Option{
		IgnoreEmpty: true,
	})

	if ignore == "" {
		return
	}
	elem := reflect.Indirect(reflect.ValueOf(to))
	ignoreLower := strings.ToUpper(ignore[0:1]) + ignore[1:]
	for i := 0; i < elem.NumField(); i++ {
		current := elem.Field(i)
		if elem.Type().Field(i).Name == ignoreLower {
			current.Set(reflect.Zero(current.Type()))
			break
		}
	}
}

func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	var upper, lower, number int

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			lower++
		case unicode.IsUpper(char):
			upper++
		case unicode.IsNumber(char):
			number++
		}
	}
	return upper > 0 && lower > 0 && number > 0
}

type Option struct {
	model.PageRequest
	Debug bool //是否显示日志
}

// ComList 用于分页查询的通用方法
func ComList[T any](list []T, option Option) ([]T, error) {
	//默认数据库操作没有日志
	DB := global.GVB_DB
	if !option.Debug {
		//不生成日志
		DB = global.GVB_DB.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)})
	}
	//图片列表排序
	if option.Sort == "" {
		option.Sort = "id desc" //默认按ID降序排列(即最新的排在前面)
	}
	//总的条数
	count := DB.Select("id").Find(&list).RowsAffected
	// 分页查询
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	if option.Limit <= 0 {
		option.Limit = int(count)
	}
	err := DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	return list, err
}
