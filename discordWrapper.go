package TogglRichPresence

import (
	"github.com/hugolgst/rich-go/client"
	"log"
	"math/rand"
	"time"
)

type DiscordWrapper struct {
	currentActivity client.Activity
	appId           string
}

func InitializeDiscordWrapper(appId string) DiscordWrapper {
	var w DiscordWrapper
	w.appId = appId
	err := client.Login(w.appId)
	if err != nil {
		log.Fatalf("Error logging in to Discord: %s. Are you sure Discord is running?", err)
	}
	return w
}

func (w DiscordWrapper) SetActivity(description, project string) {
	t := time.Now()
	w.currentActivity = client.Activity{
		Details: project,
		State:   description,
		Timestamps: &client.Timestamps{
			Start: &t,
		},
	}
	err := client.SetActivity(w.currentActivity)
	if err != nil {
		log.Fatalf("Error sending activity to Discord: %s", err)
	}
}

func (w DiscordWrapper) RefreshRichPresenceToggl(t TogglWrapper) {
	timeEntry := t.CurrentTimer()
	tags := ""
	if len(timeEntry.tags) != 0 {
		for _, tag := range timeEntry.tags {
			tags += "#" + tag + " "
		}
	}
	iconId := timeEntry.tags[rand.Intn(len(timeEntry.tags))]
	w.currentActivity = client.Activity{
		Details:    timeEntry.description,
		LargeImage: iconId,
		SmallImage: iconId,
		State:      "@" + timeEntry.project + " " + tags,
		Timestamps: &client.Timestamps{
			Start: &timeEntry.startTime,
		},
	}
	err := client.SetActivity(w.currentActivity)
	if err != nil {
		log.Fatalf("Error sending activity to Discord: %s", err)
	}
}
