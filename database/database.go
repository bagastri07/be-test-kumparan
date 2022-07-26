package database

import (
	"time"

	"github.com/bagastri07/be-test-kumparan/services/config"
	"github.com/go-sql-driver/mysql"
	_ "github.com/newrelic/go-agent/v3/integrations/nrmysql"

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

	db, err := sqlx.Connect("nrmysql", dsn)
	if err != nil {
		panic(err)
	}
	db.DB.SetConnMaxIdleTime(time.Duration(config.GetConfig().DBMaxIdleConns) * time.Second)
	db.DB.SetMaxOpenConns(conf.DBMaxOpenConns)
	db.DB.SetConnMaxLifetime(time.Duration(conf.DBMaxIdleConns) * time.Second)

	return db, err
}
