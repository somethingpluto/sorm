package test

import (
	"sorm/v3/clause"
	"sorm/v3/log"
	"testing"
)

func TestUpdateSQLBuild(t *testing.T) {
	var cla clause.Clause
	cla.Set(clause.UPDATE, "age", 30)
	cla.Set(clause.WHERE, "name", "tom")
	sql, vars := cla.Build(clause.UPDATE, clause.WHERE)
	log.Info(sql)
	log.Info(vars)
}
