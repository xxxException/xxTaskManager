package common

import (
	"errors"
	"github.com/kataras/iris"
	"reflect"
	"strconv"
)

type GlobalStruct struct {
	PageSize int
}

var Global GlobalStruct

func GetField(dataModel interface{}, ctx iris.Context) (interface{}, error) {
	//获得类型
	t := reflect.TypeOf(dataModel)

	//如果是指针,则改为指向的结构体
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, errors.New("process url into dataModel fail: dataModel is not struct")
	}

	//初始化结构体
	model := reflect.New(t).Elem()

	//获取字段的数量
	fieldCount := t.NumField()

	var fieldName string
	var parameter interface{}
	var fieldType string
	var err error

	for idx := 0; idx < fieldCount; idx++ {
		field := t.Field(idx)
		fieldName = field.Name
		//获取字类型名
		fieldType = field.Type.Name()

		parameter = ctx.URLParam(fieldName)

		if fieldType == "string" {
			parameter = ctx.URLParam(fieldName)
			value := parameter.(string)
			model.FieldByName(fieldName).SetString(value)
		} else if fieldType == "int" {
			parameter, err = ctx.URLParamInt(fieldName)
			if err != nil {
				continue
			}
			value := parameter.(int64)
			model.FieldByName(fieldName).SetInt(value)
		} else if fieldType == "bool" {
			parameter, err = ctx.URLParamBool(fieldName)
			if err != nil {
				continue
			}
			value := parameter.(bool)
			model.FieldByName(fieldName).SetBool(value)
		} else if fieldType == "float64" {
			parameter, err = ctx.URLParamFloat64(fieldName)
			if err != nil {
				continue
			}
			value := parameter.(float64)
			model.FieldByName(fieldName).SetFloat(value)
		}
	}
	return model, nil
}

func StringsToInts(strs []string) ([]int, error) {
	ints := make([]int, len(strs), len(strs))

	for idx, str := range strs {
		ints[idx], _ = strconv.Atoi(str)
	}

	return ints, nil
}
