package app_context

import (
	"gorm.io/gorm"
	"relia_system/app_context/tokenprovider"
	"relia_system/shared/logger"
)

type AppCtx interface {
	GetDBConnection() *gorm.DB
	GetTokenProvider() tokenprovider.Provider
	GetLogger() *logger.Logger
}

type appCtx struct {
	db          *gorm.DB
	tokProvider tokenprovider.Provider
	logger      *logger.Logger
}

func NewAppContext(db *gorm.DB,
	tokProvider tokenprovider.Provider,
	logger *logger.Logger) *appCtx {
	return &appCtx{db: db,
		tokProvider: tokProvider,
		logger:      logger}
}

func (ctx *appCtx) GetDBConnection() *gorm.DB {
	return ctx.db.Session(&gorm.Session{NewDB: true})
}

func (ctx *appCtx) GetTokenProvider() tokenprovider.Provider {
	return ctx.tokProvider
}

func (ctx *appCtx) GetLogger() *logger.Logger {
	return ctx.logger
}
