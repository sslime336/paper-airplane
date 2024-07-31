package db

import (
	"github.com/glebarez/sqlite"
	"github.com/sslime336/paper-airplane/config"
	"github.com/sslime336/paper-airplane/logging"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Sqlite *gorm.DB

func Init() {
	db, err := gorm.Open(sqlite.Open(config.App.Database.Sqlite.Path), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		logging.Fatal("failed to connect to database", zap.Error(err))
	}
	Sqlite = db
	if err := Sqlite.AutoMigrate(&Session{}); err != nil {
		logging.Fatal("migration failed", zap.Error(err))
	}
}
