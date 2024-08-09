package main

import (
	"github.com/sslime336/paper-airplane/config"
	"github.com/sslime336/paper-airplane/db"
	"github.com/sslime336/paper-airplane/db/orm"
	"github.com/sslime336/paper-airplane/logging"
	"gorm.io/gen"
)

func init() {
	config.Init()
	logging.Init()
	db.Init()
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dao",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(db.Sqlite)

	g.ApplyBasic(orm.Models()...)

	g.Execute()
}
