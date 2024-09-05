package db

import (
	"github.com/glebarez/sqlite"
	"github.com/sslime336/paper-airplane/config"
	"github.com/sslime336/paper-airplane/db/orm"
	"github.com/sslime336/paper-airplane/logging"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Sqlite *gorm.DB

var log *zap.Logger

func Init(config *config.App) {
	log = logging.Named("db")

	db, err := gorm.Open(sqlite.Open(config.Database.Sqlite.Path), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("failed to connect to database", zap.Error(err))
	}
	Sqlite = db
	if err := Sqlite.AutoMigrate(orm.Models()...); err != nil {
		log.Fatal("migration failed", zap.Error(err))
	}
}
