package representation

type User struct {
	Base

	Name     string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
