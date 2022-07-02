package models

type User struct {
	UserID     int    `json:"user_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Nickname   string `json:"nickname"`
	CareerName string `json:"career_name"`
	CareerId   int    `json:"career_id"`
	UUID       string `json:"uuid_fb"`
}

type Carrer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func NewUser() *User {
	return &User{}
}
