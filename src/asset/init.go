package asset

import (
	"embed"
	"io/fs"
	"regexp"
	"strings"
)

var (
	//go:embed music
	dir       embed.FS
	filenames []string
)

func init() {
	fs.WalkDir(
		dir,
		".", func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				return nil
			}

			filenames = append(filenames, path)
			return nil
		})
}

func init() {
	reg := regexp.MustCompile(`[^\/](\w|[가-힣])+(\.)`)

	for _, filename := range filenames {
		if !strings.HasSuffix(filename, ".mp3") {
			continue
		}

		byt, err := fs.ReadFile(dir, filename)
		if err != nil {
			continue
		}

		name := strings.ReplaceAll(reg.FindStringSubmatch(filename)[0], ".", "")
		assets = append(assets, NewAsset().SetName(name).SetMusic(byt))
	}
}
