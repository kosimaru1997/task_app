package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB はMySQLの接続設定を表します。
type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// NewDB はDBの新しいインスタンスを返します。
func NewDB() *DB {
	return &DB{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "password",
		Name:     "task_app",
	}
}

// ConnectDB はMySQLに接続します。
func ConnectDB() (*gorm.DB, error) {
	dbConfig := NewDB()
	dsn := dbConfig.User + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
