package user

import (
	"encoding/json"

	"github.com/meximonster/go-discordbot/content"
)

type User struct {
	content.ContentBase
}

func (u *User) Type() string {
	return "user"
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) AddImage(text string, url string) error {
	return content.AddImage("users", text, url)
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
	return content.Store("users", u.Name, images)
}
