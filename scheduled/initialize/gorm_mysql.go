package initialize

import (
	"os"
	"padmin/scheduled/global"
	"padmin/scheduled/initialize/internal"
	"padmin/scheduled/model"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//@function: GormMysql
//@description: 初始化Mysql数据库
//@return: *gorm.DB

func GormMysql() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// RegisterTables
//@function: RegisterTables
//@description: 注册数据库表专用

func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(
		model.SysCrontab{},
		model.SysInterval{},
		model.SysPeriodicTask{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
