package content

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/meximonster/go-discordbot/content/emote"
	"github.com/meximonster/go-discordbot/content/pet"
	"github.com/meximonster/go-discordbot/content/user"
	"github.com/meximonster/go-discordbot/image"
)

var Cnt = make(map[string]Content)

type Content interface {
	Type() string
	GetName() string
	AddImage(text string, url string) error
	RandomImage() (image.Image, error)
	Store() error
}

func Load() error {
	users, err := user.GetAll()
	if err != nil {
		return err
	}
	for _, u := range users {
		Cnt[u.Alias] = u
	}
	pets, err := pet.GetAll()
	if err != nil {
		return err
	}
	for _, p := range pets {
		Cnt[p.Alias] = p
	}
	emotes, err := emote.GetAll()
	if err != nil {
		return err
	}
	for _, e := range emotes {
		Cnt[e.Alias] = e
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

func Set(name string, contentType string) error {
	var c Content
	switch contentType {
	case "user":
		users := List(contentType)
		for _, u := range users {
			if u == name {
				return fmt.Errorf("user %s already exists", name)
			}
		}
		c = &user.User{Alias: name, Images: []byte{}}
	case "pet":
		pets := List(contentType)
		for _, p := range pets {
			if p == name {
				return fmt.Errorf("pet %s already exists", name)
			}
		}
		c = &pet.Pet{Alias: name, Images: []byte{}}
	case "emote":
		emotes := List(contentType)
		for _, e := range emotes {
			if e == name {
				return fmt.Errorf("emote %s already exists", name)
			}
		}
		c = &emote.Emote{Alias: name, Images: []byte{}}
	}
	Cnt[c.GetName()] = c
	return c.Store()
}

func AddImage(c Content, text string, url string) error {
	return c.AddImage(text, url)
}

func RandomImage(c Content) (image.Image, error) {
	return c.RandomImage()
}

func NewDB(db *sqlx.DB) {
	user.NewDB(db)
	pet.NewDB(db)
	emote.NewDB(db)
}

func CloseDB() {
	user.CloseDB()
	pet.CloseDB()
	emote.CloseDB()
}
