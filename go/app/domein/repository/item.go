package repository

import "app/domein/models"

type ItemRepository interface {
	Create(item *models.Item) (*models.Item, error)
	FindByID(id int) (*models.Item, error)
	Filter(itemf *models.ItemFilter) (*models.Item, error)
	Update(item *models.Item) (*models.Item, error)
	Delete(item *models.Item) error
}
