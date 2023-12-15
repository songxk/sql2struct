package main

import (
	"database/sql"
	"fmt"
	"strings"
)

func genFromTable(dbc *dbConfig, db *sql.DB) (string, error) {
	columns, err := getColumns(db, dbc)
	if err != nil {
		return "", err
	}
	table := strings.TrimLeft(dbc.Table, dbc.TablePrefix)
	tSegs := strings.Split(table, "_")
	tCamelCase := ""
	for _, v := range tSegs {
		tCamelCase = tCamelCase + ucUpper(v)
	}
	str := fmt.Sprintf("type %s%s struct {\n", ucUpper(tCamelCase), dbc.ModelSuffix)
	for _, v := range columns {
		segs := strings.Split(v.Field, "_")
		camelCase := ""
		for _, v := range segs {
			camelCase = camelCase + ucUpper(v)
		}
		str = str + fmt.Sprintf("  %s %s `gorm:\"column:%s\" json:\"%s\"` // %s\n", camelCase, getGoType(v.DataType), v.Field, ucLower(camelCase), v.Comment)
	}
	str = str + "}"
	return str, nil
}

func getGoType(mysqlType string) string {
	if mysqlType == "int" || mysqlType == "tinyint" {
		return "int"
	} else if mysqlType == "varchar" {
		return "string"
	} else if mysqlType == "bigint" {
		return "int64"
	}
	return "string"
}
