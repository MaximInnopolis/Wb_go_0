package storage

import (
	"Test_Task_0/internal/models"
	"github.com/jinzhu/gorm"
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
