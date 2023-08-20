package storage

import (
	"Test_Task_0/config"
	"Test_Task_0/internal/models"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	//"github.com/jinzhu/gorm"
)

type OrderPostgres struct {
	db *gorm.DB
}

func NewOrderPostgres(db *gorm.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (o *OrderPostgres) GetAll() ([]models.Order, error) {
	var order []models.Order
	err := o.db.Preload("Delivery").Preload("Payment").Preload("Items").Find(&order).Error
	return order, err
}

func (o *OrderPostgres) Create(order models.Order) error {
	err := o.db.Create(&order).Error
	return err
}

func ConnectToPostgres(cfg *config.Config) *gorm.DB {
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Dbname,
		cfg.Database.Password,
	)

	//db, err := gorm.Open("postgres", conn)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("Error connection database: %s", err.Error())
		return nil
	}

	err = db.AutoMigrate(&models.Order{}, &models.Item{}, &models.Delivery{}, &models.Payment{})
	if err != nil {
		logrus.Fatalf("Error migrating database: %s", err.Error())
		return nil
	}

	return db
}

func CloseDBConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
