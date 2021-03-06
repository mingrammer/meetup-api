package model

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID            uint `gorm:"primary_key" json:"id,omitempty"`
	Token         string `sql:"index" gorm:"unique" json:"token,omitempty"`
	SlackUserID   string `gorm:"unique" json:"user_id,omitempty"`
	Name          string `gorm:"unique" json:"name,omitempty"`
	AvatarURL     string `gorm:"avatar" json:"avatar_url,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	CreatedEvents []Event `gorm:"ForeignKey:OwnerID" json:"-"`
	JoinedEvents  []Event `gorm:"many2many:user_joined_events" json:"-"`
}

type Category struct {
	ID     uint `gorm:"primary_key" json:"id,omitempty"`
	Title  string `gorm:"not null;unique" json:"title,omitempty"`
	Events []Event `gorm:"ForeignKey:CategoryID" json:"-"`
}

type Place struct {
	PlaceTitle string `json:"title,omitempty"`
	PlaceLat   float64 `json:"lat,omitempty"`
	PlaceLon   float64 `json:"lon,omitempty"`
}

type Datetime struct {
	DateStart *time.Time `json:"start,omitempty"`
	DateEnd   *time.Time `json:"end,omitempty"`
}

type Event struct {
	ID            uint `gorm:"primary_key" json:"id,omitempty"`
	Title         string `gorm:"not null" json:"title,omitempty"`
	Description   string `json:"description,omitempty"`
	Place `json:"place,omitempty"`
	Datetime `json:"datetime,omitempty"`
	CategoryID    uint `json:"category_id,omitempty"`
	CategoryTitle string `json:"category_title,omitempty"`
	OwnerID       uint `json:"owner_id,omitempty"`
	OwnerName     string `json:"owner_name,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	Participants  []User `gorm:"many2many:user_joined_events" json:"-"`
	Comments      []Comment `gorm:"ForeignKey:EventID" json:"-"`
}

type Comment struct {
	ID         uint `gorm:"primary_key" json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	EventID    uint `json:"event_id,omitempty"`
	WriterID   uint `json:"writer_id,omitempty"`
	WriterName string `json:"writer_name,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}
