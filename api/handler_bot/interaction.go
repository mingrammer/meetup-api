package bot

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"strings"

// 	"github.com/ant0ine/go-json-rest/rest"
// 	"github.com/mingrammer/meetup-api/config"
// 	"github.com/nlopes/slack"
// )

// func BotResponder(w rest.ResponseWriter, r *rest.Request) {
// 	if r.Method != http.MethodPost {
// 		log.Printf("[ERROR] Invalid method: %s", r.Method)
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		return
// 	}
// 	buf, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Printf("[ERROR] Failed to read request body: %s", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	jsonStr, err := url.QueryUnescape(string(buf)[8:])
// 	if err != nil {
// 		log.Printf("[ERROR] Failed to unespace request body: %s", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	message := slack.AttachmentActionCallback{}
// 	if err := json.Unmarshal([]byte(jsonStr), &message); err != nil {
// 		log.Printf("[ERROR] Failed to decode json message from slack: %s", jsonStr)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	if message.Token != config.BotAPIConfig.SlackBot.VerificationToken {
// 		log.Printf("[ERROR] Invalid token: %s", message.Token)
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	action := message.Actions[0]
// 	switch action.Name {
// 	case actionSelect:
// 		value := action.SelectedOptions[0].Value
// 		// Overwrite original drop down message.
// 		originalMessage := message.OriginalMessage
// 		originalMessage.Attachments[0].Text = fmt.Sprintf("OK to order %s ?", strings.Title(value))
// 		originalMessage.Attachments[0].Actions = []slack.AttachmentAction{
// 			{
// 				Name:  actionStart,
// 				Text:  "Yes",
// 				Type:  "button",
// 				Value: "start",
// 				Style: "primary",
// 			},
// 			{
// 				Name:  actionCancel,
// 				Text:  "No",
// 				Type:  "button",
// 				Style: "danger",
// 			},
// 		}

// 		w.Header().Add("Content-type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w.(http.ResponseWriter)).Encode(&originalMessage)
// 		return
// 	case actionStart:
// 		title := ":ok: your order was submitted! yay!"
// 		responseMessage(w, message.OriginalMessage, title, "")
// 		return
// 	case actionCancel:
// 		title := fmt.Sprintf(":x: @%s canceled the request", message.User.Name)
// 		responseMessage(w, message.OriginalMessage, title, "")
// 		return
// 	default:
// 		log.Printf("[ERROR] ]Invalid action was submitted: %s", action.Name)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// }

// // responseMessage response to the original slackbutton enabled message.
// // It removes button and replace it with message which indicate how bot will work
// func responseMessage(w rest.ResponseWriter, original slack.Message, title, value string) {
// 	original.Attachments[0].Actions = []slack.AttachmentAction{} // empty buttons
// 	original.Attachments[0].Fields = []slack.AttachmentField{
// 		{
// 			Title: title,
// 			Value: value,
// 			Short: false,
// 		},
// 	}
// 	w.Header().Add("Content-type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w.(http.ResponseWriter)).Encode(&original)
// }
