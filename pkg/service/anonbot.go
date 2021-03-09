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
	glog.Infof("Message from %v to channel %v: %v", cmd.UserName, cmd.ChannelName, cmd.Text)
}

func (a *AnonBot) HandleCommand(cmd *slack.SlashCommand, w http.ResponseWriter) (err error) {
	a.logActivity(cmd)
	defer w.WriteHeader(http.StatusOK)
	if cmd.Text == "" {
		return
	}
	_, _, err = a.Api.PostMessage(cmd.ChannelID, slack.MsgOptionText(cmd.Text, false), slack.MsgOptionAsUser(true))
	if err != nil {
		glog.Errorf("Error writing message: %v", err)
	}
	return
}
