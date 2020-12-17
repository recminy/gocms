package requests

import (
	"cms/pkg/model"
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"io/ioutil"
	"net/http"
	"strings"
)

//应用初始化时候自动加载自定义验证规则
func Init()  {
	govalidator.AddCustomRule("unique", func(field string, rule string, message string, value interface{}) error {
		tableFieldPair := strings.Split(strings.TrimLeft(rule, "unique:"), ",")
		if len(tableFieldPair) != 2 {
			panic("unique规则错误：格式unique:tableName,FieldName")
		}
		tableName := tableFieldPair[0]
		fieldName := tableFieldPair[1]

		var count int64
		value = value.(string)
		model.DB.Table(tableName).Where(fieldName + "= ? ", value.(string)).Count(&count)

		if count > 0 {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%v 已被使用", value)
		}
		return nil
	})
}

func FormBody(r *http.Request) string {
	s, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "{}"
	}
	return string(s)
}

func FormByteBody(r *http.Request) []byte {
	s := FormBody(r)
	return []byte(s)
}
