package session

import (
	"errors"
	"reflect"
	"sorm/clause"
)

func (s *Session) Limit(values ...interface{}) *Session {
	s.clause.Set(clause.LIMIT, values...)
	return s
}

func (s *Session) Offset(values ...interface{}) *Session {
	s.clause.Set(clause.OFFSET, values...)
	return s
}

func (s *Session) Order(values ...interface{}) *Session {
	s.clause.Set(clause.ORDERBY, values...)
	return s
}

func (s *Session) Update(kv ...interface{}) (int64, error) {
	m, ok := kv[0].(map[string]interface{})
	if !ok {
		m = make(map[string]interface{})
		// 一次处理一个键值对 Age 18
		for i := 0; i < len(kv); i += 2 {
			m[kv[i].(string)] = kv[i+1]
		}
	}
	s.clause.Set(clause.UPDATE, s.refTable.Name, m)
	sql, vars := s.clause.Build(clause.UPDATE, clause.WHERE)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (s *Session) Where(desc string, args ...interface{}) *Session {
	var vars []interface{}
	vars = append(vars, desc)
	vars = append(vars, args...)
	s.clause.Set(clause.WHERE, vars...)
	return s
}

func (s *Session) Delete() (int64, error) {
	s.clause.Set(clause.DELETE, s.refTable.Name)
	sql, vars := s.clause.Build(clause.DELETE, clause.WHERE)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (s *Session) Find(values interface{}) error {
	destSlice := reflect.Indirect(reflect.ValueOf(values))
	destType := destSlice.Type().Elem()
	table := s.Model(reflect.New(destType).Elem().Interface()).RefTable()

	s.clause.Set(clause.SELECT, table.Name, table.FieldNames)
	sql, vars := s.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT, clause.OFFSET)
	rows, err := s.Raw(sql, vars...).QueryRows()
	if err != nil {
		return err
	}

	for rows.Next() {
		dest := reflect.New(destType).Elem()
		var values []interface{}
		for _, name := range table.FieldNames {
			values = append(values, dest.FieldByName(name).Addr().Interface())
		}
		err := rows.Scan(values...)
		if err != nil {
			return err
		}
		destSlice.Set(reflect.Append(destSlice, dest))
	}
	return rows.Close()
}

func (s *Session) First(value interface{}) error {
	dest := reflect.Indirect(reflect.ValueOf(value))
	destSlice := reflect.New(reflect.SliceOf(dest.Type())).Elem()
	err := s.Limit(1).Find(destSlice.Addr().Interface())
	if err != nil {
		return err
	}
	if destSlice.Len() == 0 {
		return errors.New("NOT FOUND")
	}
	dest.Set(destSlice.Index(0))
	return nil
}

func (s *Session) Insert(values ...interface{}) (int64, error) {

	s.CallMethod(BEFORE_INSERT, values...)
	recordValues := make([]interface{}, 0)
	for _, value := range values {
		table := s.RefTable()
		s.clause.Set(clause.INSERT, table.Name, table.FieldNames)
		recordValues = append(recordValues, table.RecordValues(value))
	}

	s.clause.Set(clause.VALUES, recordValues...)
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	s.CallMethod(AFTER_INSERT, values...)
	return result.RowsAffected()
}
