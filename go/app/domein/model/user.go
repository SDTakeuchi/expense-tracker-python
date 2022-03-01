package models


func FindByID(id UUID) *User {

}

type User struct {
	BaseModel
	Email string `json:"email"`
	Password string `json:"password"`
}