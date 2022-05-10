package controller

import (
	"github.com/bwmarrin/discordgo"
	"os"
	"strings"
	"fmt"
	"regexp"
)

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
	c := []byte(strings.TrimSpace(strings.ToLower(m.Content)))
	cuddleMatched, err := regexp.Match(`^cuddle me$`, c)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	sheSaidMatched, err := regexp.Match(`^it['s]* so (big|hard|huge).*$`, c)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	switch {
	case sheSaidMatched:
		s.ChannelMessageSend(m.ChannelID, "That's what she said!")
	case cuddleMatched:
		s.ChannelMessageSend(m.ChannelID, "aww, cuddles for you.")
	default:
		return
	}
}
