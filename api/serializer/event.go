package serializer

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mingrammer/meetup-api/api/model"
)

type CategorySerialzer struct {
	ID    uint `json:"id"`
	Title string `json:"title"`
}

type OwnerSerializer struct {
	ID        uint `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar"`
}

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
	ID          uint `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Place       PlaceSerializer `json:"place"`
	Datetime    DatetimeSerializer `json:"datetime"`
	Category    CategorySerialzer `json:"category"`
	Owner       OwnerSerializer `json:"owner"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func SerializeEvent(db *gorm.DB, event *model.Event) *EventSerialzer {
	owner := model.User{}
	db.Find(&owner, event.OwnerID)
	category := model.Category{}
	db.Find(&category, event.CategoryID)
	eventSerialzer := EventSerialzer{
		ID:          event.ID,
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
		Category: CategorySerialzer{
			ID:    category.ID,
			Title: category.Title,
		},
		Owner: OwnerSerializer{
			ID:        owner.ID,
			Name:      owner.Name,
			AvatarURL: owner.AvatarURL,
		},
		CreatedAt: event.CreatedAt,
		UpdatedAt: event.UpdatedAt,
	}
	return &eventSerialzer
}
