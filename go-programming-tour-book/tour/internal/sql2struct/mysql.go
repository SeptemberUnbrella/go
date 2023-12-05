package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DbModel struct {
	DBEngine *sql.DB
	DbInfo   *DBInfo
}

type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

func NewDBModel(info *DBInfo) *DbModel {
	return &DbModel{DbInfo: info} //DBModel是整个数据库连接的核心对象

}

func (m *DbModel) Connect() error {
	var err error
	s := "%s:%s@tcp(%s)/information_schema? " + "charset=%s&parseTime=True&loc=Local"
	dsn := fmt.Sprint(s,
		m.DbInfo.UserName,
		m.DbInfo.Password,
		m.DbInfo.Host,
		m.DbInfo.Charset,
	)
	m.DBEngine, err = sql.Open(m.DbInfo.DBType, dsn)
	if err != nil {
		return err
	}
	return nil
}

func (m *DbModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := "SELECT COLUMN_NAME,DATA_TYpe,COLUMNn_key, " +
		"IS_NULLABLE,COLUMN_TYPE,COLUMN_COMMENT" +
		"FROM COLUMNS WHERE TABLE_SCHEMA= ? and TABLE_NAME = ?"
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("没有数据")
	}
	defer rows.Close()
	var columns []*TableColumn
	for rows.Next() {
		var column TableColumn
		err := rows.Scan(&column.ColumnName, &column.DataType,
			&column.ColumnKey, &column.IsNullable, &column.ColumnType,
			&column.ColumnComment)

		if err != nil {
			return nil, err
		}
		columns = append(columns, &column)
	}
	return columns, nil
}

func main() {

}
