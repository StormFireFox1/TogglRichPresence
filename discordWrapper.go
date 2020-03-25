package TogglRichPresence

import (
	"github.com/hugolgst/rich-go/client"
	"log"
	"time"
)

type DiscordWrapper struct {
	currentActivity client.Activity
	appId           string
}

func (w DiscordWrapper) Initialize(appId string) {
	w.appId = appId
	err := client.Login(w.appId)
	if err != nil {
		log.Fatal(err)
	}
}

func (w DiscordWrapper) SetActivity(description, project string) {
	t := time.Now()
	w.currentActivity = client.Activity{
		Details:	project,
		State:      description,
		Timestamps: &client.Timestamps{
			Start: &t,
		},
	}
	err := client.SetActivity(w.currentActivity)
	if err != nil {
		log.Fatal(err)
	}
}

func (w DiscordWrapper) RefreshRichPresenceToggl(t TogglWrapper) {
	timeEntry := t.CurrentTimer()
	w.currentActivity = client.Activity{
		Details: timeEntry.project,
		State:   timeEntry.description,
		Timestamps: &client.Timestamps{
			Start: &timeEntry.startTime,
		},
	}
	err := client.SetActivity(w.currentActivity)
	if err != nil {
		log.Fatal(err)
	}
}