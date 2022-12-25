package dialect

import (
	"fmt"
	"reflect"
	"time"
)

type mysql struct {
}

var _Mysql = (*mysql)(nil)

// init
// @Description: 注册Mysql
//
func init() {
	RegisterDialect("mysql", &mysql{})
}

// DataTypeOf
// @Description: Mysql字段类型与Go内置类型映射
// @receiver m
// @param typ
// @return string
//
func (m *mysql) DataTypeOf(typ reflect.Value) string {
	switch typ.Kind() {
	case reflect.Bool: // 布尔值
		return "bool"
	case reflect.Int8, reflect.Uint8: // 整形
		return "tinyint"
	case reflect.Int16, reflect.Uint16:
		return "smallint"
	case reflect.Int32, reflect.Int, reflect.Uint32:
		return "integer"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Float32, reflect.Float64: // 浮点型
		return "double"
	case reflect.String: // 字符串
		return "varchar(255)"
	case reflect.Struct:
		_, ok := typ.Interface().(time.Time)
		if ok {
			return "datetime"
		}
	case reflect.Array, reflect.Slice:
		return "blob"
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind()))
}

// TableExistSQL
// @Description: 判断数据表是否存在
// @receiver m
// @param tableName
// @return string
// @return []interface{}
//
func (m *mysql) TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	return "SELECT * FROM information_schema.TABLES where TABLE_NAME = ?", args
}
