package commands

import (
	"strings"

	"github.com/m0thm4n/server-spin/general"
	"github.com/m0thm4n/server-spin/scripts"

	"time"

	"github.com/bwmarrin/discordgo"
)

// Games stores the games for the game picker

//ExecuteCommand Parses and executes the command from the calling code
func ExecuteCommand(s *discordgo.Session, m *discordgo.Message, T0 time.Time) {
	var Games []string

	msg := strings.Split(strings.TrimSpace(m.Content), "!")[1]

	if len(msg) > 2 {
		msg = strings.Split(strings.Split(m.Content, " ")[0], "!")[1]
	}

	switch msg {
	case "info":
		general.HandleInfoCommand(s, m, T0)
	case "ping":
		general.HandlePingCommand(s, m)
	case "play":
		if !strings.EqualFold(m.Author.Username, "taft") {
			HandleWrongPermissions(s, m, msg)
			return
		}

		game := strings.Join(strings.Split(m.Content, " ")[1:], " ")
		general.HandlePlayCommand(s, game)
	case "reload":
		general.HandleReloadCommand(s, m)
	case "purge":

	case "createserver":
		image := strings.Join(strings.Split(m.Content, " ")[1:2], " ")
		name := strings.Join(strings.Split(m.Content, " ")[2:3], " ")

		scripts.CreateNewContainer(image, name)
	}
}

//HandleUnknownCommand is the default for any commands not listed
func HandleUnknownCommand(s *discordgo.Session, m *discordgo.Message, msg string) {

	c, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		println("Unable to open User Channel: ", err)
		return
	}
	s.ChannelMessageSend(c.ID, "The command ` "+msg+" ` is not recognized.")
}

func HandleWrongPermissions(s *discordgo.Session, m *discordgo.Message, msg string) {

	c, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		println("Unable to open User Channel: ", err)
		return
	}
	s.ChannelMessageSend(c.ID, "The command ` "+msg+" ` is not available to you.")
}
