package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type DbConfig struct {
	DBHost   string
	DBPort   int
	DBUser   string
	DBPasswd string
	DBName   string
}

type DBConfigTemplate struct {
	// mysql connect timeout.
	ConnTimeout time.Duration

	// mysql read timeout.
	ReadTimeout time.Duration

	// mysql write timeout.
	WriteTimeout time.Duration

	// max num of connections.
	MaxOpenConns int

	// max num of idle connections.
	MaxIdleConns int

	// max life time of connection.
	KeepAlive time.Duration
}

type DbService struct {
	config         *DbConfig
	Db             *gorm.DB
	configTemplate *DBConfigTemplate
}

func NewDbService(config *DbConfig, configTem *DBConfigTemplate) *DbService {
	dbs := &DbService{config: config, configTemplate: configTem}
	return dbs
}

func (s *DbService) Init() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local&timeout=%s&readTimeout=%s&writeTimeout=%s&charset=%s",
		s.config.DBUser,
		s.config.DBPasswd,
		s.config.DBHost,
		s.config.DBPort,
		s.config.DBName,
		s.configTemplate.ConnTimeout,
		s.configTemplate.ReadTimeout,
		s.configTemplate.WriteTimeout,
		"utf8mb4",
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxOpenConns(s.configTemplate.MaxOpenConns)
	sqlDB.SetMaxIdleConns(s.configTemplate.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(s.configTemplate.KeepAlive)

	s.Db = db

	return nil
}
