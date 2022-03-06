package infra

import (
	"app/domein/models"
	"app/domein/repository"
	"github.com/jinzhu/gorm"
)

type ItemRepository struct {
	Conn *gorm.DB
}

func NewItemRepository(conn *gorm.DB) repository.ItemRepository {
	return &ItemRepository{Conn: conn}
}

func (ir *ItemRepository) Create(item *models.Item) (*models.Item, error) {
	if err := ir.Conn.Create(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (ir *ItemRepository) FindByID(id int) (*models.Item, error) {
	item := &models.Item{ID: id}
	if err := ir.Conn.First(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (ir *ItemRepository) Get(iFil *models.ItemFilter) (items []*models.Item, err error) {
	q := ir.Conn

	if iFil.Name != "" {
		name := "%" + iFil.Name + "%"
		q = q.Where("name LIKE ?", name)
	}
	if iFil.Memo != "" {
		memo := "%" + iFil.Memo + "%"
		q = q.Where("memo LIKE ?", memo)
	}

	if iFil.MaxPrice == 0 {
		q = q.Where("price > ?", iFil.MinPrice)
	} else {
		q = q.Where(q.Where("price BETWEEN ? AND ?", iFil.MinPrice, iFil.MaxPrice))
	}

	if iFil.MaxPurchaseDate.IsZero() {
		q = q.Where("purchase_date > ?", iFil.MinPurchaseDate)
	} else {
		q = q.Where("purchase_date BETWEEN ? AND ?", iFil.MinPurchaseDate, iFil.MinPurchaseDate)
	}

	if iFil.SmallCategoryId != "" {
		q = q.Where("small_category_id IN ?", iFil.SmallCategoryId)
	} 
	// if iFil.LargeCategoryId != "" {
	// 	q = q.Where("large_category_id = ?", iFil.LargeCategoryId)
	// } 

	if err := ir.Conn.Where(&iFil).Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (ir *ItemRepository) Update(item *models.Item) (*models.Item, error) {
	if err := ir.Conn.Model(&item).Update(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (ir *ItemRepository) Delete(item *models.Item) error {
	if err := ir.Conn.Delete(&item).Error; err != nil {
		return err
	}

	return nil
}
