package content

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

var cnt map[string]Content

type Content interface {
	Type() string
	AddImage(text string, url string) error
	RandomImage() (Image, error)
	Store() error
}

type ContentBase struct {
	Name               string
	Images             []Image
	LastImageURLServed string
}

type Image struct {
	Text string
	Url  string
}

func Load() error {
	return nil
}

func AddImage(table string, text string, url string) error {
	img := Image{
		Text: text,
		Url:  url,
	}
	image, err := json.Marshal(img)
	if err != nil {
		return err
	}
	return StoreImage(table, text, string(image))
}

func RandomImage(images []Image, lastImageURLServed string) (Image, error) {
	if len(images) == 0 {
		return Image{}, fmt.Errorf("no images found")
	}
	if len(images) == 1 {
		return images[0], nil
	}
	var rng int
	rand.Seed(time.Now().UnixNano())
	flag := true
	for flag {
		rng = rand.Intn(len(images))
		if images[rng].Url != lastImageURLServed {
			flag = false
		}
	}
	return images[rng], nil
}
