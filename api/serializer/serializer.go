package serializer

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mingrammer/meetup-api/api/model"
)

type PlaceSerializer struct {
	PlaceTitle     string `json:"title"`
	PlaceLatitude  float64 `json:"lat"`
	PlaceLongitude float64 `json:"lon"`
}

type DatetimeSerializer struct {
	DateStart *time.Time `json:"start"`
	DateEnd   *time.Time `json:"end"`
}

type EventSerialzer struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	Place         PlaceSerializer `json:"place"`
	Datetime      DatetimeSerializer `json:"datetime"`
	CategoryTitle string `json:"category_title"`
	OwnerName     string `json:"owner_name"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	Participants  []model.User `json:"participants"`
	Comments      []model.Comment `json:"comments"`
}

func SerializeEvent(db *gorm.DB, event *model.Event) *EventSerialzer {
	owner := model.User{}
	db.Where(&model.User{ID: event.OwnerID}).Find(&owner)
	category := model.Category{}
	db.Find(&category, event.CategoryID)
	eventSerialzer := EventSerialzer{
		Title:       event.Title,
		Description: event.Description,
		Place: PlaceSerializer{
			PlaceTitle:     event.Place.PlaceTitle,
			PlaceLatitude:  event.Place.PlaceLat,
			PlaceLongitude: event.Place.PlaceLon,
		},
		Datetime: DatetimeSerializer{
			DateStart: event.Datetime.DateStart,
			DateEnd:   event.Datetime.DateEnd,

		},
		OwnerName:     owner.Name,
		CategoryTitle: category.Title,
		CreatedAt:     event.CreatedAt,
		UpdatedAt:     event.UpdatedAt,
		Participants:  event.Participants,
		Comments:      event.Comments,
	}
	return &eventSerialzer
}
