package sorm

import (
	"database/sql"
	"sorm/v2/dialect"
	"sorm/v2/log"
	"sorm/v2/session"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver string, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return
	}

	e = &Engine{db: db, dialect: dial}

	log.Info("Connect database success")
	return e, nil
}

func (e *Engine) Close() {
	err := e.db.Close()
	if err != nil {
		log.Error(err)
	}
	log.Info("Close database success")
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db, e.dialect)
}
