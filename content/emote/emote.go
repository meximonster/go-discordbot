package emote

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/meximonster/go-discordbot/image"
)

var (
	dbC   *sqlx.DB
	table = "emotes"
)

type Emote struct {
	Name               string
	Images             []image.Image
	LastImageURLServed string
}

func NewDB(db *sqlx.DB) {
	dbC = db
}

func (e *Emote) Type() string {
	return "emote"
}

func (e *Emote) GetName() string {
	return e.Name
}

func (e *Emote) RandomImage() (image.Image, error) {
	img, err := image.RandomImage(e.Images, e.LastImageURLServed)
	if err != nil {
		return image.Image{}, err
	}
	e.LastImageURLServed = img.Url
	return img, nil
}

func (e *Emote) Store() error {
	images, err := json.Marshal(e.Images)
	if err != nil {
		return err
	}
	q := fmt.Sprintf(`INSERT INTO %s (alias,images) VALUES ($1,$2)`, table)
	dbC.MustExec(q, e.Name, images)
	return nil
}

func (e *Emote) AddImage(text string, url string) error {
	img, err := image.ValidateImage(table, text, url)
	if err != nil {
		return err
	}
	q := fmt.Sprintf(`UPDATE %s SET images = images || '%s'::jsonb WHERE alias = %s`, table, string(img), e.Name)
	dbC.MustExec(q)
	return nil
}

func GetAll() []Emote {
	emotes := []Emote{}
	dbC.Select(&emotes, `SELECT * FROM emotes`)
	return emotes
}
