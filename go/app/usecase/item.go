package usecase

import (
    "time"
	"app/domain/model"
	"app/domain/repository"
)

// interface of ItemUsecase item usecase
type ItemUsecase interface {
	Create(
        name, memo string,
        price int,
        purchaseDate time.Time,
        SmallCategoryId string) (*model.Item, error)
    Get(
        name, memo string,
        MinPrice, MaxPrice int,
        MinPurchaseDate, MaxPurchaseDate time.Time,
        SmallCategoryId string,
        LargeCategoryId string) ([]*model.Item, error)
	FindByID(id int) (*model.Item, error)
	Update(
        id int,
        name, memo string,
        price int,
        purchaseDate time.Time,
        SmallCategoryId string) (*model.Item, error)
	Delete(id int) error
}

type itemUsecase struct {
	itemRepo repository.ItemRepository
}

// NewItemUsecase item usecaseのコンストラクタ
func NewItemUsecase(itemRepo repository.ItemRepository) ItemUsecase {
	return &itemUsecase{itemRepo: itemRepo}
}

// Create itemを保存するときのユースケース
func (iu *itemUsecase) Create(title, content string) (*model.Item, error) {
	item, err := model.NewItem(title, content)
	if err != nil {
		return nil, err
	}

	createdItem, err := iu.itemRepo.Create(item)
	if err != nil {
		return nil, err
	}

	return createdItem, nil
}

// FindByID itemをIDで取得するときのユースケース
func (iu *itemUsecase) FindByID(id int) (*model.Item, error) {
	foundItem, err := iu.itemRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundItem, nil
}

// Update itemを更新するときのユースケース
func (tu *itemUsecase) Update(id int, title, content string) (*model.Item, error) {
	targetItem, err := tu.itemRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = targetItem.Set(title, content)
	if err != nil {
		return nil, err
	}

	updatedItem, err := tu.itemRepo.Update(targetItem)
	if err != nil {
		return nil, err
	}

	return updatedItem, nil
}

// Delete itemを削除するときのユースケース
func (tu *itemUsecase) Delete(id int) error {
	item, err := tu.itemRepo.FindByID(id)
	if err != nil {
		return err
	}

	err = tu.itemRepo.Delete(item)
	if err != nil {
		return err
	}

	return nil
}
