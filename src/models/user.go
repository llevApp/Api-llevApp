package models

type User struct {
	UserID     int    `json:"user_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Nickname   string `json:"nickname"`
	CareerName string `json:"career_name"`
	UUID       string `json:"uuid_fb"`
}

func NewUser() *User {
	return &User{}
}
