package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"os"
	"relia_system/app_context"
	"relia_system/app_context/tokenprovider/jwt"
	"relia_system/modules/product/product_model"
	"relia_system/modules/user/user_model"
	"relia_system/route"
	"relia_system/shared"
	"relia_system/shared/database"
	"relia_system/shared/env"
	"relia_system/shared/logger"
	"time"
)

func main() {
	appLogger := logger.LogConfigure(logger.LogConfig{
		ConsoleLoggingEnabled: true,
		FileLoggingEnabled:    true,
		EncodeLogsAsJson:      true,
		Directory:             os.Getenv("LOG_DIR"),
		Filename:              os.Getenv("LOG_FILENAME"),
		MaxSize:               500,
		MaxBackups:            500,
		MaxAge:                1,
	})

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	//jwtExpired := os.Getenv("JWT_EXPIRED")
	//jwtRefreshExpired := os.Getenv("JWT_REFRESH_EXPIRED")

	tokenProvider := jwt.NewTokenJWTProvider(jwtSecretKey, 60*60*24*30)

	dbService := database.NewDbService(&database.DbConfig{
		DBName:   os.Getenv("DB_NAME"),
		DBPasswd: os.Getenv("DB_PASSWORD"),
		DBHost:   os.Getenv("DB_HOST"),
		DBPort:   env.GetInt("DB_PORT"),
		DBUser:   os.Getenv("DB_USER"),
	}, &database.DBConfigTemplate{
		ConnTimeout:  time.Duration(env.GetInt64("DB_CONN_TO")) * time.Second,
		ReadTimeout:  time.Duration(env.GetInt64("DB_READ_TO")) * time.Second,
		WriteTimeout: time.Duration(env.GetInt64("DB_WRITE_TO")) * time.Second,
		MaxOpenConns: env.GetInt("DB_MAX_CONN"),
		MaxIdleConns: env.GetInt("DB_MAX_IDLE_CONN"),
		KeepAlive:    time.Duration(env.GetInt64("DB_KEEP_ALIVE")) * time.Second,
	})

	if err := dbService.Init(); err != nil {
		appLogger.Fatal().Msgf("Không thể kết nối đến DB: %s", err.Error())
	}

	//TODO: this one purpose for demo
	dbService.Db.AutoMigrate(&user_model.User{})
	dbService.Db.AutoMigrate(&product_model.Product{})

	appCtx := app_context.NewAppContext(dbService.Db.Debug(),
		tokenProvider,
		appLogger)

	engine := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(shared.JsonTagNameFunc)
	}

	route.SetupRouter(engine, appCtx)
	engine.Run(":8088")

	appLogger.Info().Msgf("Server listen on port 8088")

}
