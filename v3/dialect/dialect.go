package dialect

import "reflect"

var dialectMap = map[string]Dialect{}

// Dialect
// @Description: 屏蔽数据库差异
//
type Dialect interface {
	// DataTypeOf
	// @Description: Go内置数据类型与是数据库类型映射
	// @param typ
	// @return string
	//
	DataTypeOf(typ reflect.Value) string

	// TableExistSQL
	// @Description: 判断数据表是否存在
	// @param tableName
	// @return string
	// @return []interface{}
	//
	TableExistSQL(tableName string) (string, []interface{})
}

// RegisterDialect
// @Description: 注册
// @param name
// @param dialect
//
func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}

// GetDialect
// @Description: 获取连接数据库类型
// @param name
// @return dialect
// @return ok
//
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectMap[name]
	return dialect, ok
}
