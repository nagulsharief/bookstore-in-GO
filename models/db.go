package models

import (
    "log"
//    "database/sql"
//     _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func InitDB(dataSourceName string) {
    var err error
//    db, err = sql.Open("mysql", dataSourceName)
	 db, err = gorm.Open("mysql", dataSourceName)
    if err != nil {
        log.Panic(err)
    }
}