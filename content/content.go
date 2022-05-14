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
	for _, c := range cnt {
		if c.IsHuman {
			m[c.Name] = c
		}
	}
	return m
}

func GetPets() map[string]*Content {
	m := make(map[string]*Content, len(cnt))
	for _, c := range cnt {
		if c.IsPet {
			m[c.Name] = c
		}
	}
	return m
}

func GetByName(name string) (*Content, error) {
	if c, ok := cnt[name]; ok {
		return c, nil
	}
	return nil, fmt.Errorf("%s not found", name)
}

func (c *Content) RandomImage() (configuration.ImageInfo, error) {
	if len(c.Images) == 0 {
		return configuration.ImageInfo{}, fmt.Errorf("no images for %s", c.Name)
	}
	if len(c.Images) == 1 {
		return c.Images[0], nil
	}
	var rng int
	rand.Seed(time.Now().UnixNano())
	flag := true
	for flag {
		rng = rand.Intn(len(c.Images))
		if c.Images[rng].Url != c.LastImageURLServed {
			flag = false
		}
	}
	c.LastImageURLServed = c.Images[rng].Url
	return c.Images[rng], nil
}

func AddImage(name string, text string, url string) error {
	cfg := configuration.Read()
	newImage := configuration.ImageInfo{
		Text: text,
		Url:  url,
	}
	for i, c := range cfg.Content {
		if c.Name == name {
			index := i
			c.Images = append(c.Images, newImage)
			cfg.Content[index].Images = c.Images
			if _, ok := cnt[name]; ok {
				cnt[name].Images = append(cnt[name].Images, newImage)
			}
		}
	}
	err := configuration.Write(cfg)
	if err != nil {
		return err
	}
	return nil
}

func Set(name string, cntType string) {

}
