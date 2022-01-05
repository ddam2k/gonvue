package user

type Account struct {
	Id       string `json:"id"`
	password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}
