package initialize

import (
	"github.com/qingchunyibeifangzongle/project/core"
	mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func ConnectMysqlDB() *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN:                       MysqlDsn(), // DSN data source name
		DefaultStringSize:         255,        // string 类型字段的默认长度
		SkipInitializeWithVersion: false,      // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true, Logger: logger.Default.LogMode(logger.Error)}); err != nil {
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

func MysqlDsn() string {
	return core.GetToString("mysql.user") + ":" + core.GetToString("mysql.pwd") + "@tcp(" + core.GetToString("mysql.host") + ":" + core.GetToString("mysql.port") + ")/" + core.GetToString("mysql.database") + "?charset=utf8&parseTime=True&loc=Local"
}
