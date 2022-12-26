package session

import (
	"fmt"
	"reflect"
	"sorm/v2/log"
	"sorm/v2/schema"
	"strings"
)

func (s *Session) Model(value interface{}) *Session {
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

// RefTable
// @Description: 获取会话操作的数据表
// @receiver s
// @return *schema.Schema
//
func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("table not create")
	}
	return s.refTable
}

// CreateTable
// @Description: 根据table 创建数据表
// @receiver s
// @return error
//
func (s *Session) CreateTable() error {
	table := s.RefTable()
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s)", table.Name, desc)).Exec()
	return err
}

// DropTable
// @Description: 根据table删除数据表
// @receiver s
// @return error
//
func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", s.RefTable().Name)).Exec()
	return err
}

// HasTable
// @Description: 判断数据表是否存在
// @receiver s
// @return bool
//
func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExistSQL(s.refTable.Name)
	queryRow := s.Raw(sql, values...).QueryRow()
	var temp string
	_ = queryRow.Scan(&temp)
	return temp == s.RefTable().Name
}
