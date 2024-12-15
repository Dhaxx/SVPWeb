package models

import "time"

type Service struct {
	ID uint `json:"id"`
	Client uint `json:"client"`
	StartDate time.Time `json:"startDate"`
	EndDate time.Time `json:"endDate"`
	Requester string `json:"requester"`
	Tel string `json:"tel"`
	Email string `json:"email"`
	Cell string `json:"cell"`
	Initial string `json:"initial"`
	Description string `json:"description"`
	Obs string `json:"obs"`
	Finished rune `json:"finished"`
	User int `json:"user"`
	Protocol string `json:"protocol"`
	System uint `json:"system"`
	UserAlteration uint `json:"userAlteration"`
	UserFinished uint `json:"userFinished"`
	Origin uint `json:"origin"`
}

type ServiceLog struct {
	ID uint `json:"id"`
	ServiceID uint `json:"serviceId"`
	UserID uint `json:"userId"`
	Requerer string `json:"requerer"`
	Initial string `json:"initial"`
	Description string `json:"description"`
	Date time.Time `json:"date"`
}