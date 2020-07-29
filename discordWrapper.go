package TogglRichPresence

import (
	"log"
	"math/rand"
	"time"

	"github.com/hugolgst/rich-go/client"
)

type DiscordWrapper struct {
	currentActivity client.Activity
	appId           string
	defaultIcon     string
}

func InitializeDiscordWrapper(appId string, defaultIcon string) DiscordWrapper {
	var w DiscordWrapper
	w.appId = appId
	w.defaultIcon = defaultIcon
	err := client.Login(w.appId)
	if err != nil {
		log.Fatalf("Error logging in to Discord: %s. Are you sure Discord is running?", err)
	}
	return w
}

func (w DiscordWrapper) RestartWrapper() {
	client.Logout()
	err := client.Login(w.appId)
	if err != nil {
		log.Fatalf("Error restarting Discord: %s", err)
	}
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
	timeEntry, err := t.CurrentTimer()
	if err != nil {
		w.RestartWrapper()
		return
	}
	tags := ""
	iconId := w.defaultIcon
	if len(timeEntry.tags) != 0 {
		for _, tag := range timeEntry.tags {
			tags += "#" + tag + " "
		}
		iconId = timeEntry.tags[rand.Intn(len(timeEntry.tags))]
	}
	w.currentActivity = client.Activity{
		Details:    timeEntry.description,
		LargeImage: iconId,
		LargeText:  iconId,
		State:      "@" + timeEntry.project + " " + tags,
		Timestamps: &client.Timestamps{
			Start: &timeEntry.startTime,
		},
	}
	err = client.SetActivity(w.currentActivity)
	if err != nil {
		log.Fatalf("Error sending activity to Discord: %s", err)
	}
}
