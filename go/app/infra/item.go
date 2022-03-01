package infra

import (
	"app/domein/model"
	"app/domein/repository"
	"github.com/jinzhu/gorm"
	"time"
)

type ItemRepository struct {
	Conn *gorm.DB
}

func NewItemRepository(conn *gorm.DB) repository.ItemRepository {
	return &ItemRepository{Conn: conn}
}

func (ir *ItemRepository) Create(item *model.Item) (*model.Item, error) {
	if err := ir.Conn.Create(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (ir *ItemRepository) FindByID(id int) (*model.Item, error) {
	item := &model.Item{Id: id}
	if err := ir.Conn.First(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (ir *ItemRepository) Filter(
	name, memo string,
	minPrice, maxPrice int,
	minPurchaseDate, maxPurchaseDate time.Time,
	largeCategoryId uint,
	smallCategoryId uint) (*model.Item, error) {
		if name == "" {
			name = "*"
		}
		if memo == "" {
			memo = "*"
		}
		if minPrice == nil {
			
		}

		item := &model.Item{Id: id}
		if err := ir.Conn.First(&item).Error; err != nil {
			return nil, err
	}

	return item, nil
}

func (ir *ItemRepository) Update(item *model.Item) (*model.Item, error) {
	if err := ir.Conn.Model(&item).Update(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (ir *ItemRepository) Delete(item *model.Item) error {
	if err := ir.Conn.Delete(&item).Error; err != nil {
		return err
	}

	return nil
}
