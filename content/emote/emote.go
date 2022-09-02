package emote

import (
	"encoding/json"

	"github.com/meximonster/go-discordbot/content"
)

type Emote struct {
	content.ContentBase
}

func (e *Emote) Type() string {
	return "emote"
}

func (e *Emote) GetName() string {
	return e.Name
}

func (e *Emote) AddImage(text string, url string) error {
	return content.AddImage("emotes", text, url)
}

func (e *Emote) RandomImage(text string, url string) (content.Image, error) {
	img, err := content.RandomImage(e.Images, e.LastImageURLServed)
	if err != nil {
		return content.Image{}, err
	}
	e.LastImageURLServed = img.Url
	return img, nil
}

func (e *Emote) Store() error {
	images, err := json.Marshal(e.Images)
	if err != nil {
		return err
	}
	return content.Store("users", e.Name, images)
}
