package main

import (
	"encoding/json"
	"log"
	"os/exec"
	"strings"
)

type Command struct {
	Command   string   `json:"command"`
	Arguments []string `json:"args"`
}

const ScriptStart = "tell application \"Spotify\" to "

var commands = map[string]string{
	"state":      "player state",
	"play":       "play",
	"pause":      "pause",
	"duration":   "duration of current track",
	"name":       "name of current track",
	"album":      "album of current track",
	"id":         "id of current track",
	"artwork":    "artwork of current track",
	"vol_loud":   "set sound volume to 100",
	"vol_soft":   "set sound volume to 20",
	"vol_norm":   "set sound volume to 50",
	"set_volume": "set sound volume to ", //requires parameter
	"play_track": "play track ",          //requires parameter
	"position":   "player position",
}

func callSpotify(command string, args ...string) string {
	fullcmd := ScriptStart + commands[command] + strings.Join(args, " ")

	out, err := exec.Command("/usr/bin/osascript", "-e", fullcmd).Output()
	if err != nil {
		log.Fatal(err)
		log.Fatal(out)
	}
	return strings.TrimSpace(string(out))
}

func parseCommand(cmd []byte) error {

	var c Command
	err := json.Unmarshal(cmd, &c)
	if err != nil {
		return err
	}

	// Action Command
	log.Printf("Command: %s", c.Command)
	callSpotify(c.Command, c.Arguments...)

	return nil
}
