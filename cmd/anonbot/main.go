package main

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/matthewlang/anonbot/pkg/service"

	"github.com/golang/glog"
	"github.com/slack-go/slack"

	"flag"
)

var (
	oauth         string // OAuth token
	signingSecret string // Application signing secret
	clientSecret  string // Application client secret
	port          string // Port to listen on
	cmdUrl        string // URL to receive interactions
	cmd           string // Slash command for sending messages
)

func slashify(s string) string {
	if s[0] != '/' {
		return "/" + s
	}
	return s
}

func main() {
	flag.StringVar(&oauth, "oauth", "", "OAuth Token")
	flag.StringVar(&signingSecret, "ssecret", "", "Application signing secret")
	flag.StringVar(&clientSecret, "csecret", "", "Application client secret")
	flag.StringVar(&port, "p", ":1000", "Port to listen on")
	flag.StringVar(&cmdUrl, "cmdUrl", "/slash", "URL to receive slash commands")
	flag.StringVar(&cmd, "cmd", "/anon", "Name of slash command")

	flag.Parse()

	glog.Infof("Starting on port %v ...", port)

	cmd = slashify(cmd)
	cmdUrl = slashify(cmdUrl)

	api := slack.New(oauth)

	ab := service.AnonBot{Api: api}
	http.HandleFunc(cmdUrl, func(w http.ResponseWriter, r *http.Request) {
		verifier, err := slack.NewSecretsVerifier(r.Header, signingSecret)
		if err != nil {
			glog.Errorf("Could not create verifier: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		r.Body = ioutil.NopCloser(io.TeeReader(r.Body, &verifier))
		s, err := slack.SlashCommandParse(r)
		if err != nil {
			glog.Errorf("Unauthorized: %v", err)
			return
		}
		glog.V(1).Infof("Command parsed as %v for %v", s.Command, s)

		if s.Command == cmd {
			ab.HandleCommand(&s, w)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})

	glog.Infof("Listening...")
	http.ListenAndServe(port, nil)
}
