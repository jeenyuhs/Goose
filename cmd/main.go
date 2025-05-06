package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/jeenyuhs/Goose/internal/handlers"
	"github.com/jeenyuhs/Goose/internal/repository"
)


func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("BOT_TOKEN variable not set.")
	}

	server := repository.NewServer()

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}
	
	discord.AddHandler(func(session *discordgo.Session, ready *discordgo.Ready) {
		fmt.Println("Bot is now running. Press CTRL+C to exit")
	})
	
	messageHandler := handlers.NewMessageHandler(server)
	discord.AddHandler(messageHandler.HandleMessages)
	
	err = discord.Open()
	if err != nil {
		panic(err)
	}

	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan, os.Interrupt, syscall.SIGTERM)
	<-signal_chan

	fmt.Println("Gracefully shutting down...")
	
	err = discord.Close()
	if err != nil {
		panic(err)
	}
}