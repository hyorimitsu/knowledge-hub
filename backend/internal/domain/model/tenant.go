package model

import "time"

type Tenant struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Domain    string    `json:"domain" gorm:"unique"`
	Settings  Settings  `json:"settings" gorm:"embedded"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Settings struct {
	Theme    Theme    `json:"theme" gorm:"embedded"`
	Features Features `json:"features" gorm:"embedded"`
}

type Theme struct {
	PrimaryColor   string `json:"primary_color"`
	SecondaryColor string `json:"secondary_color"`
}

type Features struct {
	Comments bool `json:"comments"`
	Tags     bool `json:"tags"`
	Ratings  bool `json:"ratings"`
}
