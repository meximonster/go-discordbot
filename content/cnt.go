package content

import "fmt"

var Cnt map[string]Content

type Content interface {
	Type() string
	GetName() string
	AddImage(text string, url string) error
	RandomImage() (Image, error)
	Store() error
}

type ContentBase struct {
	Name               string
	Images             []Image
	LastImageURLServed string
}

func Load() error {
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
