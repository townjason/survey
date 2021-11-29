package mysql

import (
	"strings"
)

func addPrefixTableName(tableAlias string, tableName string, column string, tags []string, isSubQuery bool) string {
	if strings.Index(column, ".") >= 0 {
		return column
	}

	if tableAlias != "" {
		tableName = tableAlias
	}

	tn := ""
	if len(tags) > 0 {
		for _, v := range tags {
			if v == column {
				tn = "`"+ tableName + "`"+ "."
				break
			}
		}
	} else if !isSubQuery {
		tn = "`"+ tableName + "`"+ "."
	}

	return tn + column
}
