package asset

import (
	"math/rand"
)

var (
	assets        = make([]*Asset, 0)
	sequenceIndex = 0
)

func GetAssetWithRandIndex() *Asset {
	index := rand.Intn(len(assets))
	return assets[index]
}

func GetAssetWithSequence() *Asset {
	defer func() {
		sequenceIndex++
		if sequenceIndex >= len(assets) {
			sequenceIndex = 0
		}
	}()
	return assets[sequenceIndex]
}
