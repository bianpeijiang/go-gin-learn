package models

import (
	"fmt"
	"github.com/bianpeijiang/go-gin-learn/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)
import _ "github.com/jinzhu/gorm/dialects/mysql"

var db *gorm.DB

type Model struct {
	ID        int64 `gorm:"primary_key" json:"id"`
	CreatedOn int   `json:"created_on"`
	UpdatedOn int   `json:"updated_on"`
}

func init() {
	var err error
	var dbName, user, password, host, tablePrefix string

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//配置命名策略
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,        //设置为true, 使用单数表名
			TablePrefix:   tablePrefix, //为所有表名添加前缀
			//NoLowerCase:   true,        //禁止将表名和字段名转换为蛇形命名
		},
		Logger: logger.Default.LogMode(logger.Info), //设置日志级别为Info
	})
	if err != nil {
		log.Println(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	//SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	//设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	//设置了可以重新使用连接的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
}
