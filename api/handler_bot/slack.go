package bot

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/nlopes/slack"
)

const (
	channelIMOnly = "im_only_channel"
	channelAny    = "any_channel"
)

var (
	cmdListEventRegex        = regexp.MustCompile(`(오늘|이번주|다음주|이번달|다음달) 밋업`)
	cmdEventDetailRegex      = regexp.MustCompile(`(?P<id>\d+)번 밋업 정보`)
	cmdListCategoryRegex     = regexp.MustCompile(`모든 카테고리 보여줘|카테고리 리스트`)
	cmdCategorySearchRegex   = regexp.MustCompile(`(?P<category>[가-힣\w]+)으?로 밋업 검색`)
	cmdListParticipantsRegex = regexp.MustCompile(`(?P<id>\d+)번 밋업 참가자`)
	cmdListCommandRegex      = regexp.MustCompile(`모든 커맨드 보여줘|모든 명령어 보여줘|커맨드 리스트|명령어 리스트`)
)

type SlackListener struct {
	Client    *slack.Client
	BotID     string
	ChannelID string
}

// ListenAndResponse starts RTM connection and listens the incoming messages from slack
func (s *SlackListener) ListenAndResponse() {
	rtm := s.Client.NewRTM()

	// Start listening slack events
	go rtm.ManageConnection()

	// Handle slack events
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if err := s.handleMessageEvent(ev); err != nil {
				log.Printf("[ERROR] Failed to handle message: %s", err)
			}
		}
	}
}

// handleMesageEvent parses the messages and send the proper attachments to specific channels
func (s *SlackListener) handleMessageEvent(ev *slack.MessageEvent) error {
	// Only response mention to bot. Ignore else.
	if !strings.HasPrefix(ev.Msg.Text, fmt.Sprintf("<@%s> ", s.BotID)) {
		return nil
	}
	withoutPrefix := strings.TrimPrefix(ev.Msg.Text, fmt.Sprintf("<@%s> ", s.BotID))
	title, attachment, channelType := generateAttachment(withoutPrefix)
	if attachment == nil {
		return nil
	}
	params := slack.PostMessageParameters{
		Attachments: []slack.Attachment{
			*attachment,
		},
		AsUser: true,
	}
	channel := selectChannel(ev, channelType)
	if _, _, err := s.Client.PostMessage(channel, title, params); err != nil {
		return fmt.Errorf("Failed to post message: %s", err)
	}
	return nil
}

// selectChannel selects a channel from message event corresponding to channel type
func selectChannel(ev *slack.MessageEvent, channelType string) string {
	switch channelType {
	case channelIMOnly:
		return ev.User
	case channelAny:
		return ev.Channel
	default:
		return ""
	}
}
