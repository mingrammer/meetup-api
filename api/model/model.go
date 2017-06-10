package model

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	UserID        string `gorm:"primary_key" json:"user_id"`
	Token         string `gorm:"unique" json:"token"`
	Name          string `gorm:"unique" json:"name"`
	AvatarURL     string `gorm:"avatar" json:"avatar_url"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	CreatedEvents []Event `gorm:"ForeignKey:OwnerToken" json:"created_events"`
	JoinedEvents  []Event `gorm:"many2many:user_joined_events"`
}

type Category struct {
	ID     uint `gorm:"primary_key" json:"id"`
	Title  string `gorm:"not null;unique" json:"title"`
	Events []Event `gorm:"ForeignKey:CategoryID" json:"related_events"`
}

type Place struct {
	PlaceTitle     string `json:"place_title"`
	PlaceLatitude  float64 `json:"place_lat"`
	PlaceLongitude float64 `json:"place_lon"`
}

type Datetime struct {
	DateStart *time.Time `json:"datetime_start"`
	DateEnd   *time.Time `json:"datetime_end"`
}

type Event struct {
	ID           uint `gorm:"primary_key" json:"id"`
	Title        string `gorm:"not null" json:"title"`
	Description  string `json:"description"`
	Place
	Datetime
	CategoryID   uint `json:"category_id"`
	OwnerToken   string `json:"-"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	Participants []User `gorm:"many2many:user_joined_events json:"participants"`
	Comments     []Comment `gorm:"ForeignKey:EventID" json:"comments"`
}

type Comment struct {
	ID          uint `gorm:"primary_key" json:"id"`
	Content     string `json:"content"`
	EventID     uint `json:"event_id"`
	WriterToken string `json:"-"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
