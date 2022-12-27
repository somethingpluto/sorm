package test

import (
	"fmt"
	"sorm/v3/clause"
	"testing"
)

func TestSelect(t *testing.T) {
	var cla clause.Clause
	cla.Set(clause.LIMIT, 3)
	cla.Set(clause.SELECT, "user", []string{"*"})
	cla.Set(clause.WHERE, "name = ?", "tom")
	cla.Set(clause.ORDERBY, "age", "desc")
	sql, vars := cla.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)
	fmt.Printf("%s\n", sql)
	fmt.Printf("%v\n", vars)
}
