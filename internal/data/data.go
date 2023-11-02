package data

import (
	"context"
	"fmt"
	"time"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"

	"nunu-http-layout/internal/conf"
	"nunu-http-layout/pkg/log"
)

var ProviderSet = wire.NewSet(NewDB,NewRedis,NewData,NewUserRepo)

type Data struct {
	db     *gorm.DB
	rdb    *redis.Client
	logger *log.Logger
}

func NewData(db *gorm.DB, rdb *redis.Client, logger *log.Logger) *Data {
	return &Data{
		db:     db,
		rdb:    rdb,
		logger: logger,
	}
}

func NewDB(conf *conf.Data, l *log.Logger) *gorm.DB {
	logger := zapgorm2.New(l.Logger)
	logger.SetAsDefault()
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{Logger: logger})
	if err != nil {
		panic(err)
	}
	db = db.Debug()
	return db
}

func NewRedis(conf *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       int(conf.Redis.Db),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("redis error: %s", err.Error()))
	}

	return rdb
}
