package schema

import (
	"go/ast"
	"reflect"
	"sorm/v1/dialect"
)

// Field
// @Description: 结构体中的字段
//
type Field struct {
	// 字段名称
	Name string
	// 字段类型
	Type string
	// 字段标签
	Tag string
}

// Schema
// @Description: 经过解析的结构体
//
type Schema struct {
	Model interface{}
	// 结构体名称
	Name string
	// 结构体字段数组
	Fields []*Field
	// 字段名称数组
	FieldNames []string
	// 字段名称与字段结构体的映射
	fieldMap map[string]*Field
}

// GetField
// @Description: 根据字段名获取字段结构体
// @receiver schema
// @param name
// @return *Field
//
func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

func Parse(dest interface{}, d dialect.Dialect) *Schema {
	// 1.获取传入结构体值的类型
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	// 2.初始化该结构体schema
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}
	// 3.遍历结构体字段
	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		// 如果为不是匿名字段并且为对外暴露
		if !p.Anonymous && ast.IsExported(p.Name) {
			// 生成字段结构体
			filed := &Field{
				Name: p.Name,
				// 字段类型转换
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}

			// 获取key为sorm的标签
			v, ok := p.Tag.Lookup("sorm")
			if ok {
				filed.Tag = v
			}
			schema.Fields = append(schema.Fields, filed)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = filed
		}
	}
	return schema
}
