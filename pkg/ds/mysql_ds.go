package ds

import (
	"fmt"
	"os"
	"v-shi/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func LoadDB() *gorm.DB {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	name := os.Getenv("MYSQL_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&models.Category{},
		&models.ShopOwner{},
		&models.Shop{},
		&models.Location{},
		&models.Food{},
		&models.FoodImage{},
	)

	if err != nil {
		panic(err)
	}

	return db

}
