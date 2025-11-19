package bd

import (
	"fmt"
	"os"
	"user_api/domain"
	usuarioClient "user_api/repository"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
    os.Getenv("DB_USER"),     // root
    os.Getenv("DB_PASSWORD"), // kaneki15..
    os.Getenv("DB_HOST"),     // mysql_service
    os.Getenv("DB_PORT"),     // 3306
    os.Getenv("DB_NAME"),     // hoteles
)


	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Error("Connection to database failed")
		log.Fatal(err)
	}

	log.Info("Database connection established")

	usuarioClient.DB = DB
}

func StartDbEngine() {
	DB.AutoMigrate(&domain.User{})
	log.Info("Finishing Migration Database Tables")
}
