package resource

import (
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlMap = make(map[string]*gorm.DB, 0)

func GetMysql(name string) *gorm.DB {
	if db, ok := MysqlMap[name]; ok {
		return db
	}
	return nil
}

//func InitMysql(ctx context.Context, c staticconfig.MysqlStaticConf) error {
//go func() {
//	db, err := gorm.Open(mysql.New(mysql.Config{
//		DSN:                       c.DSN, // DSN data source name
//		DefaultStringSize:         256,   // string 类型字段的默认长度
//		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
//		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
//		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
//		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
//	}), &gorm.Config{Logger: zorm.NewLogger().LogMode(c.LogLevel)})
//	if err != nil {
//		TimeChan <- err
//		return
//	}
//	var sqlDB *sql.DB
//	if sqlDB, err = db.DB(); err != nil {
//		TimeChan <- err
//		return
//	}
//	sqlDB.SetMaxIdleConns(c.MaxIdle)
//	sqlDB.SetMaxOpenConns(c.MaxConn)
//	//sqlDB.SetConnMaxLifetime(time.Hour)
//
//	MysqlMap[c.Name] = db
//	TimeChan <- err
//	return
//}()
//
//select {
//case err := <-TimeChan:
//	return err
//case <-ctx.Done():
//	return errorx.New("InitDB 执行超时")
//}
//}

//type SQL struct {
//	db map[string]*gorm.DB
//}
//
//func NewSQL() *SQL {
//	return &SQL{
//		db: gormDBMap,
//	}
//}
//
//func (s *SQL) GetDB(name string) *gorm.DB {
//	if db, ok := s.db[name]; ok {
//		return db
//	}
//
//	return nil
//}
