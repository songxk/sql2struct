package main

import (
	"database/sql"
	"fmt"
	"log"
)

type mySQLColumn struct {
	Database   string `gorm:"column:TABLE_SCHEMA"`
	Table      string `gorm:"column:TABLE_NAME"`
	Field      string `gorm:"column:COLUMN_NAME"`
	IsNullable string `gorm:"column:IS_NULLABLE"`
	DataType   string `gorm:"column:DATA_TYPE"`   //varchar
	ColumnType string `gorm:"column:COLUMN_TYPE"` //varchar(32)
	Comment    string `gorm:"column:COLUMN_COMMENT"`
}

func dbCon(dbc *dbConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", dbc.User, dbc.Password, dbc.Host, dbc.Database))
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}

func getColumns(db *sql.DB, dbc *dbConfig) ([]*mySQLColumn, error) {
	sql := fmt.Sprintf("select TABLE_SCHEMA,TABLE_NAME,COLUMN_NAME,IS_NULLABLE,DATA_TYPE,COLUMN_TYPE,COLUMN_COMMENT from INFORMATION_SCHEMA.COLUMNS where TABLE_SCHEMA = '%s' and TABLE_NAME = '%s'", dbc.Database, dbc.Table)
	columns := make([]*mySQLColumn, 0)
	var err error
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		row := &mySQLColumn{}
		err = rows.Scan(&row.Database, &row.Table, &row.Field, &row.IsNullable, &row.DataType, &row.ColumnType, &row.Comment)
		if err != nil {
			log.Fatal(err)
		}
		columns = append(columns, row)
	}
	return columns, nil
}
