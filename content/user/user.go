package user

import (
	"encoding/json"

	"github.com/meximonster/go-discordbot/content"
)

type User struct {
	content.ContentBase
	DiscordID string
}

func (u *User) Type() string {
	return "user"
}

func (u *User) AddImage(text string, url string) error {
	img := content.Image{
		Text: text,
		Url:  url,
	}
	image, err := json.Marshal(img)
	if err != nil {
		return err
	}
	return content.AddImages("users", u.Name, string(image))
}

func (u *User) RandomImage(text string, url string) (content.Image, error) {
	img, err := content.RandomImage(u.Images, u.LastImageURLServed)
	if err != nil {
		return content.Image{}, err
	}
	u.LastImageURLServed = img.Url
	return img, nil
}

func (u *User) Store() error {
	images, err := json.Marshal(u.Images)
	if err != nil {
		return err
	}
	return content.Store("users", u.Name, images, u.DiscordID)
}
