package bot

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/mingrammer/meetup-api/config"
	"github.com/nlopes/slack"
)

func listEventAttachment(params map[string]string) *slack.Attachment {
	date := params["date"]
	eventList := GetEventsByDate(date)
	var eventListStringBuffer bytes.Buffer
	if len(eventList) > 0 {
		for _, event := range eventList {
			eventListStringBuffer.WriteString(fmt.Sprintf(
				"%d. %s\n장소: %s\n시작: %s\n종료: %s\n링크: %s\n\n",
				event.ID,
				event.Title,
				event.PlaceTitle,
				event.DateStart.Format("2006-01-02 15:04"),
				event.DateEnd.Format("2006-01-02 15:04"),
				fmt.Sprintf("%s/details/%d", config.BotAPIConfig.WebEndpoint, event.ID),
			))
		}
	} else {
		eventListStringBuffer = *bytes.NewBufferString(fmt.Sprintf("%s에는 밋업이 없습니다", date))
	}
	return &slack.Attachment{
		Fallback: fmt.Sprintf("%s 밋업 리스트입니다", date),
		Title:    fmt.Sprintf("%s 밋업 리스트입니다", date),
		Text:     eventListStringBuffer.String(),
	}
}

func eventDetailAttachment(params map[string]string) *slack.Attachment {
	event, err := GetEvent(params["event_id"])
	if err != nil {
		return nil
	}
	return &slack.Attachment{
		Fields: []slack.AttachmentField{
			{
				Title: "밋업 번호",
				Value: strconv.Itoa(int(event.ID)),
			},
			{
				Title: "밋업명",
				Value: event.Title,
				Short: true,
			},
			{
				Title: "카테고리",
				Value: event.CategoryTitle,
				Short: true,
			},
			{
				Title: "개설자",
				Value: event.OwnerName,
				Short: true,
			},
			{
				Title: "참가자 수",
				Value: fmt.Sprintf("%d명", len(event.Participants)),
				Short: true,
			},
			{
				Title: "시작 시간",
				Value: event.DateStart.Format("2006-01-02 15:04"),
				Short: true,
			},
			{
				Title: "종료 시간",
				Value: event.DateEnd.Format("2006-01-02 15:04"),
				Short: true,
			},
			{
				Title: "장소",
				Value: event.PlaceTitle,
			},
			{
				Title: "링크",
				Value: fmt.Sprintf("%s/details/%d", config.BotAPIConfig.WebEndpoint, event.ID),
			},
		},
		Color: "good",
	}
}

func listCommandAttachment() *slack.Attachment {
	commandList := []string{
		"밋업 리스트 보기 - (오늘|이번주|다음주|이번달|다음달) 밋업",
		"밋업 상세 정보 보기 - {id}번 밋업 정보",
		"카테고리 리스트 - 모든 카테고리 보여줘 | 카테고리 리스트",
		"카테고리 검색 - {category}(으)로 밋업 검색",
		"참가자 보기 - {id}번 밋업 참가자",
		"명령어 리스트 - 모든 커맨드 보여줘 | 모든 명령어 보여줘 | 커맨드 리스트 | 명령어 리스트",
	}
	var commandListStringBuffer bytes.Buffer
	for _, command := range commandList {
		commandListStringBuffer.WriteString(fmt.Sprintf("%s\n", command))
	}
	return &slack.Attachment{
		Title: "명령어 리스트",
		Text:  commandListStringBuffer.String(),
	}
}

func listCategoryAttachment() *slack.Attachment {
	categoryList := GetCategories()
	var categoryListStringBuffer bytes.Buffer
	for _, category := range categoryList {
		categoryListStringBuffer.WriteString(fmt.Sprintf("%s\n", category.Title))
	}
	return &slack.Attachment{
		Title: "카테고리 리스트",
		Text:  categoryListStringBuffer.String(),
	}
}

func categorySearchAttachment(params map[string]string) *slack.Attachment {
	categoryTitle := params["category_title"]
	eventList := GetEventsByCategory(categoryTitle)
	var eventListStringBuffer bytes.Buffer
	if len(eventList) > 0 {
		for _, event := range eventList {
			eventListStringBuffer.WriteString(fmt.Sprintf(
				"%d. %s\n장소: %s\n시작: %s\n종료: %s\n링크: %s\n\n",
				event.ID,
				event.Title,
				event.PlaceTitle,
				event.DateStart.Format("2006-01-02 15:04"),
				event.DateEnd.Format("2006-01-02 15:04"),
				fmt.Sprintf("%s/details/%d", config.BotAPIConfig.WebEndpoint, event.ID),
			))
		}
	} else {
		eventListStringBuffer = *bytes.NewBufferString(fmt.Sprintf("%s 관련 밋업이 없습니다", categoryTitle))
	}
	return &slack.Attachment{
		Fallback: fmt.Sprintf("%s 관련 밋업 리스트입니다", categoryTitle),
		Title:    fmt.Sprintf("%s 관련 밋업 리스트입니다", categoryTitle),
		Text:     eventListStringBuffer.String(),
	}
}

func listParticipantsAttachment(params map[string]string) *slack.Attachment {
	participantList, err := GetParticipants(params["event_id"])
	if err != nil {
		return nil
	}
	var participantListStringBuffer bytes.Buffer
	if len(participantList) > 0 {
		for _, participant := range participantList {
			participantListStringBuffer.WriteString(fmt.Sprintf("%s\n", participant.Name))
		}
	} else {
		participantListStringBuffer = *bytes.NewBufferString("0명")
	}
	return &slack.Attachment{
		Title: fmt.Sprintf("참가자 리스트 (총 %d명)", len(participantList)),
		Text:  participantListStringBuffer.String(),
	}
}

func generateAttachment(command string) (string, *slack.Attachment, string) {
	match := cmdListEventRegex.FindStringSubmatch(command)
	if len(match) == 2 {
		return "밋업 리스트", listEventAttachment(map[string]string{
			"date": match[1],
		}), channelAny
	}
	match = cmdEventDetailRegex.FindStringSubmatch(command)
	if len(match) == 2 {
		return "밋업 상세 정보", eventDetailAttachment(map[string]string{
			"event_id": match[1],
		}), channelAny
	}
	match = cmdListParticipantsRegex.FindStringSubmatch(command)
	if len(match) == 2 {
		return "참가자 리스트", listParticipantsAttachment(map[string]string{
			"event_id": match[1],
		}), channelAny
	}
	match = cmdCategorySearchRegex.FindStringSubmatch(command)
	if len(match) == 2 {
		return "검색 결과", categorySearchAttachment(map[string]string{
			"category_title": match[1],
		}), channelAny
	}
	match = cmdListCommandRegex.FindStringSubmatch(command)
	if len(match) == 1 {
		return "명령어 리스트", listCommandAttachment(), channelAny
	}
	match = cmdListCategoryRegex.FindStringSubmatch(command)
	if len(match) == 1 {
		return "카테고리 리스트", listCategoryAttachment(), channelAny
	}
	return "", nil, ""
}
