package player

import (
	"github.com/Rarkness/24h-music-bot/src/asset"
	"github.com/Rarkness/24h-music-bot/src/config"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

var messageID discord.MessageID

func sendMessage(a asset.Asset) {
	client := api.NewClient("Bot " + config.Token)

	embed := embedBuilder(a)
	comp := componentsBuilder()

	if messageID.IsValid() {
		if _, err := client.EditMessageComplex(
			config.VoiceChannelID,
			messageID,
			api.EditMessageData{
				Embeds:     &[]discord.Embed{embed},
				Components: &comp}); err != nil {
			messageID = 0
		}
		return
	}

	if msg, err := client.SendMessageComplex(
		config.VoiceChannelID,
		api.SendMessageData{
			Embeds:     []discord.Embed{embed},
			Components: comp}); err == nil {
		messageID = msg.ID
	}
}

func embedBuilder(a asset.Asset) (embed discord.Embed) {
	embed.Title = "Now Playing"
	embed.Description = a.GetName()
	return
}

func componentsBuilder() discord.ContainerComponents {
	return discord.ContainerComponents{
		&discord.ActionRowComponent{
			&discord.ButtonComponent{
				Style: discord.LinkButtonStyle("https://github.com/Rarkness/24h-music-bot"),
				Label: "See Repo",
			}}}
}
