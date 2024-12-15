package models

import "time"

type Image struct {
	ID          uint      `json:"id"`
	Service     uint      `json:"service"`
	Item        uint      `json:"item"`
	Image       []byte    `json:"image"`
	User        uint      `json:"user"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Notice      uint      `json:"notice"`
}
