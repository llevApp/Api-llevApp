package models

type User struct {
	UserID     int
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Email      string `json:"email"`
	Nickname   string `json:"nickname"`
	CareerName string `json:"career_name"`
}
