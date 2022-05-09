package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
	"strings"
)

func main() {
	data, err := os.ReadFile(".discord.token")
	if err != nil {
		log.Fatal(err)
	}
	token := strings.TrimSpace(string(data))
	log.Printf("using token: %s", token)
	dis, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	dis.AddHandler(messageCreate)
	dis.Identify.Intents = discordgo.IntentsGuildMessages
	err = dis.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer dis.Close()

	log.Printf("connection established.")
	// wait for kill
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
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
