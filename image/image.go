package image

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

type Image struct {
	Text string `json:"text"`
	Url  string `json:"url"`
}

func (i *Image) validateURL() error {
	_, err := url.ParseRequestURI(i.Url)
	if err != nil {
		return err
	}
	return nil
}

func ValidateImage(table string, text string, url string) ([]byte, error) {
	img := Image{
		Text: text,
		Url:  url,
	}
	err := img.validateURL()
	if err != nil {
		return nil, err
	}
	image, err := json.Marshal(img)
	if err != nil {
		return nil, err
	}
	return image, nil
}

func RandomImage(images []byte, lastImageURLServed string) (Image, error) {
	var imgs []Image
	err := json.Unmarshal(images, &imgs)
	if err != nil {
		return Image{}, fmt.Errorf("cannot unmasrhal images: %s", err.Error())
	}
	if len(imgs) == 0 {
		return Image{}, fmt.Errorf("no images found")
	}
	if len(imgs) == 1 {
		return imgs[0], nil
	}
	var rng int
	rand.Seed(time.Now().UnixNano())
	flag := true
	for flag {
		rng = rand.Intn(len(imgs))
		if imgs[rng].Url != lastImageURLServed {
			flag = false
		}
	}
	return imgs[rng], nil
}

func AddImage(previous []byte, new []byte) ([]byte, error) {
	var imgs []Image
	err := json.Unmarshal(previous, &imgs)
	if err != nil {
		return nil, err
	}
	var img Image
	err = json.Unmarshal(new, &img)
	if err != nil {
		return nil, err
	}
	imgs = append(imgs, img)
	all, err := json.Marshal(imgs)
	if err != nil {
		return nil, err
	}
	return all, nil
}
