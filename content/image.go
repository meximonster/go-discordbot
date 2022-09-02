package content

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

type Image struct {
	Text string
	Url  string
}

func (i *Image) Validate() error {
	_, err := url.ParseRequestURI(i.Url)
	if err != nil {
		return err
	}
	return nil
}

func AddImage(table string, text string, url string) error {
	img := Image{
		Text: text,
		Url:  url,
	}
	err := img.Validate()
	if err != nil {
		return err
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
