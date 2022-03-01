
type Item struct {
	gorm.Model
	Title           string    `json:"title"`
	Memo            string    `json:"memo"`
	IsExpense       bool      `json:"isExpense"` //default:true
	Price           int       `json:"price"`
	PurchaseDate    time.Time `json:"purchaseDate"`
	SmallCategoryId uint      `json:"smallCategoryId"`
	UserId          uint      `json:"userId"`
}
