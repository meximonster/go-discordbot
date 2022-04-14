package user

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/meximonster/go-discordbot/configuration"
)

var users = map[string]*User{}

type User struct {
	Username           string
	Id                 string
	ChannelID          string
	IsHuman            bool
	IsPet              bool
	Variety            bool
	Images             []configuration.ImageInfo
	LastImageURLServed string
}

func InitUsers(usrConfig []configuration.UserConfig) {
	for _, us := range usrConfig {
		u := new(User)
		u.Username = us.Username
		u.Id = us.UserID
		u.ChannelID = us.ChannelID
		u.IsHuman = us.IsHuman
		u.IsPet = us.IsPet
		u.Variety = us.Variety
		u.Images = us.Images

		users[u.Username] = u
	}
}

func GetUsers() map[string]*User {
	m := make(map[string]*User, len(users))
	for _, u := range users {
		if u.IsHuman {
			m[u.Username] = u
		}
	}
	return m
}

func GetPets() map[string]*User {
	m := make(map[string]*User, len(users))
	for _, u := range users {
		if u.IsPet {
			m[u.Username] = u
		}
	}
	return m
}

func GetByName(name string) (*User, error) {
	if u, ok := users[name]; ok {
		return u, nil
	}
	return nil, fmt.Errorf("user %s not found", name)
}

func (u *User) RandomImage() (configuration.ImageInfo, error) {
	if len(u.Images) == 0 {
		return configuration.ImageInfo{}, fmt.Errorf("no images for %s", u.Username)
	}
	if len(u.Images) == 1 {
		return u.Images[0]
	}
	var rng int
	rand.Seed(time.Now().UnixNano())
	flag := true
	for flag {
		rng = rand.Intn(len(u.Images))
		if u.Images[rng].Url != u.LastImageURLServed {
			flag = false
		}
	}
	u.LastImageURLServed = u.Images[rng].Url
	return u.Images[rng], nil
}
