package content

import (
	"fmt"

	"github.com/meximonster/go-discordbot/content/emote"
	"github.com/meximonster/go-discordbot/content/pet"
	"github.com/meximonster/go-discordbot/content/user"
	"github.com/meximonster/go-discordbot/image"
)

var Cnt map[string]Content

type Content interface {
	Type() string
	GetName() string
	AddImage(text string, url string) error
	RandomImage() (image.Image, error)
	Store() error
}

func Load() error {
	users := user.GetAll()
	for _, u := range users {
		o := &u
		Cnt[u.Name] = o
	}
	pets := pet.GetAll()
	for _, p := range pets {
		o := &p
		Cnt[p.Name] = o
	}
	emotes := emote.GetAll()
	for _, e := range emotes {
		o := &e
		Cnt[e.Name] = o
	}
	return nil
}

func Get() map[string]Content {
	return Cnt
}

func GetOne(name string) (Content, error) {
	if _, ok := Cnt[name]; !ok {
		return nil, fmt.Errorf("%s doesn't exist", name)
	}
	return Cnt[name], nil
}

func List(contentType string) []string {
	var s []string
	for k, v := range Cnt {
		if contentType == v.Type() {
			s = append(s, k)
		}
	}
	return s
}
