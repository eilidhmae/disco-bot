package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
	"discord-bot/controller"
)

func main() {
	token, err := controller.FetchToken()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("using token: %s", token)
	dis, err := discordgo.New("Bot " + token)
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
