package cache

import (
	"Test_Task_0/internal/models"
	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
	"sync"
)

type Cache struct {
	mx    sync.RWMutex
	items map[string]models.Order
}

func NewCache(db *gorm.DB) *Cache {
	return &Cache{items: make(map[string]models.Order)}
}

func (c *Cache) Set(order models.Order) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.items[order.OrderUid] = order
}

func (c *Cache) GetByUid(uid string) (models.Order, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.items[uid]
	return val, ok
}

func (c *Cache) GetAll() []models.Order {
	var order []models.Order
	c.mx.RLock()
	defer c.mx.RUnlock()
	for _, b := range c.items {
		order = append(order, b)
	}
	return order
}
