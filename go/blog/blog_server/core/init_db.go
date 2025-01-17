package core

import (
  "blog_server/global"
  "github.com/sirupsen/logrus"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "time"
)

func InitDB() *gorm.DB {

  dc := global.Config.DB

  db, err := gorm.Open(mysql.Open(dc.DSN()), &gorm.Config{
    DisableForeignKeyConstraintWhenMigrating: true, // 不生成外键约束
  })
  if err != nil {
    logrus.Fatalf("数据库连接失败 %s", err)
  }
  sqlDB, err := db.DB()
  sqlDB.SetMaxIdleConns(dc.Max_idle_conns)
  sqlDB.SetMaxOpenConns(dc.Max_open_conns)
  sqlDB.SetConnMaxLifetime(time.Hour)
  logrus.Infof("数据库连接成功！")
  return db
}
