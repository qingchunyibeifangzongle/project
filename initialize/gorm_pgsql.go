package initialize

import (
	"github.com/qingchunyibeifangzongle/project/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func ConnectPgsqlDB() *gorm.DB {
	pgsqlConfig := postgres.Config{
		DSN:                  PgsqlDsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true, Logger: logger.Default.LogMode(logger.Error)}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(core.GetToInt("mysql.max_id_leconns")) // 空闲中的最大连接数
		sqlDB.SetMaxOpenConns(core.GetToInt("mysql.max_open_conns")) // 打开到数据库的最大连接数
		// 设置每个链接的过期时间
		sqlDB.SetConnMaxLifetime(5 * time.Minute)
		return db
	}
}

func PgsqlDsn() string {
	return "host=" + core.GetToString("pgsql.host") + " user=" + core.GetToString("pgsql.user") + " password=" + core.GetToString("pgsql.pwd") + " dbname=" + core.GetToString("pgsql.database") + " port=" + core.GetToString("pgsql.port") + " " + "?charset=utf8&parseTime=True&loc=Local"
}
