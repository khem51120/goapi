package config

import (
	"backend/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	var err error
	var dsn = "host=rain.db.elephantsql.com user=qvngikec dbname=qvngikec port=5432 password=mnmHBZ_IcR-lnolTfddSGrgYvYdDCMXU"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("postgres err: " + err.Error())
		panic(err)
	}
	migrateDB()
	return err
}

func migrateDB() {
	DB.AutoMigrate(models.Channel{})
	DB.AutoMigrate(models.Template{})
	DB.AutoMigrate(models.Savety{})
	DB.AutoMigrate(models.Saveunit{})
	DB.AutoMigrate(models.NodeMCU{})
}
