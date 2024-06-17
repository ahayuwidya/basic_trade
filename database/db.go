package database

import (
	"basic_trade/models/entity"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	dsn_config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	// connect to DB
	db, err = gorm.Open(mysql.Open(dsn_config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}
	fmt.Println("Successfully connected to database.")

	// migrate to DB
	db.AutoMigrate(&entity.Admin{}, &entity.Product{}, &entity.Variant{})
	fmt.Println("Database migrated successfully.")
	if db == nil {
		log.Fatal("database connection is nil")
	}
}

func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("database is nil")
	}
	fmt.Println("Succcessfully get database")
	return db
}
