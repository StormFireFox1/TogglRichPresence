package main

import (
	"fmt"
	"github.com/StormFireFox1/TogglRichPresence"
	"log"
	"os"
	"os/signal"
	"sort"
	"syscall"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:     "togglApiKey",
			Value:    "",
			Usage:    "API `KEY` for Toggl account",
			EnvVars:  []string{"TOGGL_API_KEY"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "discordAppId",
			Value:    "",
			Usage:    "App `ID` for personal Discord Rich Presence app",
			EnvVars:  []string{"DISCORD_APP_ID"},
			Required: true,
		},
		&cli.StringFlag{
			Name:    "defaultIconId",
			Value:   "",
			Usage:   "Default name for icon to use if no tags present in time entry. Empty by default.",
			EnvVars: []string{"DEFAULT_ICON_ID"},
		},
		&cli.IntFlag{
			Name:    "defaultRefreshInterval",
			Value:   10,
			Usage:   "Default refresh interval for TogglRichPresence to refresh the Rich Presence with the Toggl API. Default is 10 seconds.",
			EnvVars: []string{"DEFAULT_REFRESH_INTERVAL"},
		},
		&cli.StringFlag{
			Name:    "config, c",
			Value:   "~/.config/TogglRichPresence/config.json",
			Usage:   "Load configuration from `FILE`",
			EnvVars: []string{"CONFIG_PATH"},
		},
	}

	app := &cli.App{
		Flags: flags,
		Commands: []*cli.Command{
			{
				Name:    "sync",
				Aliases: []string{"s"},
				Usage:   "Interactively syncs current time entry in Toggl to Discord Rich Presence.",
				Action: func(c *cli.Context) error {
					sigChannel := make(chan os.Signal)
					signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)
					go func() {
						<-sigChannel
						fmt.Println("Interrupt received! Stopping refresh...")
						os.Exit(1)
					}()

					discordWrapper := TogglRichPresence.InitializeDiscordWrapper(c.String("discordAppId"), c.String("defaultIconId"))
					togglWrapper := TogglRichPresence.InitializeTogglWrapper(c.String("togglApiKey"))

					log.Print("Syncing Toggl and Discord Rich Presence!")
					for {
						discordWrapper.RefreshRichPresenceToggl(togglWrapper)
						time.Sleep(time.Duration(c.Int("defaultRefreshInterval")) * time.Second)
					}
				},
			},
			{
				Name:    "set",
				Aliases: []string{"a"},
				Usage: "Stops current timer, sets a new Toggl timer to the provided `DESCRIPTION`, `PROJECT`" +
					"and `DETAILS`, and also sets the current Rich Presence activity.",
				Action: func(c *cli.Context) error {
					log.Fatal("Not yet implemented!")
					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
