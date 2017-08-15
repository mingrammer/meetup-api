package web

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mingrammer/meetup-api/api/model"
	"github.com/mingrammer/meetup-api/config"
	"github.com/nlopes/slack"
)

func hookGeneratedEvent(event *model.Event) {
	client := slack.New(config.BotAPIConfig.SlackBot.BotToken)
	params := slack.PostMessageParameters{
		Attachments: []slack.Attachment{
			{
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
			},
		},
		AsUser: true,
	}
	if _, _, err := client.PostMessage(config.BotAPIConfig.SlackBot.ChannelID, "밋업이 생성되었습니다", params); err != nil {
		log.Fatalf("Failed to post message: %s", err)
	}
}
