package models

type User struct {
	Id        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt string `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type UserDTO struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	OldPassword string `json:"old_password"`
	Avatar      string `json:"avatar"`
}
