package asset

type Asset struct {
	name  string
	music []byte
}

func NewAsset() *Asset { return &Asset{} }

func (a *Asset) SetName(name string) *Asset {
	a.name = name
	return a
}

func (a *Asset) SetMusic(music []byte) *Asset {
	a.music = music
	return a
}

func (a *Asset) GetName() string  { return a.name }
func (a *Asset) GetMusic() []byte { return a.music }
