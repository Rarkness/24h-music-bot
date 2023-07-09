package config

import (
	"log"
	"os"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/joho/godotenv"
)

var (
	Token          string
	VoiceChannelID discord.ChannelID
)

func init() {
	godotenv.Load()
	Token = os.Getenv("TOKEN")
}

func init() {
	chID := os.Getenv("VOICE_CHANNEL_ID")
	rawID, err := discord.ParseSnowflake(chID)
	if err != nil {
		log.Fatal("voice channel id is not in snowflake format.")
	}

	VoiceChannelID = discord.ChannelID(rawID)
}
