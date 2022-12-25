package session

import (
	"database/sql"
	"sorm/v1/dialect"
	"sorm/v1/log"
	"sorm/v1/schema"
	"strings"
)

// Session
// @Description: 数据库会话连接
//
type Session struct {
	db       *sql.DB
	dialect  dialect.Dialect
	refTable *schema.Schema
	sql      strings.Builder
	sqlVars  []interface{}
}

// New
// @Description: 创建会话
// @param db
// @param dialect
// @return *Session
//
func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{db: db, dialect: dialect}
}

// Clear
// @Description: 清除会话中sql语句 参数
// @receiver s
//
func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

// DB
// @Description: 获取db
// @receiver s
// @return *sql.DB
//
func (s *Session) DB() *sql.DB {
	return s.db
}

// Raw
// @Description: 生成sql语句
// @receiver s
// @param sql
// @param values
// @return *Session
//
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

// Exec
// @Description: 执行sql语句
// @receiver s
// @return result
// @return err
//
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	result, err = s.DB().Exec(s.sql.String(), s.sqlVars...)
	if err != nil {
		log.Error(err)
	}
	return result, nil
}

// QueryRow
// @Description: 查询单条记录
// @receiver s
// @return *sql.Row
//
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

// QueryRows
// @Description: 查询多条记录
// @receiver s
// @return rows
// @return err
//
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	rows, err = s.DB().Query(s.sql.String(), s.sqlVars...)
	if err != nil {
		log.Error(err)
	}
	return rows, nil
}
