package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Token         string `gorm:"primary_key" json:"token"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
	CreatedEvents []Event `gorm:"ForeignKey:OwnerID" json:"created_events"`
	JoinedEvents  []Event `gorm:"many2many:user_joined_events"`
}

type Category struct {
	ID     uint `gorm:"primary_key"`
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
	gorm.Model
	Place
	Datetime
	Title       string `gorm:"not null" json:"event_title"`
	Description string `json:"description"`
	OwnerToken  string `json:"-"`
	CategoryID  uint `json:"category_id"`
	Comments    []Comment `gorm:"ForeignKey:EventID" json:"comments"`
}

type Comment struct {
	gorm.Model
	Content     string `json:"content"`
	WriterToken string `json:"-"`
	EventID     uint `json:"event_id"`
}
