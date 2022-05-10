package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
	"disco-bot/controller"
)

func main() {
	token, err := controller.BotToken(".discord.token")
	if err != nil {
		log.Fatal(err)
	}
	dis, err := discordgo.New(token)
	if err != nil {
		log.Fatal(err)
	}

	dis.AddHandler(controller.MessageCreate)
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
