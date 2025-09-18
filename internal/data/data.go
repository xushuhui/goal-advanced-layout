package data

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"

	"goal-advanced-layout/internal/conf"
	"goal-advanced-layout/internal/data/model"
	"goal-advanced-layout/pkg/log"

	_ "github.com/go-sql-driver/mysql"
)

var ProviderSet = wire.NewSet(NewDB, NewRedis, NewData, NewUserRepo)

type Data struct {
	query  *model.Queries
	rdb    *redis.Client
	logger *log.Logger
}

func NewData(db *sql.DB, rdb *redis.Client, logger *log.Logger) *Data {
	return &Data{
		query:  model.New(db),
		rdb:    rdb,
		logger: logger,
	}
}

func NewDB(conf *conf.Data, l *log.Logger) *sql.DB {
	db, err := sql.Open("mysql", conf.Database.Source)
	if err != nil {
		l.Errorf("failed to connect to database: %v", err)
		return nil
	}

	return db
}

func NewRedis(conf *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       int(conf.Redis.Db),
	})

	return rdb
}
