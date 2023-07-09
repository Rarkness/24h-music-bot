package player

import (
	"bytes"

	"github.com/Rarkness/24h-music-bot/src/asset"
	"github.com/jonas747/dca"
)

func encode(a asset.Asset) (*dca.EncodeSession, error) {
	opt := dca.StdEncodeOptions
	opt.RawOutput = true
	opt.Bitrate = 96
	opt.Application = "lowdelay"

	encoder, err := dca.EncodeMem(bytes.NewReader(a.GetMusic()), opt)

	if err != nil {
		return nil, err
	}

	return encoder, nil
}
