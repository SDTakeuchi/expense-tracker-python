package models

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
    db  *gorm.DB
    err error
)

func Setup () {
	db, err = gorm.Open(
		setting.DatabaseSetting.Type,
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			setting.DatabaseSetting.Host,
			setting.DatabaseSetting.Port,
			setting.DatabaseSetting.User,
			setting.DatabaseSetting.Name,
			setting.DatabaseSetting.Password,
			setting.DatabaseSetting.Sslmode))
	if err != nil {
        panic(err.Error())
    }
    fmt.Println("db connected: ", &db)
    db.Set("gorm:table_options", "ENGINE=InnoDB")
    db.AutoMigrate(models.&Item{})
    db.LogMode(true)
}
