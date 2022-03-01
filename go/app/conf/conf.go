package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
)

type ConfigList struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  int
	WriteTimeout int
	DbType       string
	DbUser       string
	DbPassword   string
	DbHost       string
	DbPort       int
	DbName       string
	DbSslmode    string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("../prd.config.ini")
	if err != nil {
		fmt.Printf("Error Reading prd config file: %v\n", err)
		fmt.Println("Now trying to read local config file...")
		cfg, err = ini.Load("../local.config.ini")
		if err != nil {
			fmt.Println("Neither prd nor local config file found.")
			log.Fatal(err)
		}
	}

	Config = ConfigList{
		RunMode:      cfg.Section("web").Key("RunMode").MustString("debug"),
		HttpPort:     cfg.Section("web").Key("HttpPort").String(),
		ReadTimeout:  cfg.Section("web").Key("ReadTimeout").MustInt(100),
		WriteTimeout: cfg.Section("web").Key("WriteTimeout").MustInt(100),
		DbType:       cfg.Section("web").Key("DbType").String(),
		DbUser:       cfg.Section("web").Key("DbUser").String(),
		DbPassword:   cfg.Section("web").Key("DbPassword").String(),
		DbHost:       cfg.Section("web").Key("DbHost").String(),
		DbPort:       cfg.Section("web").Key("DbPort").Int(),
		DbName:       cfg.Section("web").Key("DbName").String(),
		DbSslmode:    cfg.Section("web").Key("DbSslmode").String(),
	}
}

func main() {

}
