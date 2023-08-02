package request

type UserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Age      int8   `json:"age"`
	Password string `json:"password"`
}
