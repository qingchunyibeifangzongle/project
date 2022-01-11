package config

import (
	"github.com/qingchunyibeifangzongle/project/core"
)

func Mysql() {
	core.Set("mysql", core.StrMap{
		"host":           core.Env("DB_HOST", "127.0.0.1"),
		"port":           core.Env("DB_PORT", "3306"),
		"database":       core.Env("DB_DATABASE", ""),
		"user":           core.Env("DB_USERNAME", ""),
		"pwd":            core.Env("DB_PASSWORD", ""),
		"charset":        "utf8mb4",
		"collation":      "utf8mb4_unicode_ci",
		"prefix":         core.Env("DB_PREFIX", ""),
		"max_id_leconns": core.Env("DB_MAX_ID_LECONNS", ""), // 空闲中的最大连接数
		"max_open_conns": core.Env("DB_MAX_OPEN_CONNS", ""), // 打开到数据库的最大连接数
	})
}

func Pgsql() {
	core.Set("pgsql", core.StrMap{
		"host":           core.Env("DB_HOST", "127.0.0.1"),
		"port":           core.Env("DB_PORT", "5432"),
		"database":       core.Env("DB_DATABASE", ""),
		"user":           core.Env("DB_USERNAME", ""),
		"pwd":            core.Env("DB_PASSWORD", ""),
		"charset":        "utf8",
		"prefix":         core.Env("DB_PREFIX", ""),
		"max_id_leconns": core.Env("DB_MAX_ID_LECONNS", ""),
		"max_open_conns": core.Env("DB_MAX_OPEN_CONNS", ""),
	})
}
