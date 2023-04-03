package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable             = "users"
	projectorTable         = "projector"
	videoWallTable         = "video_wall"
	monitorTable           = "monitor"
	mountTable             = "mount"
	categoryTable          = "category"
	commQuantityTable      = "commercial_quantity"
	usersCommQuantityTable = "users_comm_quantity"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username,
		cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
