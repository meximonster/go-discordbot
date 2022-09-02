package content

import "fmt"

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

func Load() error {
	return nil
}

func Get() map[string]Content {
	return cnt
}

func GetOne(name string) (Content, error) {
	if _, ok := cnt[name]; !ok {
		return nil, fmt.Errorf("%s doesn't exist", name)
	}
	return cnt[name], nil
}

func List(contentType string) []string {
	var s []string
	for k, v := range cnt {
		if contentType == v.Type() {
			s = append(s, k)
		}
	}
	return s
}
