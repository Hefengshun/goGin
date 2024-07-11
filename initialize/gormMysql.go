package initialize

import (
	"database/sql"
	"fmt"
	"ginDemo/config"
	"ginDemo/models/demo"
	"ginDemo/models/system"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB {
	//前提是你要先在本机用Navicat创建一个名为go_db的数据库
	configMySql := config.MySql{
		Host:     "118.89.198.69",
		Port:     "3306",
		Username: "root",
		Password: "123456",
		Database: "goGin",
		Charset:  "utf8",
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		configMySql.Username,
		configMySql.Password,
		configMySql.Host,
		configMySql.Port,
		configMySql.Database,
		configMySql.Charset)

	// 先连接到 MySQL 服务器，不指定数据库
	dsnWithoutDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=%s&parseTime=true",
		configMySql.Username,
		configMySql.Password,
		configMySql.Host,
		configMySql.Port,
		configMySql.Charset)

	sqlDB, err := sql.Open("mysql", dsnWithoutDB)
	if err != nil {
		log.Fatalf("Error connecting to MySQL server: %v", err)
	}
	defer sqlDB.Close()

	// 检查数据库是否存在
	_, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", configMySql.Database))
	if err != nil {
		log.Fatalf("Error creating database: %v", err)
	}

	// 使用 GORM 连接到指定的数据库 //这里 gorm.Open()函数与之前版本的不一样，大家注意查看官方最新gorm版本的用法
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error to Db connection, err: " + err.Error())
	}
	//这个是gorm自动创建数据表的函数。它会自动在数据库中创建一个名为users的数据表
	_ = db.AutoMigrate(&system.SysUser{}, &demo.SysDemo{}) // 可以使用逗号分割 根据表的结构体创建多个表
	return db
}
