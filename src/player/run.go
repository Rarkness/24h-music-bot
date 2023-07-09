package player

import (
	"github.com/Rarkness/24h-music-bot/src/asset"
	"github.com/Rarkness/24h-music-bot/src/config"
	"github.com/diamondburned/arikawa/v3/voice"
)

func Run(v *voice.Session) {
	for {
		playMusic(v)
	}
}

func playMusic(v *voice.Session) {
	a := map[bool]*asset.Asset{
		true:  asset.GetAssetWithRandIndex(),
		false: asset.GetAssetWithSequence(),
	}[config.Shuffle]

	encoder, err := encode(*a)
	if err != nil {
		return
	}

	go sendMessage(*a)
	defer encoder.Cleanup()

	stream(v, encoder)
}
