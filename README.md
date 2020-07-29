# TogglRichPresence

A Go library (along with simple command-line app) that displays any active Toggl
timer as a Rich Presence status in your Discord account.

A little side-project made while working from home to avoid having to tell
people what I'm doing on Discord.

## Installation

If you want to work with the library, just import the package in your Go code:

```go
import "github.com/StormFireFox1/TogglRichPresence"
```

If you don't want to develop, and just want to use it, `go get` the binary
package, and it will be installed in your PATH:
```bash
$ go get "github.com/StormFireFox1/TogglRichPresence/cmd/togglRichPresence"
```

## Configuration

The library requires a bit of configuration on your part, preferably done using
environment variables. Alternatively, if the environment variables are not set,
TogglRichPresence reads command-line flags. Otherwise, the binary fails.

There are a few extra ones that are optional, but useful:
- `DEFAULT_ICON_ID`: The default icon to display in TogglRichPresence on empty
  timer entries. By default, TogglRichPresence displays no image if the timer
  entry does not have any tags that could be used for Rich Presence icons.
  Assign a string here to change that. The string must match the name of the
  Rich Presence icon available in the Discord Developer Portal.
- `DEFAULT_REFRESH_INTERVAL`: By default, TogglRichPresence refreshes the Rich
  Presence information every 10 seconds. Use this environment variable to change
  the refresh interval. Default is 10 seconds.

## Usage

The most important thing is to have Discord running in the background. Sadly,
the Discord Rich Presence library requires Discord running on a computer (or
mobile phone) in order to read the info of the Rich Presence app.

### Discord App ID
You will need the ID of a Discord custom app you create, in order for your
account to be able to use Rich Presence on your account:

- Go to the Discord Developer Portal (discordapp.com/developrs/applications)
- Create a "New Application"
- Give it a name of your choosing (I chose "The Work Game")
- Give it an app icon, if you want
- Add a Rich Presence cover image and any assets you may want (give them names
  that relate to your Toggl tags)
- Copy the Client ID of your new custom app (That's the DISCORD_APP_ID parameter)

### Toggl API Key
You will need your Toggl API key, this can be found under "My Profile" in your
Toggl window.

After the above configuration you can run the command in the following manner:
```shelll
$ togglRichPresence sync \
  --discordAppId [Discord App ID here] \
  --togglApiKey [Toggl API Key here]
```

Alternatively, set the above parameters as environment variables.

While the command is running, it will automatically sync your most recent Toggl
timer with Discord Rich Presence. It will display the description of the time
entry, the project under the format "@ProjectName" and it will also display tags
under the format "#tag1 #tag2 ..."

Additionally, a random tag is chosen in order to select an app icon to be
displayed for the Rich Presence of the time entry.
