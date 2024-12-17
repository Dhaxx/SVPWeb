package models

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	Active string `json:"active"`
	System uint `json:"system"`
	Notice uint `json:"notice"`
	Multi string `json:"multi"`
	Control uint `json:"control"`
	PassMD5 string `json:"pass_md5"`
	Cad string `json:"cad"`
}