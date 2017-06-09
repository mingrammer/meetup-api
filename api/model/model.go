package model

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	UserID         string `gorm:"primary_key" json:"user_id"`
	Token         string `gorm:"unique" json:"token"`
	Name          string `gorm:"unique" json:"name"`
	AvatarURL        string `gorm:"avatar" json:"avatar_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     *time.Time `sql:"index" json:"deleted_at`
	CreatedEvents []Event `gorm:"ForeignKey:OwnerID" json:"created_events"`
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
	ID          uint `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"deleted_at`
	Place
	Datetime
	Title       string `gorm:"not null" json:"event_title"`
	Description string `json:"description"`
	OwnerToken  string `json:"-"`
	CategoryID  uint `json:"category_id"`
	Comments    []Comment `gorm:"ForeignKey:EventID" json:"comments"`
}

type Comment struct {
	ID          uint `gorm:"primary_key" json:"id"`
	Content     string `json:"content"`
	WriterToken string `json:"-"`
	EventID     uint `json:"event_id"`
}
