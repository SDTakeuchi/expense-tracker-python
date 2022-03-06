package infra

import (
	"app/conf"
	"app/domein/models"
	"fmt"
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
		conf.Config.DbType,
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			conf.Config.DbHost,
			conf.Config.DbPort,
			conf.Config.DbUser,
			conf.Config.DbName,
			conf.Config.DbPassword,
			conf.Config.DbSslmode))
	if err != nil {
        panic(err.Error())
    }
    fmt.Println("db connected: ", &db)
    db.Set("gorm:table_options", "ENGINE=InnoDB")
    db.AutoMigrate(&models.Item{})
    db.LogMode(true)
}
