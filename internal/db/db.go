package db

import (
	"ChatBot/config"
	model "ChatBot/internal/model"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/ziutek/mymysql/godrv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MyDB struct {
	DB     *gorm.DB
	UserDB *gorm.DB
	ChatDB *gorm.DB
}

func (db *MyDB) InitDB() {
	dsn := config.DbUser + ":" + config.DbPassword +
		"@tcp(127.0.0.1:3306)/chatbot?charset=utf8mb4&parseTime=True&loc=Local"

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}
	sqlDB.SetMaxIdleConns(1000)
	sqlDB.SetMaxOpenConns(100000)
	sqlDB.SetConnMaxLifetime(-1)
	DB, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}
	db.DB = DB

	db.UserDB = db.DB
	db.UserDB.AutoMigrate(&model.User{})
	db.ChatDB = db.DB
	db.ChatDB.AutoMigrate(&model.ChatMessage{})
}
