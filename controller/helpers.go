package controller

import (
	"github.com/bwmarrin/discordgo"
	"os"
	"strings"
	"fmt"
	"regexp"
)

var sheSaidRegex string = `^it['s]* so (big|hard|huge).*$`

func BotToken(tokenFile string) (string, error) {
	data, err := os.ReadFile(tokenFile)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Bot %s", strings.TrimSpace(string(data))), nil
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore self messages
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case "ping":
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	case "pong":
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	default:
		defaultContentHandler(s, m)
	}
}

func defaultContentHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	c := strings.TrimSpace(strings.ToLower(m.Content))
	switch {
	case matches(`^cuddle me$`, c):
		s.ChannelMessageSend(m.ChannelID, "aww, cuddles for you.")
	case matches(sheSaidRegex, c):
		s.ChannelMessageSend(m.ChannelID, "That's what she said!")
	default:
		return
	}
}

func matches(reg, pattern string) bool {
	re := regexp.MustCompile(reg)
	return re.MatchString(pattern)
}
