package bd

import (
    "fmt"
    "os"
    "time"
    "reservation_api/domain"
    log "github.com/sirupsen/logrus"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&tls=false",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

    var err error
    retries := 10
    for i := 0; i < retries; i++ {
        DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
            Logger: logger.Default.LogMode(logger.Info),
        })
        if err == nil {
            log.Info("Database connection established")
            return
        }
        log.Warnf("Database not ready, retrying in 2s... (%d/%d)", i+1, retries)
        time.Sleep(2 * time.Second)
    }

    log.Fatal("Could not connect to the database:", err)
}

// Exported function for migrations
func StartDbEngine() {
    DB.AutoMigrate(&domain.Reservation{})
    log.Info("Finishing Migration Database Tables")
}
