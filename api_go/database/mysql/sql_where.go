package mysql

import (
	. "api/util/hit"
	"api/util/log"
	_string "api/util/string"
	"errors"
	"strings"
)

type sqlWhere struct {
	tableAlias 		string
	tableName 		string
	wheres		    []string
	values 		    []interface{}
	tags 			[]string
}

func (_where *sqlWhere) generateWhereIn(isSubQuery bool, boolean string, isNot bool, column string, value interface{}) {
	not := If(isNot, "not", "").(string)
	if closure, ok := value.(func(Table)); ok {
		subTable := newTable("", _where.tableName, nil , isSubQuery)
		closure(subTable)
		subTable.getDql().generateSQL(isSubQuery, subTable.getWhere())

		_where.wheres = append(_where.wheres, " " + _string.TrimQuotes(boolean) + " " +
			addPrefixTableName(_where.tableAlias, _where.tableName, column, _where.tags, isSubQuery) + " " + not + " in (" + subTable.getDql().sql + ") ")

		_where.values = append(_where.values, append(subTable.getDql().values)...)
	}else if v, ok := value.([]interface{}); ok {
		_where.wheres = append(_where.wheres, " " + _string.TrimQuotes(boolean) + " " +
			addPrefixTableName(_where.tableAlias, _where.tableName, column, _where.tags, isSubQuery) + " " + not + " in " + "(?" +
			strings.Repeat(",?", len(v)-1) + ") ")

		_where.values = append(_where.values, append(v)...)
	} else {
		log.Error(errors.New("whereIn value error"))
	}
}

func (_where *sqlWhere) generateWhere(isSubQuery bool, boolean string, column string, operator string, value interface{}) {
	if closure, ok := value.(func(Table)); ok {

		subTable := newTable("", "", nil , true)
		closure(subTable)

		subTable.getDql().generateSQL(true, subTable.getWhere())

		if subTable.getDql().selectSql == "" {
			_where.wheres = append(_where.wheres, boolean + " ( " + subTable.getDql().sql + ") ")
		} else {
			_where.wheres = append(_where.wheres, " " + boolean + " " + addPrefixTableName(_where.tableAlias, _where.tableName, column, _where.tags, isSubQuery) + " " +
				_string.TrimQuotes(operator) + " (" + subTable.getDql().sql + ") ")
		}

		_where.values = append(_where.values, append(subTable.getDql().values)...)
	} else {
		if len(_where.wheres) == 0 && _where.tableName == "" {
			boolean = ""
		}

		_where.wheres = append(_where.wheres, " " + boolean + " " +  addPrefixTableName(_where.tableAlias, _where.tableName, column, _where.tags, isSubQuery) + " " + _string.TrimQuotes(operator) + " ? ")
		_where.values = append(_where.values, value)
	}
}

func (_where *sqlWhere) generateWhereRaw(sql string) {
	_where.wheres = append(_where.wheres,  " " + sql + " ")
}

func (_where *sqlWhere) generateWhereBetween(isNot bool, boolean string, column string, values []interface{}) {
	not := If(isNot, "not", "").(string)
	_where.wheres = append(_where.wheres, " " + _string.TrimQuotes(boolean) + " " +
		addPrefixTableName(_where.tableAlias, _where.tableName, column, _where.tags, false) + " " + not + " between ? and ? ")

	_where.values = append(_where.values, append(values)...)
}


func (_where *sqlWhere) generateWhereColumne(boolean string, column1 string, operator string, column2 string) {
	_where.wheres = append(_where.wheres, " " + boolean + " " + addPrefixTableName(_where.tableAlias, _where.tableName, column1, _where.tags, false) +
		" " + _string.TrimQuotes(operator) + " " + addPrefixTableName(_where.tableAlias, _where.tableName, column2, _where.tags, false) + " ")
}

func (_where *sqlWhere) generateWhereNull(isNot bool, column string) {
	not := If(isNot, "not", "").(string)
	_where.wheres = append(_where.wheres,  " is " + not + " null " + addPrefixTableName(_where.tableAlias, _where.tableName, column, _where.tags, false) + " ")
}