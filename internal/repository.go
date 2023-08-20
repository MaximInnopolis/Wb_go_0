package internal

import (
	"Test_Task_0/internal/cache"
	"Test_Task_0/internal/models"
	"Test_Task_0/internal/storage"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type OrderRepo interface {
	GetAll() ([]models.Order, error)
	Create(order models.Order) error
}

type CacheRepo interface {
	Set(order models.Order)
	GetByUid(uuid string) (models.Order, bool)
	GetAll() []models.Order
}

type Repository struct {
	OrderRepo
	CacheRepo
}

func NewRepository(db *gorm.DB) *Repository {
	rdb := storage.NewOrderPostgres(db)
	rcache := cache.NewCache(db)
	items, err := rdb.GetAll()
	if err != nil {
		return &Repository{
			OrderRepo: rdb,
			CacheRepo: rcache,
		}
	}
	for _, item := range items {
		rcache.Set(item)
	}
	logrus.Println("Cache loaded successfully")
	return &Repository{
		OrderRepo: rdb,
		CacheRepo: rcache,
	}
}
