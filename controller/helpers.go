package controller

import (
	"github.com/bwmarrin/discordgo"
	"os"
	"strings"
)

func FetchToken() (string, error) {
	data, err := os.ReadFile(".discord.token")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
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
	case "cuddle me":
		s.ChannelMessageSend(m.ChannelID, "aww, cuddles for you.")
	default:
		return
	}
}
