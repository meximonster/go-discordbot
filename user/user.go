package user

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/meximonster/go-discordbot/configuration"
)

var users = map[string]*User{}

type User struct {
	Username  string
	Id        string
	ChannelID string
	Images    []configuration.ImageInfo
}

func InitUsers(usrConfig []configuration.UserConfig) {
	for _, us := range usrConfig {
		u := new(User)
		u.Username = us.Username
		u.Id = us.UserID
		u.ChannelID = us.ChannelID
		u.Images = us.Images

		users[u.Username] = u
	}
}

func GetUserByName(name string) *User {
	if u, ok := users[name]; ok {
		return u
	}
	return nil
}

func (u *User) RandomImage() (configuration.ImageInfo, error) {
	if len(u.Images) == 0 {
		return configuration.ImageInfo{}, fmt.Errorf("no images for %s", u.Username)
	}
	rand.Seed(time.Now().UnixNano())
	rng := rand.Intn(len(u.Images))
	return u.Images[rng], nil
}
