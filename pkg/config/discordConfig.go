package config

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var DiscordService *discordgo.Session
var DiscordServerID = ""

func connectDiscord() (*discordgo.Session, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Unable to read .env file")
	}
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_API_KEY"))
	if err != nil {
		log.Fatal("Error creating Discord session:", err)
		return dg, err
	}
	return dg, nil
}

func init() {
	var err error
	DiscordService, err = connectDiscord()
	if err != nil {
		log.Panic("Unable to Connect to Discord Server")
		return
	}
	log.Println("Connected to Discord Server")
	DiscordServerID = os.Getenv("DISCORD_SERVER_ID")
}
