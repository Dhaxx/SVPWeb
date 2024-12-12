package models

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	Active rune `json:"active"`
	System uint `json:"system"`
	Notice uint `json:"notice"`
	Multi rune `json:"multi"`
	Control uint `json:"control"`
	PassMD5 string `json:"pass_md5"`
	Cad rune `json:"cad"`
}