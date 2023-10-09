package database

import (
	"fmt"

	"github.com/ecommerce/configure"
	"github.com/ecommerce/internal/app/user/entities"
	"github.com/ecommerce/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
   var Datab *gorm.DB
func ConnectDatabase(cfg configure.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
        Datab = db
	db.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1000")
     

	db.AutoMigrate(&entities.OTPToken{})
	db.AutoMigrate(&model.User{})

	return db, dbErr
}
