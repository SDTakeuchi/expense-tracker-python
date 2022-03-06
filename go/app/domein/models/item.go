package models

import (
	"app/lib"
	"errors"
	"time"
	null "gopkg.in/guregu/null.v4"
)

type Item struct {
	BaseModel
	Name            string    `db:"name" json:"name"`
	Memo            string    `db:"memo" json:"memo"`
	Price           int       `db:"price" json:"price"`
	PurchaseDate    time.Time `db:"purchase_date" json:"purchase_date"`
	SmallCategoryId string    `db:"small_category_id" json:"small_category_id"`
	// SmallCategory SmallCategory
	UserId string `db:"user_id" json:"user_id"`
	// User User
}

type ItemFilter struct {
	Name            null.String `json:"name"`
	Memo            null.String `json:"memo"`
	MinPrice        null.Int    `json:"min_price"`
	MaxPrice        null.Int    `json:"max_price"`
	MinPurchaseDate null.Time   `json:"min_purchase_date"`
	MaxPurchaseDate null.Time   `json:"max_purchase_date"`
	SmallCategoryId null.String `json:"small_category_id"`
	LargeCategoryId null.String `json:"large_category_id"`
}

/*
constructor is for preparing objects (returns struct and error)
setter is for updating new objects (returns error if failed to udpate)
*/

// constructor
func NewItem(
	name, memo string,
	price int,
	purchaseDate time.Time,
	smallCategoryId ,userId string) (*Item, error) {

	if name == "" {
		return nil, errors.New("項目名を入力してください")
	}

	now := time.Now()
	id := lib.GenUuid()

	item := &Item{
		BaseModel: BaseModel{
			ID:        id,
			CreatedAt: now,
			UpdatedAt: now,
		},
		Name:            name,
		Memo:            memo,
		Price:           price,
		PurchaseDate:    purchaseDate,
		SmallCategoryId: smallCategoryId,
		UserId:          userId,
	}

	return item, nil
}

// setter
func (i *Item) Set(
	name, memo string,
	price int,
	purchaseDate time.Time,
	smallCategoryId string) error {

	if name == "" {
		return errors.New("項目名を入力してください")
	}

	now := time.Now()

	i.UpdatedAt = now
	i.Name = name
	i.Memo = memo
	i.Price = price
	i.PurchaseDate = purchaseDate
	i.SmallCategoryId = smallCategoryId

	return nil
}
