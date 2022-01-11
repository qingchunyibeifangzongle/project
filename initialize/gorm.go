package initialize

import (
	"fmt"
	"github.com/qingchunyibeifangzongle/project/config"
	"github.com/qingchunyibeifangzongle/project/core"
	"github.com/qingchunyibeifangzongle/project/model"
	"gorm.io/gorm"
)

func SetUpDB() {
	//加载数据
	db := Gorm()
	if db == nil {
		//err
	}
	//migrate
	migrate(db)
}

func Gorm() *gorm.DB {
	if db_connection := core.Env("DB_CONNECTION"); db_connection != "" {
		switch db_connection {
		case "mysql":
			config.Mysql()
			return ConnectMysqlDB()
		case "pgsql":
			config.Pgsql()
			return ConnectPgsqlDB()
		default:
			config.Mysql()
			return ConnectMysqlDB()
		}
	} else {
		fmt.Println("please chose db_connection!")
		return nil
	}
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
	)
}
