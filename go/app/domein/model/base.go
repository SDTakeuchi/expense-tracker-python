package models

import (
    "time"
    "net/http"
    "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type BaseModel struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type models interface {
	GenUuid()
}

func GenUuid(m model.BaseModel) (uuid string) {
	for {
		uuid = ksuid.New()
		f, _ := m.FindByID(uuid)
		if f == nil {
			return 
		}
	}
}

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
