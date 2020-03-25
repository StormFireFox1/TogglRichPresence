# TogglRichPresence

A Go library (along with simple command-line app) that displays any active Toggl timer as a Rich Presence status in your Discord account.

A little side-project made while working from home to avoid having to tell people what I'm doing on Discord.

## Installation

If you want to work with the library, just import the package in your Go code:

```go
import "github.com/StormFireFox1/TogglRichPresence"
```

If you don't want to develop, and just want to use it, `go get` the binary package, and it will be installed in your PATH:
```shell
$ go get github.com/StormFireFox1/TogglRichPresence/cmd"
```

## Configuration

The library requires a bit of configuration on your part, preferably done using environment variables. Alternatively, if the environment
variables are not set, TogglRichPresence reads the file in the directory "~/.config/TogglRichPresence/config.json".

## Usage



