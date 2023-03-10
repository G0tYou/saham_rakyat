package config

import (
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func ConnectMysql() (*gorm.DB, error) {
	connString := os.Getenv("DB_USERNAME_MYSQL") + `:` + os.Getenv("DB_PASSWORD_MYSQL") + `@tcp(` + os.Getenv("DB_HOST_MYSQL") + `:3306)/` + os.Getenv("DB_NAME_MYSQL") + `?charset=utf8mb4&parseTime=True&loc=Local`

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		log.Println("Open connection failed:", err.Error())
	}

	return db, nil
}
