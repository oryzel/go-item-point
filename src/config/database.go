package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type DB struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func (db *DB) Connect() *gorm.DB {
	dsn := "mitraruma:mitraruma@tcp(127.0.0.1:33060)/mitraruma?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := fmt.Sprintf(db.Username,":",db.Password,"@tcp(",db.Host,":"db.Port)/db.Name,"?charset=utf8mb4&parseTime=True&loc=Local")
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot Connect Database")
	}
	return conn
}
