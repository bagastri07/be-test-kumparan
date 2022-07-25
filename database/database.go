package database

import (
	"time"

	"github.com/bagastri07/be-test-kumparan/services/config"
	"github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

type Manager struct {
	DB *sqlx.DB
}

func NewConnection(conf config.Config) (*sqlx.DB, error) {
	mysqlConfig := mysql.Config{
		DBName: conf.DBName,
		User:   conf.DBUser,
		Passwd: conf.DBPassword,
		Addr:   conf.DBHost,

		ParseTime: true,
		Net:       "tcp",
	}

	dsn := mysqlConfig.FormatDSN()

	db, err := sqlx.Connect("mysql", dsn)
	db.SetConnMaxIdleTime(time.Duration(conf.DBMaxIdleConns) * time.Second)
	db.SetMaxOpenConns(conf.DBMaxOpenConns)

	return db, err
}
