package sorm

import (
	"database/sql"
	"sorm/v0/log"
	"sorm/v0/session"
)

type Engine struct {
	db *sql.DB
}

// NewEngine
// @Description: 创建Engine
// @param driver
// @param source
// @return e
// @return err
//
func NewEngine(driver string, source string) (e *Engine, err error) {
	// 连接数据库
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	// 检查数据库是否连通
	err = db.Ping()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	e = &Engine{db: db}
	return e, nil
}

// Close
// @Description: 关闭数据库
// @receiver e
//
func (e *Engine) Close() {
	err := e.db.Close()
	if err != nil {
		log.Error(err)
		return
	}
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db)
}
