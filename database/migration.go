package database

import (
	"log"

	"github.com/ihsankarim/backend-brighted/config"
	"github.com/ihsankarim/backend-brighted/internal/features/auth"
)

func Migrate() {
	err := config.DB.AutoMigrate(
		&auth.User{},
	)
	if err != nil {
		log.Fatal("AutoMigrate error: ", err)
	}
}
