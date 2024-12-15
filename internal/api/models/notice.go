package models

import "time"

type Notice struct {
	ID int `Json:"id"`
	Title string `Json:"title"`
	System uint`Json:"system"`
	User uint `Json:"user"`
	Date time.Time `Json:"date"`
	Type uint `Json:"type"`
	Path string `Json:"path"`
}

type NoticeType struct {
	ID int `Json:"id"`
	Name string `Json:"name"`
}