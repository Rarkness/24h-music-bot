package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/Rarkness/24h-music-bot/src/config"
	"github.com/Rarkness/24h-music-bot/src/player"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/diamondburned/arikawa/v3/state/store"
	"github.com/diamondburned/arikawa/v3/voice"
)

func init() { rand.Seed(time.Now().UnixNano()) }

func main() {
	s := state.New("Bot " + config.Token)
	s.Cabinet = store.NoopCabinet
	s.AddIntents(gateway.IntentGuilds)
	voice.AddIntents(s)

	if err := s.Open(context.Background()); err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Println("Logged in.")

	v, err := voice.NewSession(s)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if err := v.JoinChannelAndSpeak(context.Background(), config.VoiceChannelID, false, true); err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer s.Close()
	defer v.Leave(context.Background())

	go player.Run(v)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
