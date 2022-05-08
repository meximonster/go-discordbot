package content

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/meximonster/go-discordbot/configuration"
)

var cnt = map[string]*Content{}

type Content struct {
	Name               string
	Id                 string
	ChannelID          string
	IsHuman            bool
	IsPet              bool
	Variety            bool
	Images             []configuration.ImageInfo
	LastImageURLServed string
}

func InitCnt(cntConfig []configuration.CntConfig) {
	for _, cfg := range cntConfig {
		c := new(Content)
		c.Name = cfg.Name
		c.Id = cfg.UserID
		c.ChannelID = cfg.ChannelID
		c.IsHuman = cfg.IsHuman
		c.IsPet = cfg.IsPet
		c.Variety = cfg.Variety
		c.Images = cfg.Images

		cnt[c.Name] = c
	}
}

func GetUsers() map[string]*Content {
	m := make(map[string]*Content, len(cnt))
	for _, u := range cnt {
		if u.IsHuman {
			m[u.Name] = u
		}
	}
	return m
}

func GetPets() map[string]*Content {
	m := make(map[string]*Content, len(cnt))
	for _, u := range cnt {
		if u.IsPet {
			m[u.Name] = u
		}
	}
	return m
}

func GetByName(name string) (*Content, error) {
	if u, ok := cnt[name]; ok {
		return u, nil
	}
	return nil, fmt.Errorf("%s not found", name)
}

func (u *Content) RandomImage() (configuration.ImageInfo, error) {
	if len(u.Images) == 0 {
		return configuration.ImageInfo{}, fmt.Errorf("no images for %s", u.Name)
	}
	if len(u.Images) == 1 {
		return u.Images[0], nil
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
