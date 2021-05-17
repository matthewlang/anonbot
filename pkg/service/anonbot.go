package service

import (
	"github.com/golang/glog"
	"github.com/slack-go/slack"

	"net/http"
)

type AnonBot struct {
	Api *slack.Client
}

func (a *AnonBot) logActivity(cmd *slack.SlashCommand) {
	glog.Infof("Message from %v to channel %v (%v): %v", cmd.UserName, cmd.ChannelName, cmd.ChannelID, cmd.Text)
}

func (a *AnonBot) HandleCommand(cmd *slack.SlashCommand, w http.ResponseWriter) (err error) {
	a.logActivity(cmd)
	if cmd.Text == "" {
		msg := "You need to supply text to post..."
		w.Write([]byte(msg))
		return
	}
	_, _, err = a.Api.PostMessage(cmd.ChannelID, slack.MsgOptionText(cmd.Text, false), slack.MsgOptionAsUser(true))
	if err != nil {
		glog.Errorf("Error writing message: %v", err)
		msg := "Oh no, I could't post that. Is this a private channel or group message? If so, you need to invite me."
		w.Write([]byte(msg))
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
