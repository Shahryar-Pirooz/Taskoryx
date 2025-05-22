package app

import (
	"context"
	"fmt"
	"tasoryx/config"
	"tasoryx/internal/task"
	taskPort "tasoryx/internal/task/port"
	"tasoryx/internal/user"
	UserPort "tasoryx/internal/user/port"
	"tasoryx/pkg/adapters/storage"
	"tasoryx/pkg/cache"
	appCtx "tasoryx/pkg/context"
	appDB "tasoryx/pkg/db"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App interface {
	TaskService(ctx context.Context) taskPort.Service
	UserService(ctx context.Context) UserPort.Service
	DB() *gorm.DB
	Config() config.Config
	Cache() *cache.RedisAdapter
}

type app struct {
	db           *gorm.DB
	cfg          config.Config
	userService  UserPort.Service
	taskService  taskPort.Service
	redisAdapter cache.RedisAdapter
}

func (a *app) userServiceWithDB(db *gorm.DB) UserPort.Service {
	return user.NewService(storage.NewUserRepo(db))
}

func (a *app) taskServiceWithDB(db *gorm.DB) taskPort.Service {
	return task.NewTask(storage.NewTaskRepo(db))
}

func (a *app) UserService(ctx context.Context) UserPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.userService == nil {
			a.userService = a.userServiceWithDB(a.db)
		}
		return a.userService
	}
	return a.userServiceWithDB(db)
}

func (a *app) TaskService(ctx context.Context) taskPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.taskService == nil {
			a.taskService = a.taskServiceWithDB(a.db)
		}
		return a.taskService
	}
	return a.taskServiceWithDB(db)
}

func (a *app) setDB() {
	db := appDB.NewPSQLConnection(a.cfg.Database)
	a.db = db
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) Cache() *cache.RedisAdapter {
	return &a.redisAdapter
}

func NewApp(cfg config.Config) App {
	a := &app{
		cfg: cfg,
	}
	a.setDB()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	a.redisAdapter = *cache.NewRedisAdapter(redisClient)
	return a
}
