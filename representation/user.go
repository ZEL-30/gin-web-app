package representation

type User struct {
	Base

	Name     string `json:"name"`
	Password string `json:"password"`
}
