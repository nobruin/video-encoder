package database

import (
	"encoder/domain"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DbType        string
	DebugMode     bool
	AutoMigrateDb bool
	Env           string
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	db := NewDb()
	db.Env = "test"
	db.DbType = "sqlilte3"
	db.AutoMigrateDb = true
	db.DebugMode = true
	db.Dsn = "memory"
	connection, err := db.Connect()

	if err != nil {
		log.Fatal("Db Test error:  %w", err)
	}

	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	d.Db, err = gorm.Open(d.DbType, d.Dsn)
	if err != nil {
		return nil, err
	}

	d.Db.LogMode(d.DebugMode)

	if d.AutoMigrateDb {
		d.Db.AutoMigrate(&domain.Job{}, &domain.Video{})
		d.Db.Model(domain.Job{}).AddForeignKey("video_id", "videos (id)", "CASCADE", "CASCADE")
	}

	return d.Db, nil
}
