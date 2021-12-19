package main
 
import (
    "log"
 
    "github.com/slack-go/slack"
)
 
func main() {
 
    OAUTH_TOKEN := "OAUTH_TOKEN"
    CHANNEL_ID := "CHANNEL_ID"
 
    api := slack.New(OAUTH_TOKEN)
    attachment := slack.Attachment{
        Pretext: "A message from Josh",
        Text:    "You never see me coming!",
    }
 
    channelId, timestamp, err := api.PostMessage(
        CHANNEL_ID,
        slack.MsgOptionText("This is the main message", false),
        slack.MsgOptionAttachments(attachment),
        slack.MsgOptionAsUser(true),
    )
 
    if err != nil {
        log.Fatalf("%s\n", err)
    }
 
    log.Printf("Message successfully sent to Channel %s at %s\n", channelId, timestamp)
}
