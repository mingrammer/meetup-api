package bot

import (
	"errors"
	"time"

	"github.com/jinzhu/now"
	db "github.com/mingrammer/meetup-api/api/database"
	"github.com/mingrammer/meetup-api/api/model"
)

func GetEventsByDate(date string) []model.Event {
	var dateStart time.Time
	var dateEnd time.Time
	var dateStartString string
	var dateEndString string
	dateStart = time.Now().Local()
	switch date {
	case "오늘":
		dateEnd = now.EndOfDay()
	case "이번주":
		dateEnd = now.EndOfWeek()
	case "다음주":
		dateStart = now.BeginningOfWeek().AddDate(0, 0, 7)
		dateEnd = now.EndOfWeek().AddDate(0, 0, 7)
	case "이번달":
		dateEnd = now.EndOfMonth()
	case "다음달":
		dateStart = now.BeginningOfMonth().AddDate(0, 1, 0)
		dateEnd = now.EndOfMonth().AddDate(0, 1, 0)
	}
	dateStartString = dateStart.Format("2006-01-02")
	dateEndString = dateEnd.Format("2006-01-02")
	events := []model.Event{}
	db.DBConn.Select("id, title, place_title, date_start, date_end").Where("date_start > ? AND date_end < ?", dateStartString, dateEndString).Find(&events)
	return events
}

func GetEventsByCategory(categoryTitle string) []model.Event {
	events := []model.Event{}
	db.DBConn.Select("id, title, place_title, date_start, date_end").Where("category_title = ?", categoryTitle).Limit(5).Find(&events)
	return events
}

func GetEvent(eventID string) (model.Event, error) {
	var event model.Event
	if err := db.DBConn.Find(&event, eventID).Error; err != nil {
		return model.Event{}, errors.New("해당 밋업이 존재하지 않습니다")
	}
	users := []model.User{}
	db.DBConn.Model(&event).Association("Participants").Find(&users)
	event.Participants = users
	return event, nil
}

func GetParticipants(eventID string) ([]model.User, error) {
	var event model.Event
	if err := db.DBConn.Find(&event, eventID).Error; err != nil {
		return []model.User{}, errors.New("해당 밋업이 존재하지 않습니다")
	}
	users := []model.User{}
	db.DBConn.Model(&event).Association("Participants").Find(&users)
	return users, nil
}
