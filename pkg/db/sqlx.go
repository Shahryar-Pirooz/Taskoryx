package db

import (
	"fmt"
	"tasoryx/config"
	"tasoryx/pkg/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var (
	DB  *sqlx.DB
	log *zap.Logger
)

func Init(cnf config.Database) {
	log = logger.GetLogger()
	host := cnf.Host
	port := cnf.Port
	user := cnf.User
	pass := cnf.Password
	name := cnf.Name

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, pass, name, host, port)

	var err error

	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Panic(fmt.Sprintf("could not connect to postgresql: %v", err))
	}
	log.Info("Connected to PostgreSQL")
	EnsureAllTables()
}

func EnsureAllTables() {
	EnsureTaskTables()
	EnsureUserTables()
}

func EnsureTaskTables() {
	taskSchema := `CREATE TABLE IF NOT EXISTS tasks (
	id SERIAL PRIMARY KEY,      
	title TEXT NOT NULL,
	description TEXT,
	status SMALLINT NOT NULL,
	date_due  TIMESTAMP,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted_at TIMESTAMP
	);`

	_, err := DB.Exec(taskSchema)
	if err != nil {
		log.Panic(fmt.Sprintf("failed to create task table: %v", err))
	}
	log.Info("users table checked/created")
}

func EnsureUserTables() {
	userSchema := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50),
		email VARCHAR(100),
		password TEXT,
		role SMALLINT,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL,
		deleted_at TIMESTAMP
		);`
	_, err := DB.Exec(userSchema)
	if err != nil {
		log.Panic(fmt.Sprintf("failed to create user table: %v", err))
	}
	log.Info("users table checked/created")
}
