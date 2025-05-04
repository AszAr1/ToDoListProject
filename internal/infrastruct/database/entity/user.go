package entity

type UserEntity struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	//Picture string maybe will be late
}
