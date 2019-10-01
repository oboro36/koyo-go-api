package models

import (
	"github.com/jmoiron/sqlx"
	_"github.com/denisenkom/go-mssqldb"
)

var db *sqlx.DB
var err error

//Initialize Database
func InitDB(connectType string,dataSourceName string) {
	db, err = sqlx.Connect(connectType,dataSourceName)
		if err != nil {
			panic(err)
		}
}