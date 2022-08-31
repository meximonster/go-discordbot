package pet

import (
	"encoding/json"

	"github.com/meximonster/go-discordbot/content"
)

type Pet struct {
	content.ContentBase
}

func (p *Pet) Type() string {
	return "pet"
}

func (p *Pet) AddImage(text string, url string) error {
	img := content.Image{
		Text: text,
		Url:  url,
	}
	image, err := json.Marshal(img)
	if err != nil {
		return err
	}
	return content.AddImages("users", p.Name, string(image))
}

func (p *Pet) RandomImage(text string, url string) (content.Image, error) {
	img, err := content.RandomImage(p.Images, p.LastImageURLServed)
	if err != nil {
		return content.Image{}, err
	}
	p.LastImageURLServed = img.Url
	return img, nil
}

func (p *Pet) Store() error {
	images, err := json.Marshal(p.Images)
	if err != nil {
		return err
	}
	return content.Store("users", p.Name, images, "")
}
